// +build darwin

package gtkosx

/*
#cgo pkg-config: atk,gtk-mac-integration-gtk3
#include <gtkosxapplication.h>

GtkosxApplication * newOSXApplication() {
    return g_object_new(GTKOSX_TYPE_APPLICATION, NULL);
}

void set_menu_bar(GtkosxApplication *app, void *menu) {
	gtkosx_application_set_menu_bar(app, GTK_MENU_SHELL(menu));
}

void insert_app_menu_item(GtkosxApplication *app, void *menuitem, gint index) {
	gtkosx_application_insert_app_menu_item(app, GTK_WIDGET(menuitem), index);
}
*/
import "C"
import (
	"unsafe"

	"github.com/geoffholden/gotk3/gtk"
)

type GTKOSXApplication struct {
	ptr *C.GtkosxApplication
}

func Init() GTKOSXApplication {
	var ret GTKOSXApplication
	ret.ptr = C.newOSXApplication()
	return ret
}

func (app *GTKOSXApplication) SetMenuBar(menubar *gtk.MenuBar) {
	C.set_menu_bar(app.ptr, unsafe.Pointer(menubar.Native()))
}

func (app *GTKOSXApplication) InsertAppMenuItem(item *gtk.MenuItem, index int) {
	C.insert_app_menu_item(app.ptr, unsafe.Pointer(item.Native()), C.gint(index))
}

func (app *GTKOSXApplication) Ready() {
	C.gtkosx_application_ready(app.ptr)
}
