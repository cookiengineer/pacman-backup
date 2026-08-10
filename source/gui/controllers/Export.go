package controllers

import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/structs"
import "os"
import "time"

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

			console := structs.NewConsole(nil, nil, 0)
			done    := make(chan bool, 1)
			ticker  := time.NewTicker(100 * time.Millisecond)

			go func() {
				done <- actions.Export(console, syncFolder, pkgsFolder)
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
								view.SetStatus("Export completed successfully")
							} else {
								view.SetStatus("<span foreground='red'>Export failed</span>")
							}

						})

						return

					}

				}

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

			console := structs.NewConsole(nil, nil, 0)
			done    := make(chan bool, 1)
			ticker  := time.NewTicker(100 * time.Millisecond)

			go func() {
				done <- actions.Cleanup(console, syncFolder, pkgsFolder)
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
								view.SetStatus("Cleanup completed successfully")
							} else {
								view.SetStatus("<span foreground='red'>Cleanup failed</span>")
							}

						})

						return

					}

				}

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
