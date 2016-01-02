// 2 january 2016
package ui

import (
	"fmt"
	"runtime"
	"sync"
)

// #cgo CFLAGS: -mmacosx-version-min=10.7 -DMACOSX_DEPLOYMENT_TARGET=10.7
// #cgo LDFLAGS: -mmacosx-version-min=10.7 -lobjc -framework Foundation -framework AppKit -framework WebKit -lpthread
// #include "cocoa_darwin.h"
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

// QueueMain queues f to be executed on the GUI thread when
// next possible. It returns immediately; that is, it does not wait
// for the function to actually be executed. QueueMain is the only
// function that can be called from other goroutines, and its
// primary purpose is to allow communication between other
// goroutines and the GUI thread. Calling QueueMain after Quit
// has been called results in undefined behavior.
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
	if shouldQuitFunc() {
		return 1
	}
	return 0
}
