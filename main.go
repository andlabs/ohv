// 5 october 2014
package main

import (
	"runtime"
)

// #include "cocoa_darwin.h"
import "C"

var m *MainWindow		// keep on heap (TODO really needed?)

func main() {
	LoadLibraries()
	quit := make(chan struct{})
	go func() {
		runtime.LockOSThread()
		C.initUIThread()
		m = NewMainWindow()
		m.Show()
		C.runUIThread()
		quit <- struct{}{}
	}()
	<-quit
}
