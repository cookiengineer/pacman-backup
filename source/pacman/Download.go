package pacman

import "pacman-backup/console"
import "os/exec"

func Download(config string, name string) bool {

	var result bool

	cmd := exec.Command("pacman", "-Sw", "--noconfirm", "--config", config, name)
	err := cmd.Run()

	if err == nil {
		result = true
	} else {
		console.Error(err.Error())
	}

	return result

}
