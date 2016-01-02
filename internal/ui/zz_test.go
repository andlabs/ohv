// 2 january 2016
package ui

import (
	"testing"
)

func TestIt(t *testing.T) {
	err := Main(func() {
		w := NewWindow("Test", 320, 240)
		w.OnClosing(func() bool {
			Quit()
			return true
		})
		w.Move(100, 100)
		w.Show()
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
