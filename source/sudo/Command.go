package sudo

import "os/exec"
import "strings"

func Command(args ...string) *exec.Cmd {

	if IsRoot() == true {

		cmd := exec.Command(args[0], args[1:]...)

		return cmd

	} else {

		cmd := exec.Command("sudo", append([]string{"-S", "-k", args[0]}, args[1:]...)...)
		cmd.Stdin = strings.NewReader(GetPassword() + "\n")

		return cmd

	}

}

