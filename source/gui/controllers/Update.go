package controllers

import "pacman-backup/gui/dialogs"
import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/pacman"
import "pacman-backup/structs"
import "pacman-backup/sudo"
import "fmt"

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

			dialogs.RequestSudo(window, func() {

				if sudo.NeedsSudo() == true {
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
