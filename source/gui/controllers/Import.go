package controllers

import "pacman-backup/gui/dialogs"
import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/pacman"
import "pacman-backup/structs"
import "pacman-backup/sudo"
import "os"

func NewImport(window *gtk.Window) *views.Import {

	var view *views.Import

	view = views.NewImport(window.AsPtr(),
		func(folder string) {

			if folder == "" {
				view.SetStatus("<span foreground='red'>Please select a folder</span>")
				return
			}

			dialogs.RequestSudo(window, func() {

				if sudo.NeedsSudo() == true {
					view.SetStatus("<span foreground='red'>Sudo password required</span>")
					return
				}

				syncFolder := folder + "/sync"
				pkgsFolder := folder + "/pkgs"
				os.MkdirAll(syncFolder, 0755)
				os.MkdirAll(pkgsFolder, 0755)

				view.ShowTerminal()
				view.ClearTerminal()
				view.SetStatus("Importing ...")

				go func() {

					console := structs.NewConsole(nil, nil, 0)
					result := actions.Import(console, syncFolder, pkgsFolder)

					gtk.RunOnMain(func() {

						view.RenderConsole(console)
						view.ScrollToBottom()

						if result {
							view.SetStatus("Import completed successfully")
						} else {
							view.SetStatus("<span foreground='red'>Import failed</span>")
						}

					})

				}()

			})

		},
		func(folder string) {

			if folder == "" {
				view.SetStatus("<span foreground='red'>Please select a folder</span>")
				return
			}

			dialogs.RequestSudo(window, func() {

				if sudo.NeedsSudo() == true {
					view.SetStatus("<span foreground='red'>Sudo password required</span>")
					return
				}

				config := pacman.InitConfig("/etc/pacman.conf")
				mirror := config.ToMirror()
				syncFolder := folder + "/sync"
				pkgsFolder := folder + "/pkgs"
				os.MkdirAll(syncFolder, 0755)
				os.MkdirAll(pkgsFolder, 0755)

				view.ShowTerminal()
				view.ClearTerminal()
				view.SetStatus("Upgrading ...")

				go func() {

					console := structs.NewConsole(nil, nil, 0)
					result := actions.Upgrade(console, mirror, syncFolder, pkgsFolder)

					gtk.RunOnMain(func() {

						view.RenderConsole(console)
						view.ScrollToBottom()

						if result {
							view.SetStatus("Upgrade completed successfully")
						} else {
							view.SetStatus("<span foreground='red'>Upgrade failed</span>")
						}

					})

				}()

			})

		},
		func(folder_entry *gtk.Entry) {

			gtk.ShowFolderDialog(window.AsPtr(), func(path string) {

				if path != "" {
					gtk.RunOnMain(func() {
						folder_entry.SetText(path)
					})
				}

			})

		},
	)

	return view

}
