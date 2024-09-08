package pacman

import "os/exec"

func Sync(config string) bool {

	var result bool

	cmd := exec.Command("pacman", "-Sy", "--noconfirm", "--config", config)
	err := cmd.Run()

	if err == nil {
		result = true
	}

	return result

}
