// 5 october 2014
package main

import (
	"unsafe"
)

// #cgo pkg-config: gtk+-3.0 webkit2gtk-3.0
// #cgo CFLAGS: --std=c99
// #include "gtk_unix.h"
import "C"

type MainWindow struct {
	mw		*C.MainWindow
}

func NewMainWindow() *MainWindow {
	m := new(MainWindow)
	m.mw = C.newMainWindow(unsafe.Pointer(m))
	return m
/*
	m.navsel.Connect("changed", func () {
		xmodel, xiter, selected := m.navsel.GetSelected()
		if !selected {
			return
		}
		model := (*C.GtkTreeModel)(unsafe.Pointer(xmodel))
		iter := *(*C.GtkTreeIter)(unsafe.Pointer(&xiter))
		t := navtreeTopic(C.gtk_tree_model_get_path(model, &iter))
		s, err := t.Prepare()
		if err != nil {
			// TODO
			println(err)
			return
		}
		m.browser.LoadHtml(s, "")
	})
*/
}

//export mainWindowNavigateTo
func mainWindowNavigateTo(sel *C.GtkTreeSelection, data unsafe.Pointer) {
	// ...
}
