// 2 january 2016
package ui

import (
	"fmt"
	"runtime"
	"sync"
)

// #cgo CFLAGS: -mmacosx-version-min=10.7 -DMACOSX_DEPLOYMENT_TARGET=10.7
// #cgo LDFLAGS: -mmacosx-version-min=10.7 -lobjc -framework Foundation -framework AppKit -framework WebKit -lpthread
// #include "ui.h"
import "C"

func Main(f func()) error {
	err := make(chan error)
	go main(f, err)
	return <-err
}

func main(f func(), errchan chan error) {
	runtime.LockOSThread()
	e := C.initUIThread()
	if e != nil {
		errchan <- fmt.Errorf("%s", C.GoString(e))
		return
	}
	QueueMain(f)
	C.runUIThread()
	errchan <- nil
}

func Quit() {
	C.stopUIThread()
}

// These prevent the passing of Go functions into C land.
// TODO make an actual sparse list instead of this monotonic map thingy
var (
	qmmap = make(map[C.uintptr_t]func())
	qmcurrent = C.uintptr_t(0)
	qmlock sync.Mutex
)

func QueueMain(f func()) {
	qmlock.Lock()
	defer qmlock.Unlock()

	n := C.uintptr_t(0)
	for {
		n = qmcurrent
		qmcurrent++
		if qmmap[n] == nil {
			break
		}
	}
	qmmap[n] = f
	C.queueUIThread(n)
}

//export doQueued
func doQueued(n C.uintptr_t) {
	qmlock.Lock()

	f := qmmap[n]
	delete(qmmap, n)

	// allow QueueMain() to be called by a queued function
	qmlock.Unlock()

	f()
}

// no need to lock this; this API is only safe on the main thread
var shouldQuitFunc func() bool

func OnShouldQuit(f func() bool) {
	shouldQuitFunc = f
}

//export shouldQuit
func shouldQuit() C.int {
	if shouldQuitFunc == nil {
		return 0
	}
	return frombool(shouldQuitFunc())
}
