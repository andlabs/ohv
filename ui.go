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

//export mainWindowDoFollowLink
func mainWindowDoFollowLink(data unsafe.Pointer, curl *C.char) {
	m := (*MainWindow)(data)
	u, err := url.Parse(C.GoString(curl))
	if err != nil {
		// TODO
		panic(err)
	}
	t := m.current.Source().Lookup(u)
	path := C.gtk_tree_path_new()
	defer C.gtk_tree_path_free(path)
	for t != nil {
		p := t.Parent()
		children := Library
		if p != nil {
			children = p.Children()
		}
		i := C.gint(0)
		for _, c := range children {
			if c == t {
				break
			}
			i++
		}
		if i >= C.gint(len(children)) {
			panic("parent without known child in TODO()")
		}
		C.gtk_tree_path_prepend_index(path, i)
		t = p
	}
	// without the following line, the selection change won't work (this has always been the case :/ )
	C.gtk_tree_view_expand_to_path((*C.GtkTreeView)(unsafe.Pointer(m.mw.navtree)), path)
	// but don't expand the row itself
	C.gtk_tree_view_collapse_row((*C.GtkTreeView)(unsafe.Pointer(m.mw.navtree)), path)
	// and we need to scroll there too
	C.gtk_tree_view_scroll_to_cell((*C.GtkTreeView)(unsafe.Pointer(m.mw.navtree)), path, nil,
		C.FALSE, 0, 0)		// TODO change to TRUE?
	// and FINALLY make the change
	C.gtk_tree_selection_select_path(m.mw.navsel, path)
}
