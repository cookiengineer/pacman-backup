package console

import "os"
import "strings"

func Progress(message string) {

	if features[FeatureProgress] == true {

		message = strings.ReplaceAll(message, "\n", "")
		message = sanitize(message)
		offset := toOffset()

		if len(MESSAGES) > 0 {

			last_method := MESSAGES[len(MESSAGES)-1].Method

			if last_method == "Progress" {
				os.Stdout.WriteString("\033[A\033[2K\r")
				MESSAGES[len(MESSAGES)-1] = NewMessage("Progress", message)
			} else {
				MESSAGES = append(MESSAGES, NewMessage("Progress", message))
			}

		} else {
			MESSAGES = append(MESSAGES, NewMessage("Progress", message))
		}

		if COLORS == true {
			os.Stdout.WriteString("\u001b[40m" + offset + toSeparator(message) + message + "\u001b[K\u001b[0m\n")
		} else {
			os.Stdout.WriteString(offset + toSeparator(message) + message + "\n")
		}

	}

}
