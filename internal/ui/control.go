// 2 january 2016
package ui

type Control interface {
	Destroy()
	Handle() uintptr
}
