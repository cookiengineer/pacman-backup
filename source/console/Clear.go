package console

import "os"

func Clear() {

	var has_errors bool = false

	for m := 0; m < len(MESSAGES); m++ {

		if MESSAGES[m].Method == "Error" {
			has_errors = true
			break
		}

	}

	MESSAGES = append(MESSAGES, NewMessage("Clear", ""))

	if has_errors == false {

		// clear screen and reset cursor
		os.Stdout.WriteString("\u001b[2J\u001b[0f")

		// clear scroll buffer
		os.Stdout.WriteString("\u001b[3J")

	}

}
