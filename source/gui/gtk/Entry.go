package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>

extern const char* gtk_editable_get_text_wrapper(GtkEditable* editable);
extern void gtk_editable_set_text_wrapper(GtkEditable* editable, const char* text);
*/
import "C"
import "unsafe"

type Entry struct {
	Widget
}

func NewEntry() *Entry {

	ptr := C.gtk_entry_new()

	return &Entry{
		Widget: Widget{widget: (*C.GtkWidget)(unsafe.Pointer(ptr))},
	}

}

func (e *Entry) Text() string {
	cText := C.gtk_editable_get_text_wrapper((*C.GtkEditable)(unsafe.Pointer(e.widget)))
	return C.GoString(cText)
}

func (e *Entry) SetText(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.gtk_editable_set_text_wrapper((*C.GtkEditable)(unsafe.Pointer(e.widget)), cText)
}

func (e *Entry) SetPlaceholder(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.gtk_entry_set_placeholder_text((*C.GtkEntry)(unsafe.Pointer(e.widget)), cText)
}

func (e *Entry) SetVisibility(visible bool) {
	if visible {
		C.gtk_entry_set_visibility((*C.GtkEntry)(unsafe.Pointer(e.widget)), C.TRUE)
	} else {
		C.gtk_entry_set_visibility((*C.GtkEntry)(unsafe.Pointer(e.widget)), C.FALSE)
	}
}
