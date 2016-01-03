// 5 october 2014
package main

import (
/* TODO
	"os"
	"io/ioutil"
	"net/url"
	"unsafe"
*/

	"github.com/andlabs/ohv/internal/ui"
)

type Window struct {
	w		*ui.Window
	search	*ui.SearchField
	navtree	*ui.Tree
	page		*ui.WebView
	current	Topic
}

func NewWindow() *Window {
	w := new(Window)
	w.w = ui.NewWindow("ohv", xxx, xxx)
	w.search = ui.NewSearchField()
	w.navtree = ui.NewNavtree()
	w.page = ui.NewWebView()

	w.w.OnClosing(w.onClosing)

	margin := ui.NewMargin(w.search)
	box := ui.NewBox(margin, w.navtree)
	splitter := ui.NewSplitter(box, w.page)
	w.SetChild(splitter)
	splitter.SetPosition(xxx)

	w.navtree.SetModel(LibraryModel)

	w.w.Show()
	w.navtree.OnSelected(w.navigate)

	return w
}

func (w *Window) onClosing() bool {
	ui.Quit()
	return true
}

func (w *Window) navigate() {
	node := w.navtree.Selected()
	if node == nil {
		return
	}
	m.current = node.(Topic)
	prepared, err := m.current.Prepare()
	if err != nil {
		// TODO
		println(err)
		return
	}
/* TODO
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
*/
	w.page.NavigateFile(prepared.Path)
}

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
