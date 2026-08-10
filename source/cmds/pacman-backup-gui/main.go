package main

import "pacman-backup/gui/controllers"
import "pacman-backup/gui/gtk"
import "os"

func main() {

	app := gtk.NewApplication("engineer.cookie.pacman-backup-gui")

	app.OnActivate(func() {

		window := gtk.NewWindow(app)
		window.SetTitle("Pacman Backup")
		window.SetDefaultSize(800, 600)

		notebook := gtk.NewNotebook()
		window.SetChild(notebook.AsPtr())

		update_view  := controllers.NewUpdate(window)
		update_label := gtk.NewLabel("Update")
		notebook.AppendPage(update_view.AsPtr(), update_label.AsPtr())

		export_view  := controllers.NewExport(window)
		export_label := gtk.NewLabel("Export")
		notebook.AppendPage(export_view.AsPtr(), export_label.AsPtr())

		import_view  := controllers.NewImport(window)
		import_label := gtk.NewLabel("Import")
		notebook.AppendPage(import_view.AsPtr(), import_label.AsPtr())

		mirror_view  := controllers.NewMirror(window)
		mirror_label := gtk.NewLabel("Mirror")
		notebook.AppendPage(mirror_view.AsPtr(), mirror_label.AsPtr())

		window.Present()

	})

	os.Exit(app.Run())

}
