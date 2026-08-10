package sudo

import "context"
import "os/exec"
import "strings"
import "time"

func Command(args ...string) *exec.Cmd {

	if IsRoot() == true {

		cmd := exec.Command(args[0], args[1:]...)

		return cmd

	} else {

		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		_ = cancel

		cmd := exec.CommandContext(ctx, "sudo", append([]string{"-S", "-k", args[0]}, args[1:]...)...)
		cmd.Stdin = strings.NewReader(GetPassword() + "\n")

		return cmd

	}

}

