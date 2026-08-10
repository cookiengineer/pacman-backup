package views

import "pacman-backup/gui/gtk"
import "pacman-backup/structs"
import "unsafe"

type Update struct {
	Reviews          *gtk.ListBox
	root             *gtk.Box
	reviews_header   *gtk.Label
	reviews_wrapper  *gtk.ScrolledWindow
	terminal         *gtk.TextView
	terminal_wrapper *gtk.ScrolledWindow
}

func (view *Update) AsPtr() unsafe.Pointer {
	return view.root.AsPtr()
}

func NewUpdate(parent unsafe.Pointer, onSync func(), onDownload func()) *Update {

	view := &Update{}

	view.root = gtk.NewBox(gtk.OrientationVertical, 8)
	view.root.SetMarginStart(12)
	view.root.SetMarginEnd(12)
	view.root.SetMarginTop(12)
	view.root.SetMarginBottom(12)

	header := gtk.NewLabel("")
	header.SetMarkup("<b>System Updates</b>")
	header.SetXAlign(0.0)
	header.SetMarginBottom(2)
	view.root.Append(header.AsPtr())

	description := gtk.NewLabel("Synchronize package databases and download all packages from online mirrors.")
	description.SetWrap(true)
	description.SetXAlign(0.0)
	description.SetMarginBottom(12)
	view.root.Append(description.AsPtr())

	buttons := gtk.NewBox(gtk.OrientationHorizontal, 8)
	buttons.SetMarginBottom(6)
	view.root.Append(buttons.AsPtr())

	sync_button := gtk.NewButton("Sync Databases")
	sync_button.OnClick(func() {
		if onSync != nil {
			onSync()
		}
	})
	buttons.Append(sync_button.AsPtr())

	download_button := gtk.NewButton("Download Packages")
	download_button.OnClick(func() {
		if onDownload != nil {
			onDownload()
		}
	})
	buttons.Append(download_button.AsPtr())

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
	view.reviews_header.SetMarkup("<b>Configuration files to review (.pacnew / .pacsave):</b>")
	view.reviews_header.SetXAlign(0.0)
	view.reviews_header.SetMarginBottom(4)
	view.reviews_header.SetVisible(false)
	view.root.Append(view.reviews_header.AsPtr())

	view.reviews_wrapper = gtk.NewScrolledWindow()
	view.reviews_wrapper.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)
	view.reviews_wrapper.SetVExpand(true)
	view.reviews_wrapper.SetSizeRequest(-1, 120)
	view.reviews_wrapper.SetVisible(false)
	view.root.Append(view.reviews_wrapper.AsPtr())

	view.Reviews = gtk.NewListBox()
	view.reviews_wrapper.SetChild(view.Reviews.AsPtr())

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
}

func (view *Update) ScrollToBottom() {
	view.terminal.ScrollToBottom()
}

func (view *Update) RenderConsole(console *structs.Console) {
	RenderConsole(console, view.terminal)
}

func (view *Update) SetPacnewFiles(files []string) {

	view.Reviews.Clear()

	for _, file := range files {
		view.Reviews.Append(file)
	}

	if len(files) > 0 {
		view.reviews_header.SetVisible(true)
		view.reviews_wrapper.SetVisible(true)
	} else {
		view.reviews_header.SetVisible(false)
		view.reviews_wrapper.SetVisible(false)
	}

}
