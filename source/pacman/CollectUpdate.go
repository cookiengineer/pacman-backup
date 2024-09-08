package pacman

import "pacman-backup/structs"
import "os"
import "os/exec"

func CollectUpdate(config string, name string) structs.Package {

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	var result structs.Package = structs.NewPackage("pacman")

	cmd := exec.Command("pacman", "-Si", "--noconfirm", "--config", config, name)
	buffer, err := cmd.Output()

	if err == nil {
		ParsePackage(string(buffer), &result)
	}

	return result

}
