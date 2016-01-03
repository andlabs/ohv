// 5 october 2014
package main

import (
	"github.com/andlabs/ohv/internal/ui"
)

func main() {
	LoadLibraries()
	err := ui.Main(func() {
		NewWindow()
	})
	if err != nil {
		panic(err)
	}
}
