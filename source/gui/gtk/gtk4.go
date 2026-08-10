package gtk

/*
#cgo pkg-config: gtk4

#include <gtk/gtk.h>
#include <stdlib.h>

extern void goSignalCallback(int id);
extern void goMainCallback(int id);

extern gpointer intToPointer(int id);
extern int pointerToInt(gpointer ptr);

extern void cSignalCallback(GtkWidget *widget, gpointer data);
extern gboolean cMainCallback(gpointer data);
extern void cListBoxRowActivated(GtkListBox *box, GtkListBoxRow *row, gpointer data);
*/
import "C"
import "sync"
import "unsafe"

var (
	signalCallbacks sync.Map
	signalCounter   int
	mainCallbacks   sync.Map
	mainCounter     int
)

func registerCallback(fn func()) int {
	signalCounter++
	signalCallbacks.Store(signalCounter, fn)
	return signalCounter
}

func unregisterCallback(id int) {
	signalCallbacks.Delete(id)
}

//export goSignalCallback
func goSignalCallback(id C.int) {
	val, ok := signalCallbacks.Load(int(id))
	if ok {
		val.(func())()
	}
}

//export goMainCallback
func goMainCallback(id C.int) {
	val, ok := mainCallbacks.LoadAndDelete(int(id))
	if ok {
		val.(func())()
	}
}

func gpointerFromInt(id int) C.gpointer {
	return C.intToPointer(C.int(id))
}

func intFromGpointer(ptr C.gpointer) int {
	return int(C.pointerToInt(ptr))
}

func RunOnMain(fn func()) {

	mainCounter++
	id := mainCounter

	mainCallbacks.Store(id, fn)

	C.g_main_context_invoke(
		nil,
		(C.GSourceFunc)(unsafe.Pointer(C.cMainCallback)),
		gpointerFromInt(id),
	)

}

func connectSignal(widget unsafe.Pointer, signal string, fn func()) {

	cSignal := C.CString(signal)
	defer C.free(unsafe.Pointer(cSignal))

	id := registerCallback(fn)

	C.g_signal_connect_data(
		C.gpointer(widget),
		cSignal,
		C.GCallback(unsafe.Pointer(C.cSignalCallback)),
		gpointerFromInt(id),
		nil,
		0,
	)

}

func connectListBoxSignal(widget unsafe.Pointer, fn func()) {

	id := registerCallback(fn)
	cSignal := C.CString("row-activated")
	defer C.free(unsafe.Pointer(cSignal))

	C.g_signal_connect_data(
		C.gpointer(widget),
		cSignal,
		C.GCallback(unsafe.Pointer(C.cListBoxRowActivated)),
		gpointerFromInt(id),
		nil,
		0,
	)

}
