package controllers

import "pacman-backup/gui/gtk"
import "pacman-backup/gui/views"
import "pacman-backup/actions"
import "pacman-backup/structs"
import "context"
import "os"
import "time"

func NewMirror(window *gtk.Window) *views.Mirror {

	var serve_context context.Context
	var serve_cancel  context.CancelFunc
	var view          *views.Mirror

	view = views.NewMirror(window.AsPtr(),
		func(folder string) {

			if folder == "" {
				view.SetStatus("<span foreground='red'>Please select a folder</span>")
				return
			}

			syncFolder := folder + "/sync"
			pkgsFolder := folder + "/pkgs"
			os.MkdirAll(syncFolder, 0755)
			os.MkdirAll(pkgsFolder, 0755)

			serve_context, serve_cancel = context.WithCancel(context.Background())
			_ = serve_context

			view.ShowTerminal()
			view.ClearTerminal()
			view.AppendTerminal("Starting mirror server on http://localhost:15678 ...\n")
			view.SetStatus("Mirror running on http://localhost:15678")

			console := structs.NewConsole(nil, nil, 0)
			done    := make(chan bool, 1)
			ticker  := time.NewTicker(100 * time.Millisecond)

			go func() {
				done <- actions.Serve(console, syncFolder, pkgsFolder)
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
								view.SetStatus("Mirror stopped successfully")
							} else {
								view.SetStatus("<span foreground='red'>Mirror failed</span>")
							}

							view.Start.SetSensitive(true)
							view.Stop.SetSensitive(false)

						})

						return

					}

				}

			}()

		},
		func() {

			if serve_cancel != nil {

				serve_cancel()
				view.SetStatus("Stopping mirror ...")

			}

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
