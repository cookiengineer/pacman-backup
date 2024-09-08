package console

import "os"
import "strings"

func GroupEndResult(result bool, message string) {

	if features[FeatureGroup] == true {

		if OFFSET > 0 {
			OFFSET--
		}

		message = sanitize(message)
		offset := toOffset()

		if result == true {

			message = strings.TrimSpace(message + " succeeded")
			MESSAGES = append(MESSAGES, NewMessage("GroupEnd", message))

			if COLORS == true {
				os.Stdout.WriteString("\u001b[42m" + offset + "\\" + toSeparator(message) + message + "\u001b[K\u001b[0m\n")
			} else {
				os.Stdout.WriteString(offset + "\\" + toSeparator(message) + message + "\n")
			}

		} else {

			message = strings.TrimSpace(message + " failed")
			MESSAGES = append(MESSAGES, NewMessage("GroupEnd", message))

			if COLORS == true {
				os.Stdout.WriteString("\u001b[41m" + offset + "\\" + toSeparator(message) + message + "\u001b[K\u001b[0m\n")
			} else {
				os.Stdout.WriteString(offset + "\\" + toSeparator(message) + message + "\n")
			}

		}

	}

}
