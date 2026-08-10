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

	config := pacman.InitConfig()
	dbpath := config.Options.DBPath
	cachedir := config.Options.CacheDir

	view = views.NewUpdate(window.AsPtr(), config.Repositories.Core,
		func() {

			dialogs.RequestPassword(window.AsPtr(), func(password string) {

				if password == "" {
					return
				}

				mirror := view.GetSelectedMirror()

				view.ShowTerminal()
				view.ClearTerminal()
				view.AppendTerminal(fmt.Sprintf("Synchronizing databases from %s ...\n\n", mirror))

				go func() {

					console := structs.NewConsole(nil, nil, 0)
					actions.Download(console, mirror, dbpath+"/sync", cachedir)

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

				mirror := view.GetSelectedMirror()

				view.ShowTerminal()
				view.ClearTerminal()
				view.AppendTerminal("Downloading packages from " + mirror + " ...\n\n")

				go func() {

					console := structs.NewConsole(nil, nil, 0)
					actions.Download(console, mirror, dbpath+"/sync", cachedir)

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
