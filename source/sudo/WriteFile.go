package sudo

import "bytes"
import "os"
import "os/exec"

func WriteFile(file string, data []byte, permission os.FileMode) error {

	if NeedsSudo() {

		var stdin bytes.Buffer

		stdin.WriteString(GetPassword() + "\n")
		stdin.Write(data)

		cmd := exec.Command("sudo", "-S", "-k", "dd", "of="+file, "status=none")
		cmd.Stdin = &stdin

		// TODO: Ensure correct file permissions, use chmod?

		return cmd.Run()

	}

	return os.WriteFile(file, data, permission)

}
