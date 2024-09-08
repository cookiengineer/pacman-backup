package console

import "os"
import "strings"

func Info(message string) {

	if features[FeatureInfo] == true {

		message = sanitize(message)
		offset := toOffset()
		MESSAGES = append(MESSAGES, NewMessage("Info", message))

		if strings.Contains(message, "\n") {

			var lines = strings.Split(message, "\n")

			if COLORS == true {

				for l := 0; l < len(lines); l++ {
					os.Stdout.WriteString("\u001b[42m" + offset + toSeparator(lines[l]) + lines[l] + "\u001b[K\n")
				}

				os.Stdout.WriteString("\u001b[0m")

			} else {

				for l := 0; l < len(lines); l++ {
					os.Stdout.WriteString(offset + toSeparator(lines[l]) + lines[l] + "\n")
				}

			}

		} else {

			if COLORS == true {
				os.Stdout.WriteString("\u001b[42m" + offset + toSeparator(message) + message + "\u001b[K\u001b[0m\n")
			} else {
				os.Stdout.WriteString(offset + toSeparator(message) + message + "\n")
			}

		}

	}

}
