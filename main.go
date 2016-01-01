// 5 october 2014
package main

import (
	"runtime"
)

// #include "cocoa_darwin.h"
import "C"

func main() {
	runtime.LockOSThread()
	C.initUIThread()
	LoadLibraries()
	w := NewWindow()
	w.Show()
	C.runUIThread()
}
