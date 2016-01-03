// 3 january 2016
package ui

// #include "ui.h"
import "C"

type Splitter struct {
	id		C.id
	left		Control
	right		Control
}

func NewSplitter(left Control, right Control) *Splitter {
	s := new(Splitter)
	s.left = left
	s.right = right
	s.id = C.newSplitter(fromuintptr(s.left.Handle()),
		fromuintptr(s.right.Handle()))
	return s
}

func (s *Splitter) Destroy() {
	s.left.Destroy()
	s.right.Destroy()
	C.splitterDestroy(s.id)
}

func (s *Splitter) Handle() uintptr {
	return touintptr(s.id)
}

func (s *Splitter) SetPosition(pos int) {
	C.splitterSetPosition(s.id, C.intmax_t(pos))
}
