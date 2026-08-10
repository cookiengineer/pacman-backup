package gtk

/*
#include <gtk/gtk.h>
*/
import "C"
import "unsafe"

type Label struct {
	Widget
}

func NewLabel(text string) *Label {

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	ptr := C.gtk_label_new(cText)

	return &Label{
		Widget: Widget{widget: (*C.GtkWidget)(unsafe.Pointer(ptr))},
	}

}

func (l *Label) SetMarkup(markup string) {
	cMarkup := C.CString(markup)
	defer C.free(unsafe.Pointer(cMarkup))
	C.gtk_label_set_markup((*C.GtkLabel)(unsafe.Pointer(l.widget)), cMarkup)
}

func (l *Label) SetWrap(wrap bool) {
	if wrap {
		C.gtk_label_set_wrap((*C.GtkLabel)(unsafe.Pointer(l.widget)), C.TRUE)
	} else {
		C.gtk_label_set_wrap((*C.GtkLabel)(unsafe.Pointer(l.widget)), C.FALSE)
	}
}

func (l *Label) SetXAlign(align float32) {
	C.gtk_label_set_xalign((*C.GtkLabel)(unsafe.Pointer(l.widget)), C.float(align))
}
