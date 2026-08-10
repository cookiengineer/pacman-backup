package pacman

import "pacman-backup/sudo"
import "bytes"
import "errors"
import "os/exec"
import "strings"

func Download(config string, name string) (error) {

	var err error = nil

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd1 *exec.Cmd

	// Download without dependency checks for better UI
	if NeedsSudo(config) == true {
		cmd1 = sudo.Command("pacman", "-Swdd", "--noconfirm", "--config", config, name)
	} else {
		cmd1 = exec.Command("pacman", "-Swdd", "--noconfirm", "--config", config, name)
	}

	cmd1.Stdout = &stdout
	cmd1.Stderr = &stderr

	err1 := cmd1.Run()

	if err1 == nil {

		err = nil

	} else {

		lines := strings.Split(strings.TrimSpace(stderr.String()), "\n")

		if len(lines) > 0 {
			err = errors.New(lines[len(lines)-1])
		}

	}

	return err

}
