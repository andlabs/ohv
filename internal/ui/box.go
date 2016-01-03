// 3 january 2016
package ui

// #include "ui.h"
import "C"

type Box struct {
	id		C.id
	top		Control
	bottom	Control
}

func NewBox(top Control, bottom Control) *Box {
	b := new(Box)
	b.top = top
	b.bottom = bottom
	b.id = C.newBox(fromuintptr(b.top.Handle()),
		fromuintptr(b.bottom.Handle()))
	return b
}

func (b *Box) Destroy() {
	b.top.Destroy()
	b.bottom.Destroy()
	C.boxDestroy(b.id)
}

func (b *Box) Handle() uintptr {
	return touintptr(b.id)
}
