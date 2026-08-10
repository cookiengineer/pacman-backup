package views

import "pacman-backup/gui/gtk"
import "pacman-backup/structs"
import "unsafe"

type Update struct {
	Reviews          *gtk.ListBox
	Mirrors          *gtk.DropDown
	mirrors_list     []string
	root             *gtk.Box
	reviews_header   *gtk.Label
	terminal         *gtk.TextView
	terminal_length  int
	terminal_wrapper *gtk.ScrolledWindow
}

func (view *Update) AsPtr() unsafe.Pointer {
	return view.root.AsPtr()
}

func NewUpdate(parent unsafe.Pointer, mirrors []string, onSync func()) *Update {

	view := &Update{
		mirrors_list: mirrors,
	}

	view.root = gtk.NewBox(gtk.OrientationVertical, 8)
	view.root.SetMarginStart(12)
	view.root.SetMarginEnd(12)
	view.root.SetMarginTop(12)
	view.root.SetMarginBottom(12)

	header := gtk.NewLabel("")
	header.SetMarkup("<b>System Update</b>")
	header.SetXAlign(0.0)
	header.SetMarginBottom(2)
	view.root.Append(header.AsPtr())

	description := gtk.NewLabel("Synchronize package databases and download all packages from online mirrors.")
	description.SetWrap(true)
	description.SetXAlign(0.0)
	description.SetMarginBottom(12)
	view.root.Append(description.AsPtr())

	view.Mirrors = gtk.NewDropDown(mirrors)
	view.Mirrors.SetMarginBottom(12)
	view.root.Append(view.Mirrors.AsPtr())

	buttons := gtk.NewBox(gtk.OrientationHorizontal, 8)
	buttons.SetMarginBottom(6)
	view.root.Append(buttons.AsPtr())

	sync_button := gtk.NewButton("Download Databases and Packages")
	sync_button.OnClick(func() {
		if onSync != nil {
			onSync()
		}
	})
	buttons.Append(sync_button.AsPtr())

	view.terminal_wrapper = gtk.NewScrolledWindow()
	view.terminal_wrapper.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)
	view.terminal_wrapper.SetVExpand(true)
	view.terminal_wrapper.SetSizeRequest(-1, 300)
	view.terminal_wrapper.SetMarginBottom(8)
	view.terminal_wrapper.SetVisible(false)
	view.root.Append(view.terminal_wrapper.AsPtr())

	view.terminal = gtk.NewTextView()
	view.terminal.SetEditable(false)
	view.terminal.SetMonospace(true)
	view.terminal.SetCursorVisible(false)
	view.terminal.SetTerminalStyle()
	view.terminal_wrapper.SetChild(view.terminal.AsPtr())

	view.reviews_header = gtk.NewLabel("")
	view.reviews_header.SetMarkup("<b>Please review these files:</b>")
	view.reviews_header.SetXAlign(0.0)
	view.reviews_header.SetMarginBottom(4)
	view.reviews_header.SetVisible(false)
	view.root.Append(view.reviews_header.AsPtr())

	view.Reviews = gtk.NewListBox()
	view.root.Append(view.Reviews.AsPtr())

	return view

}

func (view *Update) ShowTerminal() {
	view.terminal_wrapper.SetVisible(true)
}

func (view *Update) AppendTerminal(text string) {
	view.terminal.Append(text)
}

func (view *Update) ClearTerminal() {
	view.terminal.Clear()
	view.terminal_length = 0
}

func (view *Update) ScrollToBottom() {
	view.terminal.ScrollToBottom()
}

func (view *Update) RenderTerminal(console *structs.Console) {

	if console.Length() > view.terminal_length {
		RenderConsole(console, view.terminal, view.terminal_length)
		view.terminal_length = console.Length()
	}

}

func (view *Update) SetPacnewFiles(files []string) {

	view.Reviews.Clear()

	for _, file := range files {

		row   := gtk.NewListBoxRow()
		label := gtk.NewLabel(file)
		label.SetXAlign(0.0)
		label.SetHAlign(gtk.AlignFill)

		row.SetChild(&label.Widget)
		view.Reviews.AppendRow(row)

	}

	if len(files) > 0 {
		view.reviews_header.SetVisible(true)
		view.Reviews.SetVisible(true)
	} else {
		view.reviews_header.SetVisible(false)
		view.Reviews.SetVisible(false)
	}

}

func (view *Update) GetSelectedMirror() string {

	idx := view.Mirrors.GetSelected()

	if int(idx) < len(view.mirrors_list) {
		return view.mirrors_list[idx]
	}

	return ""

}
