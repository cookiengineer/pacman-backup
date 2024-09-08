package console

import "os"

func GroupEnd(message string) {

	if features[FeatureGroup] == true {

		if OFFSET > 0 {
			OFFSET--
		}

		message = sanitize(message)
		offset := toOffset()
		MESSAGES = append(MESSAGES, NewMessage("GroupEnd", message))

		if COLORS == true {
			os.Stdout.WriteString("\u001b[40m" + offset + "\\" + toSeparator(message) + message + "\u001b[K\u001b[0m\n")
		} else {
			os.Stdout.WriteString(offset + "\\" + toSeparator(message) + message + "\n")
		}

	}

}
