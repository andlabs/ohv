// 5 october 2014
package main

import (
	"runtime"
)

// #include "cocoa_darwin.h"
import "C"

func main() {
	LoadLibraries()
	quit := make(chan struct{})
	go func() {
		runtime.LockOSThread()
		C.initUIThread()
		w := NewWindow()
		w.Show()
		C.runUIThread()
		quit <- struct{}{}
	}()
	<-quit
}
