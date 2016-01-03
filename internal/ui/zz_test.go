// 2 january 2016
package ui

import (
	"net/url"
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
		wv := NewWebView()
		w.SetChild(wv)
		w.Show()
		QueueMain(func() {
			u, _ := url.Parse("http://google.com/")
			wv.Navigate(u)
		})
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
