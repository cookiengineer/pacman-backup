package gtk

/*
#include <gtk/gtk.h>
*/
import "C"
import "unsafe"

const (
	PolicyNever     = 0
	PolicyAutomatic = 1
	PolicyAlways    = 2
)

type ScrolledWindow struct {
	Widget
}

func NewScrolledWindow() *ScrolledWindow {

	ptr := C.gtk_scrolled_window_new()

	return &ScrolledWindow{
		Widget: Widget{widget: (*C.GtkWidget)(unsafe.Pointer(ptr))},
	}

}

func (s *ScrolledWindow) SetChild(child unsafe.Pointer) {
	C.gtk_scrolled_window_set_child((*C.GtkScrolledWindow)(unsafe.Pointer(s.widget)), (*C.GtkWidget)(child))
}

func (s *ScrolledWindow) SetPolicy(hscroll, vscroll int) {
	C.gtk_scrolled_window_set_policy(
		(*C.GtkScrolledWindow)(unsafe.Pointer(s.widget)),
		C.GtkPolicyType(hscroll),
		C.GtkPolicyType(vscroll),
	)
}
