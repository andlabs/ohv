// 5 october 2014
package main

import (
	"reflect"
	"unsafe"
)

// #include "gtk_unix.h"
import "C"

func pathToSlice(path *C.GtkTreePath) []int {
	var slice reflect.SliceHeader

	slice.Data = uintptr(unsafe.Pointer(C.gtk_tree_path_get_indices(path)))
	slice.Len = int(C.gtk_tree_path_get_depth(path))
	slice.Cap = slice.Len
	s := (*[]C.gint)(unsafe.Pointer(&slice))
	n := make([]int, len(*s))
	for i := 0; i < len(*s); i++ {
		n[i] = int((*s)[i])
	}
	return n
}

func navtreeTopic(path *C.GtkTreePath) Topic {
	n := pathToSlice(path)
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
func navtreePathValid(path *C.GtkTreePath) C.gboolean {
	t := navtreeTopic(path)
	if t == nil {
		return C.FALSE
	}
	return C.TRUE
}

//export navtreeItemName
func navtreeItemName(path *C.GtkTreePath) *C.char {
	t := navtreeTopic(path)
	if t == nil {
		return nil
	}
	return C.CString(t.Name())		// freed on the C side
}

//export navtreeBookCount
func navtreeBookCount() C.gint {
	return C.gint(len(Library))
}

//export navtreeChildCount
func navtreeChildCount(path *C.GtkTreePath) C.gint {
	t := navtreeTopic(path)
	if t == nil {
		return 0
	}
	return C.gint(len(t.Children()))
}

func navtreePathTo(t Topic) *C.GtkTreePath {
	path := C.gtk_tree_path_new()
	for t != nil {
		p := t.Parent()
		children := Library
		if p != nil {
			children = p.Children()
		}
		i := C.gint(0)
		for _, c := range children {
			if c == t {
				break
			}
			i++
		}
		if i >= C.gint(len(children)) {
			panic("parent without known child in navtreePathTo()")
		}
		C.gtk_tree_path_prepend_index(path, i)
		t = p
	}
	return path
}
