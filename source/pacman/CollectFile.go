package pacman

import "pacman-backup/structs"
import "os/exec"

func CollectFile(config string, filepath string) structs.Package {

	var result structs.Package

	cmd := exec.Command("pacman", "-Qpi", "--noconfirm", "--config", config, filepath)
	buffer, err := cmd.Output()

	if err == nil {
		ParsePackage(string(buffer), &result)
	}

	return result

}
