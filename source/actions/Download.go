package actions

import "pacman-backup/pacman"
import "pacman-backup/structs"
import "os"
import "strconv"
import "strings"

func Download(console *structs.Console, mirror_url string, sync_folder string, pkgs_folder string) bool {

	console.Group("actions/Download")

	result := false
	config := pacman.NewConfig(mirror_url, sync_folder, pkgs_folder)
	err1 := os.WriteFile("/tmp/pacman-backup.conf", []byte(config.String()), 0666)

	if err1 == nil {

		updates := pacman.CollectUpdates(console, "/tmp/pacman-backup.conf")
		result = true

		if len(updates) > 0 {

			console.Log("Found " + strconv.Itoa(len(updates)) + " Updates")

			for u := 0; u < len(updates); u++ {

				update := updates[u]

				console.Progress("Package " + update.Name + "@" + update.Version.String())

				err := pacman.Download("/tmp/pacman-backup.conf", update.Name)

				if err != nil {

					message := err.Error()

					if strings.Contains(message, "you cannot perform this operation unless you are root") {
						result = false
						break
					} else {
						console.Error("Package " + update.Name + "@" + update.Version.String() + " not downloaded")
						result = false
					}

				}

			}

		} else {
			console.Warn("Found 0 Updates")
		}

	} else {
		console.Error(err1.Error())
	}

	if result {
		console.GroupEnd("actions/Download succeeded")
	} else {
		console.GroupEnd("actions/Download failed")
	}

	return result

}
