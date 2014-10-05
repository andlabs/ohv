// 5 october 2014
package main

import (
	"github.com/reusee/ggir/gtk"
	"github.com/andlabs/ohv/webkit2"
)

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
// see ggir bug #1
/*	m.navscroll = gtk.ScrolledWindowNew(nil, nil)
	m.navscroll.SetShadowType(gtk.SHADOW_IN)
	m.navscroll.Add(m.navtree)
	m.paned.Add1(m.navscroll)
*/	m.paned.Add1(m.navtree)
	m.browser = webkit2.WebViewNew()
	m.paned.Add2(m.browser)

	return m
}

func main() {
	gtk.Init(nil, nil)
	NewMainWindow().window.ShowAll()
	gtk.Main()
}
