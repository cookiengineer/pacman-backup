package controllers

import "pacman-backup/gui/dialogs"
import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/pacman"
import "pacman-backup/structs"
import "pacman-backup/sudo"
import "fmt"
import "time"

func NewUpdate(window *gtk.Window) *views.Update {

	var view *views.Update

	config := pacman.InitConfig("/etc/pacman.conf")
	dbpath := config.Options.DBPath
	cachedir := config.Options.CacheDir

	view = views.NewUpdate(window.AsPtr(), config.Repositories.Core,
		func() {

			dialogs.RequestSudo(window, func() {

				if sudo.NeedsSudo() == true {
					return
				}

				mirror := view.GetSelectedMirror()

				view.ShowTerminal()
				view.ClearTerminal()
				view.AppendTerminal(fmt.Sprintf("Synchronizing databases and packages from %s ...\n\n", mirror))

				console := structs.NewConsole(nil, nil, 0)
				done    := make(chan bool, 1)
				ticker  := time.NewTicker(100 * time.Millisecond)

				go func() {
					done <- actions.Download(console, mirror, dbpath+"/sync", cachedir)
				}()

				go func() {

					for {

						select {
						case <-ticker.C:

							gtk.RunOnMain(func() {
								view.RenderTerminal(console)
								view.ScrollToBottom()
								view.SetPacnewFiles(scan_pacnew_files())
							})

						case <-done:

							gtk.RunOnMain(func() {
								view.RenderTerminal(console)
								view.ScrollToBottom()
								view.SetPacnewFiles(scan_pacnew_files())
							})

							return

						}

					}

				}()

			})

		},
	)

	return view

}
