package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Widget struct {
	widget *C.GtkWidget
}

func (w *Widget) AsPtr() unsafe.Pointer {
	return unsafe.Pointer(w.widget)
}

func (w *Widget) SetMarginStart(margin int) {
	C.gtk_widget_set_margin_start(w.widget, C.int(margin))
}

func (w *Widget) SetMarginEnd(margin int) {
	C.gtk_widget_set_margin_end(w.widget, C.int(margin))
}

func (w *Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(w.widget, C.int(margin))
}

func (w *Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(w.widget, C.int(margin))
}

func (w *Widget) SetHAlign(align Align) {
	C.gtk_widget_set_halign(w.widget, C.GtkAlign(align))
}

func (w *Widget) SetHExpand(expand bool) {
	if expand {
		C.gtk_widget_set_hexpand(w.widget, C.TRUE)
	} else {
		C.gtk_widget_set_hexpand(w.widget, C.FALSE)
	}
}

func (w *Widget) SetVAlign(align Align) {
	C.gtk_widget_set_valign(w.widget, C.GtkAlign(align))
}

func (w *Widget) SetVExpand(expand bool) {
	if expand {
		C.gtk_widget_set_vexpand(w.widget, C.TRUE)
	} else {
		C.gtk_widget_set_vexpand(w.widget, C.FALSE)
	}
}

func (w *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(w.widget, C.int(width), C.int(height))
}

func (w *Widget) SetVisible(visible bool) {
	if visible {
		C.gtk_widget_set_visible(w.widget, C.TRUE)
	} else {
		C.gtk_widget_set_visible(w.widget, C.FALSE)
	}
}

func (w *Widget) SetSensitive(sensitive bool) {
	if sensitive {
		C.gtk_widget_set_sensitive(w.widget, C.TRUE)
	} else {
		C.gtk_widget_set_sensitive(w.widget, C.FALSE)
	}
}
