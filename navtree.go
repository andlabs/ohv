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

//export navtreePathValid
func navtreePathValid(path *C.GtkTreePath) C.gboolean {
	// TODO merge with all of the below
	n := pathToSlice(path)
	if n[0] >= len(Library) {
		return C.FALSE
	}
	t := Library[n[0]]
	for _, i := range n[1:] {
		if i >= len(t.Children()) {
			return C.FALSE
		}
		t = t.Children()[i]
	}
	return C.TRUE
}

//export navtreeItemName
func navtreeItemName(path *C.GtkTreePath) *C.char {
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
