package pacman

import "pacman-backup/console"
import "os/exec"

func Sync(config string) bool {

	var result bool

	cmd := exec.Command("pacman", "-Sy", "--noconfirm", "--config", config)
	err := cmd.Run()

	if err == nil {
		result = true
	} else {
		console.Error(err.Error())
	}

	return result

}
