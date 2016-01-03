// 2 january 2016
package ui

import (
	"net/url"
)

// #include "ui.h"
import "C"

type WebView struct {
	id	C.id
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
