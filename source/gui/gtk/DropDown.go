package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>

static GtkWidget* _dropdown_new(char** strings) {
	return gtk_drop_down_new_from_strings((const char* const*)strings);
}
*/
import "C"

import "unsafe"

type DropDown struct {
	Widget
}

func NewDropDown(strings []string) *DropDown {

	cStrings := make([]*C.char, len(strings)+1)

	for i, s := range strings {
		cStrings[i] = C.CString(s)
	}

	cStrings[len(strings)] = nil

	ptr := C._dropdown_new(&cStrings[0])

	for i := 0; i < len(strings); i++ {
		C.free(unsafe.Pointer(cStrings[i]))
	}

	return &DropDown{
		Widget: Widget{widget: (*C.GtkWidget)(unsafe.Pointer(ptr))},
	}

}

func (d *DropDown) SetSelected(index uint) {
	C.gtk_drop_down_set_selected((*C.GtkDropDown)(unsafe.Pointer(d.widget)), C.guint(index))
}

func (d *DropDown) GetSelected() uint {
	return uint(C.gtk_drop_down_get_selected((*C.GtkDropDown)(unsafe.Pointer(d.widget))))
}

func (d *DropDown) OnChanged(fn func()) {
	connectSignal(unsafe.Pointer(d.widget), "notify::selected", fn)
}
