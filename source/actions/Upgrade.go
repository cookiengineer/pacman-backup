package actions

import "pacman-backup/console"
import "pacman-backup/pacman"
import "os"
import "strconv"
import "strings"

func isIgnored(config *pacman.Config, name_and_version string) bool {

	var result bool = false

	name := name_and_version[0:strings.Index(name_and_version, "@")]

	if result == false {

		for i := 0; i < len(config.Options.HoldPkg); i++ {

			if config.Options.HoldPkg[i] == name {
				result = true
				break
			}

		}

	}

	if result == false {

		for i := 0; i < len(config.Options.IgnoreGroup); i++ {

			if config.Options.IgnoreGroup[i] == name {
				result = true
				break
			}

		}

	}

	if result == false {

		for i := 0; i < len(config.Options.IgnorePkg); i++ {

			if config.Options.IgnorePkg[i] == name {
				result = true
				break
			}

		}

	}

	return result

}

func Upgrade(mirror_url string, sync_folder string, pkgs_folder string) bool {

	console.Group("Upgrade")

	var result bool

	local_config := pacman.InitConfig()
	config := pacman.NewConfig(mirror_url, sync_folder, pkgs_folder)
	config.Options.SyncFirst = []string{"archlinux-keyring"}

	if len(local_config.Options.IgnoreGroup) > 0 {

		for i := 0; i < len(local_config.Options.IgnoreGroup); i++ {
			config.Options.IgnoreGroup = append(config.Options.IgnoreGroup, local_config.Options.IgnoreGroup[i])
		}

	}

	if len(local_config.Options.IgnorePkg) > 0 {

		for i := 0; i < len(local_config.Options.IgnorePkg); i++ {
			config.Options.IgnorePkg = append(config.Options.IgnorePkg, local_config.Options.IgnorePkg[i])
		}

	}

	if len(local_config.Options.HoldPkg) > 0 {

		for h := 0; h < len(local_config.Options.HoldPkg); h++ {
			config.Options.HoldPkg = append(config.Options.HoldPkg, local_config.Options.HoldPkg[h])
		}

	}

	err1 := os.WriteFile("/tmp/pacman-backup.conf", []byte(config.String()), 0666)

	if err1 == nil {

		files := pacman.CollectFiles("/tmp/pacman-backup.conf", pkgs_folder)
		updates := pacman.CollectUpdates("/tmp/pacman-backup.conf")

		console.Log("Found " + strconv.Itoa(len(updates)) + " available updates")
		console.Log("Found " + strconv.Itoa(len(files)) + " cached updates")

		cache := make(map[string]bool, 0)
		verified := true

		if len(files) > 0 && len(updates) > 0 {

			for u := 0; u < len(updates); u++ {
				cache[updates[u].Name + "@" + updates[u].Version.String()] = false
			}

			for f := 0; f < len(files); f++ {

				file := files[f]

				value, ok := cache[file.Name + "@" + file.Version.String()]

				if ok == true && value == false {
					cache[file.Name + "@" + file.Version.String()] = true
				}

			}

			for name, is_cached := range cache {

				if is_cached == false && !isIgnored(&config, name) {
					console.Error("-> Package " + name + " not available")
					verified = false
				}

			}

		}

		if verified == true {
			result = pacman.Upgrade("/tmp/pacman-backup.conf")
		} else {
			console.Error("")
			console.Error("Execute \"sudo pacman-backup download\" to repair local cache")
			console.Error("")
		}

	}

	console.GroupEnd("Upgrade")

	return result

}
