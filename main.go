// 5 october 2014
package main

import (
	"runtime"
)

// #include "gtk_unix.h"
import "C"

var m *MainWindow		// keep on heap

func main() {
	LoadLibraries()
	quit := make(chan struct{})
	go func() {
		runtime.LockOSThread()
		C.gtk_init(nil, nil)
		m = NewMainWindow()
		m.Show()
		C.gtk_main()
		quit <- struct{}{}
	}()
	<-quit
}
