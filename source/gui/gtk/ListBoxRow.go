package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type ListBoxRow struct {
	Widget
}

func NewListBoxRow() *ListBoxRow {

	ptr := C.gtk_list_box_row_new()

	row := &ListBoxRow{
		Widget: Widget{
			widget: (*C.GtkWidget)(unsafe.Pointer(ptr)),
		},
	}

	return row

}

func (r *ListBoxRow) SetText(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	label := C.gtk_label_new(cText)

	// Make the label fill the row horizontally.
	C.gtk_widget_set_halign(
		label,
		C.GtkAlign(C.GTK_ALIGN_FILL),
	)

	// Align the text itself to the left.
	C.gtk_label_set_xalign(
		(*C.GtkLabel)(unsafe.Pointer(label)),
		C.gfloat(0.0),
	)

	C.gtk_widget_set_visible(label, C.TRUE)

	C.gtk_list_box_row_set_child(
		(*C.GtkListBoxRow)(unsafe.Pointer(r.widget)),
		label,
	)
}

func (r *ListBoxRow) SetChild(child *Widget) {
	C.gtk_list_box_row_set_child(
		(*C.GtkListBoxRow)(unsafe.Pointer(r.widget)),
		child.widget,
	)
}

func (r *ListBoxRow) Child() *Widget {
	child := C.gtk_list_box_row_get_child(
		(*C.GtkListBoxRow)(unsafe.Pointer(r.widget)),
	)

	if child == nil {
		return nil
	}

	return &Widget{
		widget: (*C.GtkWidget)(unsafe.Pointer(child)),
	}
}
