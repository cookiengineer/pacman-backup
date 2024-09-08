package console

import "os"

func Group(message string) {

	if features[FeatureGroup] == true {

		offset := toOffset()
		message = sanitize(message)
		MESSAGES = append(MESSAGES, NewMessage("Group", message))
		OFFSET++

		if COLORS == true {
			os.Stdout.WriteString("\u001b[40m" + offset + "/" + toSeparator(message) + message + "\u001b[K\u001b[0m\n")
		} else {
			os.Stdout.WriteString(offset + "/" + toSeparator(message) + message + "\n")
		}

	}

}
