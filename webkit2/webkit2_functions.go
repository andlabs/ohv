package webkit2

/*
#include <webkit2/webkit2.h>
#include <stdlib.h>
#cgo pkg-config: webkitgtk2-3.0
*/
import "C"

import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/gtk"
import (
	"errors"
	"reflect"
	"unsafe"
)

func init() {
	var _ unsafe.Pointer
	var _ reflect.SliceHeader
	_ = errors.New("")
}

var _ = gobject.UnusedFix_
var _ = gtk.UnusedFix_

func DownloadErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_download_error_quark()
	return
}

func FaviconDatabaseErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_favicon_database_error_quark()
	return
}

func JavascriptErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_javascript_error_quark()
	return
}

func NetworkErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_network_error_quark()
	return
}

func PluginErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_plugin_error_quark()
	return
}

func PolicyErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_policy_error_quark()
	return
}

func PrintErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_print_error_quark()
	return
}

func SnapshotErrorQuark() (return__ C.GQuark) {
	return__ = C.webkit_snapshot_error_quark()
	return
}
