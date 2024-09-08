package actions

import "pacman-backup/console"
import "pacman-backup/pacman"
import "os"

func Download(mirror_url string, sync_folder string, pkgs_folder string) bool {

	console.Group("Download")

	config := pacman.NewConfig(mirror_url, sync_folder, pkgs_folder)
	err1 := os.WriteFile("/tmp/pacman-backup.conf", []byte(config.String()), 0666)

	if err1 == nil {

		updates := pacman.CollectUpdates("/tmp/pacman-backup.conf")
		result := true

		if len(updates) > 0 {

			for u := 0; u < len(updates); u++ {

				update := updates[u]

				console.Progress("> Package " + update.Name + " " + update.Version.String())

				check := pacman.Download("/tmp/pacman-backup.conf", update.Name)

				if check == false {
					console.Error("> Package " + update.Name + " " + update.Version.String() + " not downloaded")
					result = false
				}

			}

		}

		return result

	}

	console.GroupEnd("")

	return false

}
