package main

import "pacman-backup/gui/controllers"
import "pacman-backup/gui/gtk"
import "bytes"
import "os"
import "os/exec"
import "strings"

func runSudoStream(password string, onLine func(string), onDone func(), name string, args ...string) {
	cmdArgs := append([]string{"-S", name}, args...)
	cmd := exec.Command("sudo", cmdArgs...)
	cmd.Stdin = strings.NewReader(password + "\n")

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	cmd.Start()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := stdoutPipe.Read(buf)
			if n > 0 {
				text := string(buf[:n])
				gtk.RunOnMain(func() { onLine(text) })
			}
			if err != nil {
				break
			}
		}
	}()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := stderrPipe.Read(buf)
			if n > 0 {
				text := string(buf[:n])
				gtk.RunOnMain(func() { onLine(text) })
			}
			if err != nil {
				break
			}
		}
	}()

	go func() {
		cmd.Wait()
		gtk.RunOnMain(func() { onDone() })
	}()
}

func runSudo(password string, name string, args ...string) (string, string) {
	cmdArgs := append([]string{"-S", name}, args...)
	cmd := exec.Command("sudo", cmdArgs...)
	cmd.Stdin = strings.NewReader(password + "\n")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	output := stdout.String()
	errOutput := stderr.String()

	if strings.Contains(errOutput, "incorrect password") {
		return output, "Incorrect sudo password"
	}

	if err != nil {
		return output, errOutput + "\n" + err.Error()
	}

	return output, ""
}

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
