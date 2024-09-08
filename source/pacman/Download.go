package pacman

import "os/exec"

func Download(config string, name string) bool {

	var result bool

	cmd := exec.Command("pacman", "-Sw", "--noconfirm", "--config", config, name)
	err := cmd.Run()

	if err == nil {
		result = true
	}

	return result

}
