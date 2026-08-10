package dialogs

import "pacman-backup/gui/gtk"
import "pacman-backup/sudo"

func RequestSudo(parent *gtk.Window, on_result func()) {

	if sudo.NeedsSudo() == true {

		gtk.ShowSudoDialog(parent.AsPtr(), func(password string) {

			if password != "" {
				sudo.SetPassword(password)
			}

			on_result()

		})

	} else {
		on_result()
	}

}
