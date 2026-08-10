package views

import "pacman-backup/gui/gtk"
import "pacman-backup/structs"
import "unsafe"

type Import struct {
	Folder           *gtk.Entry
	root             *gtk.Box
	status           *gtk.Label
	terminal         *gtk.TextView
	terminal_wrapper *gtk.ScrolledWindow
}

func (view *Import) AsPtr() unsafe.Pointer {
	return view.root.AsPtr()
}

func NewImport(parent unsafe.Pointer, onImport func(folder string), onUpgrade func(folder string), onBrowse func(entry *gtk.Entry)) *Import {

	view := &Import{}

	view.root = gtk.NewBox(gtk.OrientationVertical, 8)
	view.root.SetMarginStart(12)
	view.root.SetMarginEnd(12)
	view.root.SetMarginTop(12)
	view.root.SetMarginBottom(12)

	header := gtk.NewLabel("")
	header.SetMarkup("<b>Import &amp; Upgrade</b>")
	header.SetXAlign(0.0)
	header.SetMarginBottom(2)
	view.root.Append(header.AsPtr())

	description := gtk.NewLabel("Import pacman database and packages from external storage, then upgrade the offline machine.")
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

	import_button := gtk.NewButton("Import Packages")
	import_button.OnClick(func() {
		if onImport != nil {
			folder := view.Folder.Text()
			onImport(folder)
		}
	})
	buttons.Append(import_button.AsPtr())

	upgrade_button := gtk.NewButton("Upgrade Packages")
	upgrade_button.OnClick(func() {
		if onUpgrade != nil {
			folder := view.Folder.Text()
			onUpgrade(folder)
		}
	})
	buttons.Append(upgrade_button.AsPtr())

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

func (view *Import) ShowTerminal() {
	view.terminal_wrapper.SetVisible(true)
}

func (view *Import) ClearTerminal() {
	view.terminal.Clear()
}

func (view *Import) ScrollToBottom() {
	view.terminal.ScrollToBottom()
}

func (view *Import) RenderConsole(console *structs.Console) {
	RenderConsole(console, view.terminal)
}

func (view *Import) SetStatus(text string) {

	view.status.SetMarkup(text)

	if text != "" {
		view.status.SetVisible(true)
	} else {
		view.status.SetVisible(false)
	}

}
