package pacman

import "os/exec"

func Upgrade(config string) bool {

	var result bool

	cmd := exec.Command("pacman", "-Su", "--noconfirm", "--config", config)
	err := cmd.Run()

	if err == nil {
		result = true
	}

	return result

}
