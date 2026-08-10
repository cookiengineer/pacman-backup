package views

import "pacman-backup/gui/gtk"
import "pacman-backup/structs"
import "fmt"

func RenderConsole(console *structs.Console, text_view *gtk.TextView, offset int) {

	messages := console.Messages

	if len(messages) > offset {

		for m := offset; m < len(messages); m++ {

			message := messages[m]

			switch message.Method {
			case "Error":
				text_view.Append(fmt.Sprintf("%s | %s\n", "[ERROR]", message.Value))
			case "Warn":
				text_view.Append(fmt.Sprintf("%s | %s\n", "[WARN] ", message.Value))
			case "Info":
				text_view.Append(fmt.Sprintf("%s | %s\n", "[INFO] ", message.Value))
			case "Progress":
				// Do Nothing
			case "Group":
				text_view.Append(fmt.Sprintf("%s-\\ %s\n", "-------", message.Value))
			case "GroupEnd":
				text_view.Append(fmt.Sprintf("%s-/ %s\n", "-------", message.Value))
			case "Log":
				text_view.Append(fmt.Sprintf("%s | %s\n", "       ", message.Value))
			default:
			}

		}

	}

}
