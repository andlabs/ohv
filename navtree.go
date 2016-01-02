// 5 october 2014
package main

// #include "cocoa_darwin.h"
import "C"

func indexArrayToSlice(ia C.indexArray) []int {
	n := make([]int, C.indexArrayLen(ia))
	for i := 0; i < len(n); i++ {
		n[i] = int(C.indexArrayIndex(ia, C.intmax_t(i)))
	}
	return n
}

func navtreeTopic(ia C.indexArray) Topic {
	n := indexArrayToSlice(ia)
	if n[0] >= len(Library) {
		return nil
	}
	t := Library[n[0]]
	for _, i := range n[1:] {
		if i >= len(t.Children()) {
			return nil
		}
		t = t.Children()[i]
	}
	return t
}

//export navtreePathValid
func navtreePathValid(ia C.indexArray) C.int {
	t := navtreeTopic(ia)
	if t == nil {
		return 0
	}
	return 1
}

//export navtreeItemName
func navtreeItemName(ia C.indexArray) *C.char {
	t := navtreeTopic(ia)
	if t == nil {
		return nil
	}
	return C.CString(t.Name())		// freed on the C side
}

//export navtreeBookCount
func navtreeBookCount() C.intmax_t {
	return C.intmax_t(len(Library))
}

//export navtreeChildCount
func navtreeChildCount(ia C.indexArray) C.intmax_t {
	t := navtreeTopic(ia)
	if t == nil {
		return 0
	}
	return C.intmax_t(len(t.Children()))
}

func navtreePathTo(t Topic) C.indexArray {
	ia := C.newIndexArray()
	for t != nil {
		p := t.Parent()
		children := Library
		if p != nil {
			children = p.Children()
		}
		i := C.intmax_t(0)
		for _, c := range children {
			if c == t {
				break
			}
			i++
		}
		if i >= C.intmax_t(len(children)) {
			panic("parent without known child in navtreePathTo()")
		}
		C.indexArrayPrepend(ia, i)
		t = p
	}
	return ia
}

//export navtreeSelected
func navtreeSelected(ov C.goid, ia C.indexArray) {
	navigate(ov, navtreeTopic(ia))
}
