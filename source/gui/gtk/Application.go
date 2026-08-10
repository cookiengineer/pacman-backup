package gtk

/*
#include <gtk/gtk.h>
*/
import "C"
import "unsafe"

type Application struct {
	app *C.GtkApplication
}

func NewApplication(id string) *Application {

	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cId))

	app := C.gtk_application_new(cId, C.G_APPLICATION_DEFAULT_FLAGS)

	return &Application{
		app: (*C.GtkApplication)(unsafe.Pointer(app)),
	}

}

func (a *Application) Ptr() *C.GtkApplication {
	return a.app
}

func (a *Application) Run() int {
	return int(C.g_application_run((*C.GApplication)(unsafe.Pointer(a.app)), 0, nil))
}

func (a *Application) OnActivate(fn func()) {
	connectSignal(unsafe.Pointer(a.app), "activate", fn)
}
