// 5 october 2014
package main

import (
	"unsafe"
	"github.com/reusee/ggir/gtk"
	"github.com/andlabs/ohv/webkit2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: --std=c99
// #include "gtk_unix.h"
import "C"

type MainWindow struct {
	window	*gtk.Window
	paned	*gtk.Paned
	navtree	*gtk.TreeView
	navscroll	*gtk.ScrolledWindow
	browser	*webkit2.WebView
}

func NewMainWindow() *MainWindow {
	m := new(MainWindow)

	m.window = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	m.window.SetTitle("ohv")
	m.window.Connect("destroy", gtk.MainQuit)

	m.paned = gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	m.window.Add(m.paned);

	m.navtree = gtk.TreeViewNew()
	C.newModel(m.navtree.CPointer)
	// see ggir bug #1
	m.navscroll = gtk.NewScrolledWindowFromCPointer(unsafe.Pointer(C.gtk_scrolled_window_new(nil, nil)))
	m.navscroll.SetShadowType(gtk.SHADOW_IN)
	m.navscroll.Add(m.navtree)
	m.paned.Add1(m.navscroll)

	m.browser = webkit2.WebViewNew()
	m.paned.Add2(m.browser)

	return m
}
