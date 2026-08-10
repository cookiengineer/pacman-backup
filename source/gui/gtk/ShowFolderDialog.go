package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>

extern void goFileDialogCallback(char* path, int id);
extern void cFileDialogCallback(GObject* source, GAsyncResult* result, gpointer data);
extern gpointer intToPointer(int id);
*/
import "C"
import "sync"
import "unsafe"

var file_dialog_callbacks sync.Map
var file_dialog_counter int

//export goFileDialogCallback
func goFileDialogCallback(path *C.char, id C.int) {
	val, ok := file_dialog_callbacks.LoadAndDelete(int(id))
	if ok {
		val.(func(string))(C.GoString(path))
	}
}

func ShowFolderDialog(parent unsafe.Pointer, onResult func(string)) {

	dialog := C.gtk_file_dialog_new()

	cTitle := C.CString("Select Folder")
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_file_dialog_set_title((*C.GtkFileDialog)(unsafe.Pointer(dialog)), cTitle)

	file_dialog_counter++
	id := file_dialog_counter
	file_dialog_callbacks.Store(id, onResult)

	var parentWin *C.GtkWindow
	if parent != nil {
		parentWin = (*C.GtkWindow)(parent)
	} else {
		parentWin = nil
	}

	C.gtk_file_dialog_select_folder(
		(*C.GtkFileDialog)(unsafe.Pointer(dialog)),
		parentWin,
		nil,
		C.GAsyncReadyCallback(unsafe.Pointer(C.cFileDialogCallback)),
		gpointerFromInt(id),
	)

	C.g_object_unref(C.gpointer(unsafe.Pointer(dialog)))

}
