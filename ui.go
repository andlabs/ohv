// 5 october 2014
package main

import (
	"net/url"
	"unsafe"
)

// #cgo pkg-config: gtk+-3.0 webkit2gtk-3.0
// #cgo CFLAGS: --std=c99
// #include "gtk_unix.h"
import "C"

type MainWindow struct {
	mw		*C.MainWindow
	current	Topic
}

func NewMainWindow() *MainWindow {
	m := new(MainWindow)
	m.mw = C.newMainWindow(unsafe.Pointer(m))
	return m
}

func (m *MainWindow) Show() {
	C.gtk_widget_show_all(m.mw.window)
}

//export mainWindowDoNavigateTo
func mainWindowDoNavigateTo(data unsafe.Pointer, model *C.GtkTreeModel, iter *C.GtkTreeIter) {
	m := (*MainWindow)(data)
	t := navtreeTopic(C.gtk_tree_model_get_path(model, iter))
	m.current = t
	s, err := m.current.Prepare()
	if err != nil {
		// TODO
		println(err)
		return
	}
	cs := (*C.gchar)(unsafe.Pointer(C.CString(s)))
	defer C.free(unsafe.Pointer(cs))
	C.webkit_web_view_load_html(m.mw.browser, cs, nil)
}

//export mainWindowDoMSXHELP
func mainWindowDoMSXHELP(data unsafe.Pointer, curl *C.char) {
//	m := (*MainWindow)(data)
	u, err := url.Parse(C.GoString(curl))
	if err != nil {
		// TODO
		panic(err)
	}
	println(u.Query().Get("Id"))
}
