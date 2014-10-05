// 5 october 2014
package main

import (
	// ...
)

// #include "gtk_unix.h"
import "C"

//export navtreePathValid
func navtreePathValid(path *C.GtkTreePath) C.gboolean {
	// TODO
	return C.FALSE
}

//export navtreeItemName
func navtreeItemName(path *C.GtkTreePath) *C.char {
	// TODO
	return nil
}

//export navtreeBookCount
func navtreeBookCount() C.gint {
	// TODO
	return 0
}

//export navtreeChildCount
func navtreeChildCount(path *C.GtkTreePath) C.gint {
	// TODO
	return 0
}
