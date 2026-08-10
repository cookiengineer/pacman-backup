package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>

extern const char* gtk_editable_get_text_wrapper(GtkEditable* editable);
*/
import "C"
import "unsafe"

func ShowSudoDialog(parent unsafe.Pointer, onResult func(string)) {

	window := C.gtk_window_new()
	widget := (*C.GtkWidget)(unsafe.Pointer(window))

	cTitle := C.CString("Sudo Password")
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_window_set_title((*C.GtkWindow)(unsafe.Pointer(window)), cTitle)

	C.gtk_window_set_modal((*C.GtkWindow)(unsafe.Pointer(window)), C.TRUE)
	C.gtk_window_set_default_size((*C.GtkWindow)(unsafe.Pointer(window)), 400, 150)

	if parent != nil {
		C.gtk_window_set_transient_for((*C.GtkWindow)(unsafe.Pointer(window)), (*C.GtkWindow)(parent))
	}

	box := C.gtk_box_new(C.GTK_ORIENTATION_VERTICAL, 8)
	C.gtk_widget_set_margin_start(box, 12)
	C.gtk_widget_set_margin_end(box, 12)
	C.gtk_widget_set_margin_top(box, 12)
	C.gtk_widget_set_margin_bottom(box, 12)
	C.gtk_window_set_child((*C.GtkWindow)(unsafe.Pointer(window)), box)

	cPrompt := C.CString("Enter sudo password:")
	defer C.free(unsafe.Pointer(cPrompt))
	label := C.gtk_label_new(cPrompt)
	C.gtk_label_set_xalign((*C.GtkLabel)(unsafe.Pointer(label)), 0.0)
	C.gtk_widget_set_visible(label, C.TRUE)
	C.gtk_box_append((*C.GtkBox)(unsafe.Pointer(box)), label)

	entry := C.gtk_entry_new()
	C.gtk_entry_set_visibility((*C.GtkEntry)(unsafe.Pointer(entry)), C.FALSE)
	cPlaceholder := C.CString("password")
	defer C.free(unsafe.Pointer(cPlaceholder))
	C.gtk_entry_set_placeholder_text((*C.GtkEntry)(unsafe.Pointer(entry)), cPlaceholder)
	C.gtk_widget_set_visible(entry, C.TRUE)
	C.gtk_box_append((*C.GtkBox)(unsafe.Pointer(box)), entry)

	btnBox := C.gtk_box_new(C.GTK_ORIENTATION_HORIZONTAL, 8)
	C.gtk_widget_set_halign(btnBox, C.GTK_ALIGN_END)
	C.gtk_widget_set_visible(btnBox, C.TRUE)
	C.gtk_box_append((*C.GtkBox)(unsafe.Pointer(box)), btnBox)

	cCancel := C.CString("Cancel")
	defer C.free(unsafe.Pointer(cCancel))
	cancelBtn := C.gtk_button_new_with_label(cCancel)
	C.gtk_widget_set_visible(cancelBtn, C.TRUE)
	C.gtk_box_append((*C.GtkBox)(unsafe.Pointer(btnBox)), cancelBtn)

	cOk := C.CString("OK")
	defer C.free(unsafe.Pointer(cOk))
	okBtn := C.gtk_button_new_with_label(cOk)
	C.gtk_widget_set_visible(okBtn, C.TRUE)
	C.gtk_box_append((*C.GtkBox)(unsafe.Pointer(btnBox)), okBtn)

	var entryPtr *C.GtkWidget
	entryPtr = entry
	var windowPtr *C.GtkWidget
	windowPtr = widget

	connectSignal(unsafe.Pointer(cancelBtn), "clicked", func() {
		C.gtk_window_destroy((*C.GtkWindow)(unsafe.Pointer(windowPtr)))
		onResult("")
	})

	connectSignal(unsafe.Pointer(okBtn), "clicked", func() {
		text := C.GoString(C.gtk_editable_get_text_wrapper((*C.GtkEditable)(unsafe.Pointer(entryPtr))))
		C.gtk_window_destroy((*C.GtkWindow)(unsafe.Pointer(windowPtr)))
		onResult(text)
	})

	C.gtk_window_present((*C.GtkWindow)(unsafe.Pointer(window)))

}
