// 12 december 2015

package ui

// #include "ui.h"
// static inline uintptr_t fromid(id x)
// {
// 	return (uintptr_t) x;
// }
// static inline id toid(uintptr_t x)
// {
// 	return (id) x;
// }
import "C"

func tobool(b C.int) bool {
	return b != 0
}

func frombool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func toid(id uintptr) C.id {
	return C.toid(C.uintptr_t(id))
}

func fromid(id C.id) uintptr {
	return uintptr(C.fromid(id))
}
