package gtk

/*
#include <gtk/gtk.h>
*/
import "C"
import "unsafe"

type Notebook struct {
	widget *C.GtkWidget
}

func NewNotebook() *Notebook {

	ptr := C.gtk_notebook_new()

	return &Notebook{
		widget: (*C.GtkWidget)(unsafe.Pointer(ptr)),
	}

}

func (n *Notebook) AsPtr() unsafe.Pointer {
	return unsafe.Pointer(n.widget)
}

func (n *Notebook) AppendPage(child unsafe.Pointer, tabLabel unsafe.Pointer) int {
	return int(C.gtk_notebook_append_page(
		(*C.GtkNotebook)(unsafe.Pointer(n.widget)),
		(*C.GtkWidget)(child),
		(*C.GtkWidget)(tabLabel),
	))
}
