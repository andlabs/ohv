// 2 january 2016
package ui

import (
//	"net/url"
	"testing"
)

func TestIt(t *testing.T) {
	err := Main(func() {
		w := NewWindow("Test", 640, 480)
		w.OnClosing(func() bool {
			Quit()
			return true
		})
		w.Move(100, 100)
		t := NewTree()
		w.SetChild(t)
		w.Show()
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
}

/*
func TestIt(t *testing.T) {
	err := Main(func() {
		w := NewWindow("Test", 640, 480)
		w.OnClosing(func() bool {
			Quit()
			return true
		})
		w.Move(100, 100)
		wv := NewWebView()
		wv.OnLoadFailed(func(sysError uintptr) {
			w.MsgBoxSysError(sysError)
		})
		w.SetChild(wv)
		w.Show()
		QueueMain(func() {
			u, _ := url.Parse("http://aphsrguphgpihpbtdnb.com/")
			wv.Navigate(u)
		})
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
*/
