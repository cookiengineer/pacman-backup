package gtk

/*
#include <gtk/gtk.h>
*/
import "C"
import "unsafe"

type TextView struct {
	Widget
}

func NewTextView() *TextView {
	ptr := C.gtk_text_view_new()

	return &TextView{
		Widget: Widget{widget: (*C.GtkWidget)(unsafe.Pointer(ptr))},
	}
}

func (t *TextView) SetEditable(editable bool) {
	if editable {
		C.gtk_text_view_set_editable((*C.GtkTextView)(unsafe.Pointer(t.widget)), C.TRUE)
	} else {
		C.gtk_text_view_set_editable((*C.GtkTextView)(unsafe.Pointer(t.widget)), C.FALSE)
	}
}

func (t *TextView) SetMonospace(monospace bool) {
	if monospace {
		C.gtk_text_view_set_monospace((*C.GtkTextView)(unsafe.Pointer(t.widget)), C.TRUE)
	} else {
		C.gtk_text_view_set_monospace((*C.GtkTextView)(unsafe.Pointer(t.widget)), C.FALSE)
	}
}

func (t *TextView) SetCursorVisible(visible bool) {
	if visible {
		C.gtk_text_view_set_cursor_visible((*C.GtkTextView)(unsafe.Pointer(t.widget)), C.TRUE)
	} else {
		C.gtk_text_view_set_cursor_visible((*C.GtkTextView)(unsafe.Pointer(t.widget)), C.FALSE)
	}
}

func (t *TextView) Append(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	buffer := C.gtk_text_view_get_buffer((*C.GtkTextView)(unsafe.Pointer(t.widget)))
	C.gtk_text_buffer_insert_at_cursor(buffer, cText, C.int(-1))
}

func (t *TextView) Clear() {
	buffer := C.gtk_text_view_get_buffer((*C.GtkTextView)(unsafe.Pointer(t.widget)))
	empty := C.CString("")
	defer C.free(unsafe.Pointer(empty))
	C.gtk_text_buffer_set_text(buffer, empty, C.int(-1))
}

func (t *TextView) ScrollToBottom() {
	buffer := C.gtk_text_view_get_buffer((*C.GtkTextView)(unsafe.Pointer(t.widget)))

	var iter C.GtkTextIter
	C.gtk_text_buffer_get_end_iter(buffer, &iter)

	endName := C.CString("end_scroll")
	defer C.free(unsafe.Pointer(endName))
	mark := C.gtk_text_buffer_create_mark(buffer, endName, &iter, C.FALSE)
	C.gtk_text_view_scroll_mark_onscreen((*C.GtkTextView)(unsafe.Pointer(t.widget)), mark)
}

func (t *TextView) SetTerminalStyle() {
	cssProvider := C.gtk_css_provider_new()

	css := C.CString(`
		textview {
			background-color: #0a0a0a;
			color: #c0c0c0;
			font-family: monospace;
			padding: 4px;
		}
		textview text {
			background-color: #0a0a0a;
			color: #c0c0c0;
		}
	`)
	defer C.free(unsafe.Pointer(css))

	C.gtk_css_provider_load_from_string((*C.GtkCssProvider)(unsafe.Pointer(cssProvider)), css)

	context := C.gtk_widget_get_style_context(t.widget)
	C.gtk_style_context_add_provider(
		context,
		(*C.GtkStyleProvider)(unsafe.Pointer(cssProvider)),
		C.GTK_STYLE_PROVIDER_PRIORITY_APPLICATION,
	)
}
