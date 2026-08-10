package pacman

import "pacman-backup/structs"
import "pacman-backup/sudo"
import "os"
import "os/exec"

func CollectUpdate(config string, name string) (structs.Package, error) {

	var result structs.Package = structs.NewPackage("pacman")
	var err error = nil

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	var cmd1 *exec.Cmd

	if NeedsSudo(config) == true {
		cmd1 = sudo.Command("pacman", "-Si", "--noconfirm", "--config", config, name)
	} else {
		cmd1 = exec.Command("pacman", "-Si", "--noconfirm", "--config", config, name)
	}

	buffer, err1 := cmd1.Output()

	if err1 == nil {
		ParsePackage(string(buffer), &result)
	} else {
		err = err1
	}

	return result, err

}
