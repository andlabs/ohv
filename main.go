// 5 october 2014
package main

// TODO menus needed for keyboard shortcuts to work

import (
	"github.com/andlabs/ohv/internal/ui"
)

func main() {
	LoadLibraries()
	err := ui.Main(func() {
		Library.MakeModel()
		NewWindow()
	})
	if err != nil {
		panic(err)
	}
}
