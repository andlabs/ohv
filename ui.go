// 5 october 2014
package main

import (
	"os"
	"io/ioutil"
	"net/url"
	"unsafe"
)

// #cgo CFLAGS: -mmacosx-version-min=10.7 -DMACOSX_DEPLOYMENT_TARGET=10.7
// #cgo LDFLAGS: -mmacosx-version-min=10.7 -lobjc -framework Foundation -framework AppKit -framework WebKit
// #include "cocoa_darwin.h"
import "C"

type Window struct {
	w		C.goid
	sf		C.goid
	navtree	C.goid
	browser	C.goid
	current	Topic
}

var goids = make(map[C.goid]*Window)

func NewWindow() *Window {
	w := new(MainWindow)
	w.w = C.newWindow()
	w.sf = C.newSearchField()
	w.navtree = C.newNavTree()
	// TODO
	w.browser = C.newSearchField()

	goids[w.w] = w
	goids[w.sf] = w
	goids[w.navtree] = w
	goids[w.browser] = w

	C.navtreeBegin(w.navtree)

	return w
}

func (w *Window) Show() {
	C.windowShow(w.w)
}

//export uiOnWindowClosing
func uiOnWindowClosing(win C.goid) C.int {
	C.stopUIThread()
	return 1
}

/* TODO
//export mainWindowDoNavigateTo
func mainWindowDoNavigateTo(data unsafe.Pointer, model *C.GtkTreeModel, iter *C.GtkTreeIter) {
	m := (*MainWindow)(data)
	t := navtreeTopic(C.gtk_tree_model_get_path(model, iter))
	m.current = t
	prepared, err := m.current.Prepare()
	if err != nil {
		// TODO
		println(err)
		return
	}
	// path must be a valid URI
	// TODO urlencode
	s := "file://" + prepared.Path
	cs := (*C.gchar)(unsafe.Pointer(C.CString(s)))
	defer C.free(unsafe.Pointer(cs))
	group := C.webkit_web_view_get_group(m.mw.browser)
	C.webkit_web_view_group_remove_all_user_style_sheets(group)
	if prepared.CSSPath != "" {
		// sigh, this can't be loaded from a file...
		f, err := os.Open(prepared.CSSPath)
		if err != nil {
			// TODO
			panic(err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			// TODO
			panic(err)
		}
		ccss := (*C.gchar)(unsafe.Pointer(C.CString(string(b))))
		defer C.free(unsafe.Pointer(ccss))
		cbasedir := (*C.gchar)(unsafe.Pointer(C.CString(string(prepared.CSSBaseDir))))
		defer C.free(unsafe.Pointer(cbasedir))
		C.webkit_web_view_group_add_user_style_sheet(group, ccss, cbasedir, nil, nil, C.WEBKIT_INJECTED_CONTENT_FRAMES_ALL)
	}
	C.webkit_web_view_load_uri(m.mw.browser, cs)
}
*/

/* TODO
//export mainWindowDoFollowLink
func mainWindowDoFollowLink(data unsafe.Pointer, curl *C.char) {
	m := (*MainWindow)(data)
	u, err := url.Parse(C.GoString(curl))
	if err != nil {
		// TODO
		panic(err)
	}
	t := m.current.Source().Lookup(u)
	// TODO change this call to return the last index separately for the below
	path := navtreePathTo(t)
	defer C.gtk_tree_path_free(path)
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
*/
