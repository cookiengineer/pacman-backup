package actions

import "pacman-backup/pacman"
import "pacman-backup/structs"
import "os"

func Sync(console *structs.Console, mirror_url string, sync_folder string, pkgs_folder string) bool {

	console.Group("actions/Sync")

	var result bool

	config := pacman.NewConfig(mirror_url, sync_folder, pkgs_folder)
	err1 := os.WriteFile("/tmp/pacman-backup.conf", []byte(config.String()), 0666)

	if err1 == nil {

		err2 := pacman.Sync("/tmp/pacman-backup.conf")

		if err2 == nil {
			result = true
		} else {
			console.Error(err2.Error())
			result = false
		}

	}

	if result {
		console.GroupEnd("actions/Sync succeeded")
	} else {
		console.GroupEnd("actions/Sync failed")
	}

	return result

}
