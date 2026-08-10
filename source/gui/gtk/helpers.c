#include <gtk/gtk.h>
#include <stdlib.h>

extern void goSignalCallback(int id);
extern gboolean goMainCallback(int id);
extern void goFileDialogCallback(char* path, int id);

gpointer intToPointer(int id) {
	return GINT_TO_POINTER(id);
}

int pointerToInt(gpointer ptr) {
	return GPOINTER_TO_INT(ptr);
}

void cSignalCallback(GtkWidget *widget, gpointer data) {
	int id = GPOINTER_TO_INT(data);
	goSignalCallback(id);
}

gboolean cMainCallback(gpointer data) {
	int id = GPOINTER_TO_INT(data);
	return goMainCallback(id);
}

void cListBoxRowActivated(GtkListBox *box, GtkListBoxRow *row, gpointer data) {
	int id = GPOINTER_TO_INT(data);
	goSignalCallback(id);
}

void cFileDialogCallback(GObject *source, GAsyncResult *result, gpointer data) {
	int id = GPOINTER_TO_INT(data);
	GError *error = NULL;
	GFile *file = gtk_file_dialog_select_folder_finish(GTK_FILE_DIALOG(source), result, &error);

	if (file != NULL) {
		char *path = g_file_get_path(file);
		goFileDialogCallback(path, id);
		g_free(path);
		g_object_unref(file);
	} else {
		goFileDialogCallback("", id);
		if (error != NULL) {
			g_error_free(error);
		}
	}
}

const char* gtk_editable_get_text_wrapper(GtkEditable *editable) {
	return gtk_editable_get_text(editable);
}

void gtk_editable_set_text_wrapper(GtkEditable *editable, const char *text) {
	gtk_editable_set_text(editable, text);
}
