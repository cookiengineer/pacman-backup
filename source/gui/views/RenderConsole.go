package views

import "pacman-backup/gui/gtk"
import "pacman-backup/structs"

func RenderConsole(console *structs.Console, textView *gtk.TextView) {

	messages := console.Messages

	for _, msg := range messages {

		prefix := ""

		switch msg.Method {
		case "Error":
			prefix = "[ERROR] "
		case "Warn":
			prefix = "[WARN]  "
		case "Info":
			prefix = "[INFO]  "
		case "Progress":
			prefix = "        "
		case "Group":
			prefix = "------> "
		case "GroupEnd":
			prefix = "<------ "
		default:
			prefix = "        "
		}

		textView.Append(prefix + msg.Value + "\n")

	}

}
