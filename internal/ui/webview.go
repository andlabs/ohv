// 2 january 2016
package ui

import (
	"net/url"
)

// #include "ui.h"
import "C"

type WebView struct {
	id				C.id
	onLoadFailed		func(sysError uintptr)
	onLinkClicked		func(target *url.URL) bool
}

var webviews = make(map[C.id]*WebView)

func NewWebView() *WebView {
	w := new(WebView)
	w.id = C.newWebView()
	webviews[w.id] = w
	return w
}

func (w *WebView) Destroy() {
	delete(webviews, w.id)
	C.webViewDestroy(w.id)
}

func (w *WebView) Handle() uintptr {
	return touintptr(w.id)
}

func (w *WebView) Navigate(to *url.URL) {
	nsurl := fromURL(to)
	C.webViewNavigate(w.id, nsurl)
}

func (w *WebView) NavigateFile(to string) {
	nsurl := fromFileURL(to)
	C.webViewNavigate(w.id, nsurl)
}

func (w *WebView) OnLoadFailed(f func(sysError uintptr)) {
	w.onLoadFailed = f
}

//export doWebViewLoadFailed
func doWebViewLoadFailed(ww C.id, sysError C.id) {
	w := webviews[ww]
	if w.onLoadFailed != nil {
		w.onLoadFailed(touintptr(sysError))
	}
}

func (w *WebView) OnLinkClicked(f func(target *url.URL) bool) {
	w.onLinkClicked = f
}

//export doWebViewLinkClicked
func doWebViewLinkClicked(ww C.id, nsurl C.id) C.int {
	// TODO rename webViews
	w := webviews[ww]
	if w.onLinkClicked == nil {
		return 0
	}
	return frombool(w.onLinkClicked(toURL(nsurl)))
}
