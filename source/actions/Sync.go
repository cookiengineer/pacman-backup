package actions

import "pacman-backup/pacman"
import "os"

func Sync(mirror_url string, sync_folder string, pkgs_folder string) bool {

	var result bool

	config := pacman.NewConfig(mirror_url, sync_folder, pkgs_folder)

	err1 := os.WriteFile("/tmp/pacman-backup.conf", []byte(config.String()), 0666)

	if err1 == nil {
		result = pacman.Sync("/tmp/pacman-backup.conf")
	}

	return result

}
