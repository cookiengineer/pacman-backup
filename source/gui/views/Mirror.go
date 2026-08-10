package views

import "pacman-backup/gui/gtk"
import "pacman-backup/structs"
import "unsafe"

type Mirror struct {
	Folder           *gtk.Entry
	Mirror           *gtk.Entry
	Start            *gtk.Button
	Stop             *gtk.Button
	root             *gtk.Box
	status           *gtk.Label
	terminal         *gtk.TextView
	terminal_wrapper *gtk.ScrolledWindow
}

func (view *Mirror) AsPtr() unsafe.Pointer {
	return view.root.AsPtr()
}

func NewMirror(parent unsafe.Pointer, onStart func(folder string), onStop func(), onBrowse func(entry *gtk.Entry)) *Mirror {

	view := &Mirror{}

	view.root = gtk.NewBox(gtk.OrientationVertical, 8)
	view.root.SetMarginStart(12)
	view.root.SetMarginEnd(12)
	view.root.SetMarginTop(12)
	view.root.SetMarginBottom(12)

	server_header := gtk.NewLabel("")
	server_header.SetMarkup("<b>Local Mirror</b>")
	server_header.SetXAlign(0.0)
	server_header.SetMarginBottom(2)
	view.root.Append(server_header.AsPtr())

	server_description := gtk.NewLabel("Serve a local mirror for offline machines on the network.")
	server_description.SetWrap(true)
	server_description.SetXAlign(0.0)
	server_description.SetMarginBottom(12)
	view.root.Append(server_description.AsPtr())

	folder_row := gtk.NewBox(gtk.OrientationHorizontal, 8)
	folder_row.SetMarginBottom(8)
	view.root.Append(folder_row.AsPtr())

	folder_label := gtk.NewLabel("Folder:")
	folder_label.SetMarginEnd(4)
	folder_row.Append(folder_label.AsPtr())

	view.Folder = gtk.NewEntry()
	view.Folder.SetPlaceholder("/mnt/usb-drive")
	view.Folder.SetHExpand(true)
	folder_row.Append(view.Folder.AsPtr())

	browse_button := gtk.NewButton("Browse")
	browse_button.OnClick(func() {
		if onBrowse != nil {
			onBrowse(view.Folder)
		}
	})
	folder_row.Append(browse_button.AsPtr())

	mirror_row := gtk.NewBox(gtk.OrientationHorizontal, 8)
	mirror_row.SetMarginBottom(12)
	view.root.Append(mirror_row.AsPtr())

	mirror_label := gtk.NewLabel("Mirror URL:")
	mirror_label.SetMarginEnd(4)
	mirror_row.Append(mirror_label.AsPtr())

	view.Mirror = gtk.NewEntry()
	view.Mirror.SetText("http://localhost:15678")
	view.Mirror.SetHExpand(true)
	mirror_row.Append(view.Mirror.AsPtr())

	serve_row := gtk.NewBox(gtk.OrientationHorizontal, 8)
	serve_row.SetMarginBottom(2)
	view.root.Append(serve_row.AsPtr())

	view.Start = gtk.NewButton("Serve Mirror")
	view.Start.OnClick(func() {

		if onStart != nil {

			folder := view.Folder.Text()
			view.Start.SetSensitive(false)
			view.Stop.SetSensitive(true)

			onStart(folder)

		}

	})
	serve_row.Append(view.Start.AsPtr())

	view.Stop = gtk.NewButton("Stop Mirror")
	view.Stop.SetSensitive(false)
	view.Stop.OnClick(func() {

		if onStop != nil {

			view.Stop.SetSensitive(false)
			view.Start.SetSensitive(true)
			onStop()

		}

	})
	serve_row.Append(view.Stop.AsPtr())

	view.terminal_wrapper = gtk.NewScrolledWindow()
	view.terminal_wrapper.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)
	view.terminal_wrapper.SetVExpand(true)
	view.terminal_wrapper.SetSizeRequest(-1, 200)
	view.terminal_wrapper.SetMarginBottom(8)
	view.terminal_wrapper.SetVisible(false)
	view.root.Append(view.terminal_wrapper.AsPtr())

	view.terminal = gtk.NewTextView()
	view.terminal.SetEditable(false)
	view.terminal.SetMonospace(true)
	view.terminal.SetCursorVisible(false)
	view.terminal.SetTerminalStyle()
	view.terminal_wrapper.SetChild(view.terminal.AsPtr())

	view.status = gtk.NewLabel("")
	view.status.SetWrap(true)
	view.status.SetXAlign(0.0)
	view.status.SetMarginBottom(4)
	view.root.Append(view.status.AsPtr())

	return view

}

func (view *Mirror) ShowTerminal() {
	view.terminal_wrapper.SetVisible(true)
}

func (view *Mirror) RenderConsole(console *structs.Console) {
	RenderConsole(console, view.terminal)
}

func (view *Mirror) AppendTerminal(text string) {
	view.terminal.Append(text)
}

func (view *Mirror) ClearTerminal() {
	view.terminal.Clear()
}

func (view *Mirror) ScrollToBottom() {
	view.terminal.ScrollToBottom()
}

func (view *Mirror) SetStatus(text string) {
	view.status.SetMarkup(text)
}
