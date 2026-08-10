package dialogs

import "pacman-backup/gui/gtk"
import "unsafe"

var cached_sudo_password string

func RequestPassword(parent unsafe.Pointer, on_result func(string)) {

	if cached_sudo_password != "" {
		on_result(cached_sudo_password)
		return
	}

	gtk.ShowPasswordDialog(parent, func(password string) {

		if password != "" {
			cached_sudo_password = password
		}

		on_result(password)

	})

}
