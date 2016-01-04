// 12 december 2015

package ui

import (
	"net/url"
)

// #include "ui.h"
// static inline uintptr_t fromid(id x)
// {
// 	return (uintptr_t) x;
// }
// static inline id toid(uintptr_t x)
// {
// 	return (id) x;
// }
// static inline void freestr(char *s)
// {
// 	free(s);
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

func touintptr(id C.id) uintptr {
	return uintptr(C.fromid(id))
}

func fromuintptr(id uintptr) C.id {
	return C.toid(C.uintptr_t(id))
}

func toURL(nsurl C.id) *url.URL {
	cs := C.fromNSURL(nsurl)
	u, err := url.Parse(C.GoString(cs))
	if err != nil {
		// TODO shouldn't happen
		panic(err)
	}
	C.freestr(cs)
	return u
}

func fromURL(u *url.URL) C.id {
	s := u.String()
	cs := C.CString(s)			// freed on the C side
	return C.toNSURL(cs)
}

func fromFileURL(path string, anchor string) C.id {
	var canchor *C.char

	cpath := C.CString(path)		// both are freed on the C side
	if anchor != "" {
		canchor = C.CString(anchor)
	}
	return C.toFileNSURL(cpath, canchor)
}
