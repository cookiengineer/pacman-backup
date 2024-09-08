package actions

import "pacman-backup/console"
import "pacman-backup/pacman"
import "os"

func Upgrade(sync_folder string, pkgs_folder string) bool {

	console.Group("Upgrade")

	var result bool

	config := pacman.NewConfig("", sync_folder, pkgs_folder)
	err1 := os.WriteFile("/tmp/pacman-backup.conf", []byte(config.String()), 0666)

	if err1 == nil {

		files := pacman.CollectFiles("/tmp/pacman-backup.conf", pkgs_folder)
		updates := pacman.CollectUpdates("/tmp/pacman-backup.conf")
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

				if is_cached == false {
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
