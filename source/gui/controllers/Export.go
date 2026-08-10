package controllers

import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/structs"
import "os"

func NewExport(window *gtk.Window) *views.Export {

	var view *views.Export

	view = views.NewExport(window.AsPtr(),
		func(folder string) {

			if folder == "" {
				view.SetStatus("<span foreground='red'>Please select a folder</span>")
				return
			}

			syncFolder := folder + "/sync"
			pkgsFolder := folder + "/pkgs"
			os.MkdirAll(syncFolder, 0755)
			os.MkdirAll(pkgsFolder, 0755)

			view.ShowTerminal()
			view.ClearTerminal()
			view.SetStatus("Exporting ...")

			go func() {

				console := structs.NewConsole(nil, nil, 0)
				result := actions.Export(console, syncFolder, pkgsFolder)

				gtk.RunOnMain(func() {

					view.RenderConsole(console)
					view.ScrollToBottom()

					if result {
						view.SetStatus("Export completed successfully")
					} else {
						view.SetStatus("<span foreground='red'>Export failed</span>")
					}

				})

			}()

		},
		func(folder string) {

			if folder == "" {
				view.SetStatus("<span foreground='red'>Please select a folder</span>")
				return
			}

			syncFolder := folder + "/sync"
			pkgsFolder := folder + "/pkgs"

			view.ShowTerminal()
			view.ClearTerminal()
			view.SetStatus("Cleaning up ...")

			go func() {

				console := structs.NewConsole(nil, nil, 0)
				result := actions.Cleanup(console, syncFolder, pkgsFolder)

				gtk.RunOnMain(func() {

					view.RenderConsole(console)
					view.ScrollToBottom()

					if result {
						view.SetStatus("Cleanup completed successfully")
					} else {
						view.SetStatus("<span foreground='red'>Cleanup failed</span>")
					}

				})

			}()

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
