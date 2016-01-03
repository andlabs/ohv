// 3 january 2016
package ui

// #include "ui.h"
import "C"

type Margin struct {
	id		C.id
	child		Control
}

func NewMargin(child Control) *Margin {
	m := new(Margin)
	m.child = child
	m.id = C.newMargin(fromuintptr(m.child.Handle()))
	return m
}

func (m *Margin) Destroy() {
	m.child.Destroy()
	C.marginDestroy(m.id)
}

func (m *Margin) Handle() uintptr {
	return touintptr(m.id)
}
