package pacman

import "pacman-backup/sudo"
import "os"
import "os/exec"

func Sync(config string) (error) {

	var err error = nil

	os.Setenv("TZ", "Europe/Greenwich")
	os.Setenv("LC_TIME", "en_US")

	var cmd1 *exec.Cmd

	if NeedsSudo(config) == true {
		cmd1 = sudo.Command("pacman", "-Sy", "--noconfirm", "--config", config)
	} else {
		cmd1 = exec.Command("pacman", "-Sy", "--noconfirm", "--config", config)
	}

	err1 := cmd1.Run()

	if err1 == nil {
		err = nil
	} else {
		err = err1
	}

	return err

}
