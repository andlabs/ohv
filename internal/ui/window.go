// 2 january 2016
package ui

// #include "ui.h"
import "C"

var windows = make(map[C.id]*Window)

type Window struct {
	id			C.id
	child			Control
	onClosing		func() bool
}

func NewWindow(title string, width int, height int) *Window {
	w := new(Window)
	ctitle := C.CString(title)		// freed on C side
	w.id = C.newWindow(ctitle, C.int(width), C.int(height))
	windows[w.id] = w
	return w
}

func (w *Window) Destroy() {
	if w.child != nil {
		c := w.child
		w.SetChild(nil)
		c.Destroy()
	}
	delete(windows, w.id)
	C.windowDestroy(w.id)
}

func (w *Window) Handle() uintptr {
	return touintptr(w.id)
}

func (w *Window) Move(x int, y int) {
	C.windowMove(w.id, C.int(x), C.int(y))
}

func (w *Window) Center() {
	C.windowCenter(w.id)
}

func (w *Window) Show() {
	C.windowShow(w.id)
}

func (w *Window) SetChild(c Control) {
	if w.child != nil {
		C.windowUnsetChild(w.id, fromuintptr(w.child.Handle()))
	}
	w.child = c
	if w.child != nil {
		C.windowSetChild(w.id, fromuintptr(w.child.Handle()))
	}
}

func (w *Window) OnClosing(f func() bool) {
	w.onClosing = f
}

//export doWindowClosing
func doWindowClosing(id C.id) C.int {
	w := windows[id]
	if w.onClosing == nil {
		return 0
	}
	return frombool(w.onClosing())
}

func (w *Window) MsgBoxSysError(sysError uintptr) {
	C.windowMsgBoxSysError(w.id, fromuintptr(sysError))
}
