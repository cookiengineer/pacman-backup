package controllers

import "pacman-backup/gui/dialogs"
import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/pacman"
import "pacman-backup/structs"
import "pacman-backup/sudo"
import "os"
import "time"

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

				console := structs.NewConsole(nil, nil, 0)
				done    := make(chan bool, 1)
				ticker  := time.NewTicker(100 * time.Millisecond)

				go func() {
					done <- actions.Import(console, syncFolder, pkgsFolder)
				}()

				go func() {

					for {

						select {
						case <-ticker.C:

							gtk.RunOnMain(func() {
								view.RenderTerminal(console)
								view.ScrollToBottom()
							})

						case result := <-done:

							gtk.RunOnMain(func() {

								view.RenderTerminal(console)
								view.ScrollToBottom()

								if result {
									view.SetStatus("Import completed successfully")
								} else {
									view.SetStatus("<span foreground='red'>Import failed</span>")
								}

							})

							return

						}

					}

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

				console := structs.NewConsole(nil, nil, 0)
				done    := make(chan bool, 1)
				ticker  := time.NewTicker(100 * time.Millisecond)

				go func() {
					done <- actions.Upgrade(console, mirror, syncFolder, pkgsFolder)
				}()

				go func() {

					for {

						select {
						case <-ticker.C:

							gtk.RunOnMain(func() {
								view.RenderTerminal(console)
								view.ScrollToBottom()
							})

						case result := <-done:

							gtk.RunOnMain(func() {

								view.RenderTerminal(console)
								view.ScrollToBottom()

								if result {
									view.SetStatus("Upgrade completed successfully")
								} else {
									view.SetStatus("<span foreground='red'>Upgrade failed</span>")
								}

							})

							return

						}

					}

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
