// 5 october 2014
package main

import (
/* TODO
	"os"
	"io/ioutil"
*/
	"net/url"

	"github.com/andlabs/ohv/internal/ui"
)

type Window struct {
	w		*ui.Window
	search	*ui.SearchEntry
	navtree	*ui.Tree
	page		*ui.WebView
	current	Topic
	results	*SearchResults
}

// the default sizes and positions here are from my devhelp config

func NewWindow() *Window {
	w := new(Window)
	w.w = ui.NewWindow("ohv", 1095, 546)
	w.search = ui.NewSearchEntry()
	w.navtree = ui.NewTree()
	w.page = ui.NewWebView()

	margin := ui.NewMargin(w.search)
	box := ui.NewBox(margin, w.navtree)
	splitter := ui.NewSplitter(box, w.page)
	w.w.SetChild(splitter)
	splitter.SetPosition(250)
//	w.w.Move(100, 100)
	w.w.Center()

	w.navtree.SetModel(Library.Model())

	ui.OnShouldQuit(func() bool {
		return true
	})
	w.w.OnClosing(w.onClosing)
	w.search.OnChanged(w.searching)
	w.navtree.OnSelected(w.navigate)
	w.page.OnLoadFailed(w.loadFailed)
	w.page.OnLinkClicked(w.linkClicked)

	w.w.Show()

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
	if w.results != nil {		// searching
		w.current = w.results.Topic(node)
	} else {				// navigating
		w.current = node.(Topic)
	}
	prepared, err := w.current.Prepare()
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
	w.page.NavigateFile(prepared.Path, prepared.Anchor)
}

func (w *Window) linkClicked(target *url.URL) bool {
	if w.results != nil {		// abandon search on link clicked
		// don't restore the current since the current is about to change
		w.dismissSearch(false)
	}
	t := w.current.Source().Lookup(target)
	w.navtree.SetSelected(t)
	// TODO the above triggers a navigation for us
	return false
}

func (w *Window) loadFailed(sysError uintptr) {
	w.w.MsgBoxSysError(sysError)
}

func (w *Window) searching() {
	var err error

	text := w.search.Text()
	if text == "" {			// search cancelled
		// restore the current so that whatever the last search item was is now selected
		w.dismissSearch(true)
		return
	}
	if w.results != nil {		// overwrite existing search
		w.navtree.SetModel(nil)
		w.results.Dismiss()
	}
	w.results, err = LibrarySearcher.Search(text)
	if err != nil {
		// TODO
		return
	}
	w.navtree.SetModel(w.results.Model())
}

func (w *Window) dismissSearch(restoreCurrent bool) {
	// TODO NSSearchField can send "" twice
	if w.results == nil {
		return
	}
	w.navtree.SetModel(Library.Model())
	// TODO separate SetSelected from OnSelected so this won't need to be first
	w.results.Dismiss()
	w.results = nil
	if restoreCurrent && w.current != nil {
		w.navtree.SetSelected(w.current)
	}
}
