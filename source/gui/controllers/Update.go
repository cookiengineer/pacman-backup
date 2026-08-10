package controllers

import "pacman-backup/gui/dialogs"
import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/pacman"
import "pacman-backup/structs"
import "fmt"

func NewUpdate(window *gtk.Window) *views.Update {

	var view *views.Update

	view = views.NewUpdate(window.AsPtr(),
		func() {

			dialogs.RequestPassword(window.AsPtr(), func(password string) {

				if password == "" {
					return
				}

				config := pacman.InitConfig()
				mirror := config.ToMirror()

				view.ShowTerminal()
				view.ClearTerminal()
				view.AppendTerminal(fmt.Sprintf("Synchronizing databases from %s ...\n\n", mirror))

				go func() {

					console := structs.NewConsole(nil, nil, 0)
					actions.Download(console, mirror, config.Options.DBPath+"/sync", config.Options.CacheDir)

					gtk.RunOnMain(func() {
						view.RenderConsole(console)
						view.ScrollToBottom()
						view.SetPacnewFiles(scan_pacnew_files())
					})

				}()

			})

		},
		func() {

			dialogs.RequestPassword(window.AsPtr(), func(password string) {

				if password == "" {
					return
				}

				config := pacman.InitConfig()
				mirror := config.ToMirror()

				view.ShowTerminal()
				view.ClearTerminal()
				view.AppendTerminal("Downloading packages from " + mirror + " ...\n\n")

				go func() {

					console := structs.NewConsole(nil, nil, 0)
					actions.Download(console, mirror, config.Options.DBPath+"/sync", config.Options.CacheDir)

					gtk.RunOnMain(func() {
						view.RenderConsole(console)
						view.ScrollToBottom()
						view.SetPacnewFiles(scan_pacnew_files())
					})

				}()

			})

		},
	)

	return view

}
