package views

import "pacman-backup/gui/gtk"
import "pacman-backup/structs"
import "unsafe"

type Export struct {
	Folder           *gtk.Entry
	root             *gtk.Box
	status           *gtk.Label
	terminal         *gtk.TextView
	terminal_length  int
	terminal_wrapper *gtk.ScrolledWindow
}

func (view *Export) AsPtr() unsafe.Pointer {
	return view.root.AsPtr()
}

func NewExport(parent unsafe.Pointer, onExport func(folder string), onCleanup func(folder string), onBrowse func(entry *gtk.Entry)) *Export {

	view := &Export{}

	view.root = gtk.NewBox(gtk.OrientationVertical, 8)
	view.root.SetMarginStart(12)
	view.root.SetMarginEnd(12)
	view.root.SetMarginTop(12)
	view.root.SetMarginBottom(12)

	header := gtk.NewLabel("")
	header.SetMarkup("<b>Export and Cleanup</b>")
	header.SetXAlign(0.0)
	header.SetMarginBottom(2)
	view.root.Append(header.AsPtr())

	description := gtk.NewLabel("Export pacman database and packages to external storage, then clean up old package versions.")
	description.SetWrap(true)
	description.SetXAlign(0.0)
	description.SetMarginBottom(12)
	view.root.Append(description.AsPtr())

	folder_row := gtk.NewBox(gtk.OrientationHorizontal, 8)
	folder_row.SetMarginBottom(12)
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

	buttons := gtk.NewBox(gtk.OrientationHorizontal, 8)
	buttons.SetMarginBottom(12)
	view.root.Append(buttons.AsPtr())

	export_button := gtk.NewButton("Export Packages")
	export_button.OnClick(func() {
		if onExport != nil {
			folder := view.Folder.Text()
			onExport(folder)
		}
	})
	buttons.Append(export_button.AsPtr())

	cleanup_button := gtk.NewButton("Cleanup Packages")
	cleanup_button.OnClick(func() {
		if onCleanup != nil {
			folder := view.Folder.Text()
			onCleanup(folder)
		}
	})
	buttons.Append(cleanup_button.AsPtr())

	view.status = gtk.NewLabel("")
	view.status.SetWrap(true)
	view.status.SetXAlign(0.0)
	view.status.SetMarginBottom(4)
	view.root.Append(view.status.AsPtr())

	view.terminal_wrapper = gtk.NewScrolledWindow()
	view.terminal_wrapper.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)
	view.terminal_wrapper.SetVExpand(true)
	view.terminal_wrapper.SetSizeRequest(-1, 200)
	view.terminal_wrapper.SetVisible(false)
	view.root.Append(view.terminal_wrapper.AsPtr())

	view.terminal = gtk.NewTextView()
	view.terminal.SetEditable(false)
	view.terminal.SetMonospace(true)
	view.terminal.SetCursorVisible(false)
	view.terminal.SetTerminalStyle()
	view.terminal_wrapper.SetChild(view.terminal.AsPtr())

	return view

}

func (view *Export) ShowTerminal() {
	view.terminal_wrapper.SetVisible(true)
}

func (view *Export) ClearTerminal() {
	view.terminal.Clear()
	view.terminal_length = 0
}

func (view *Export) ScrollToBottom() {
	view.terminal.ScrollToBottom()
}

func (view *Export) RenderTerminal(console *structs.Console) {

	if console.Length() > view.terminal_length {
		RenderConsole(console, view.terminal, view.terminal_length)
		view.terminal_length = console.Length()
	}

}

func (view *Export) SetStatus(text string) {

	view.status.SetMarkup(text)

	if text != "" {
		view.status.SetVisible(true)
	} else {
		view.status.SetVisible(false)
	}

}
