package webkit2

/*
#include <webkit2/webkit2.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "reflect"
import "errors"
import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/gtk"

func init() {
	_ = unsafe.Pointer(nil)
	_ = reflect.ValueOf(nil)
	_ = errors.New("")
}

var UnusedFix_ bool
var _ = gobject.UnusedFix_
var _ = gtk.UnusedFix_

/*
Creates a new #WebKitContextMenu object to be used as a submenu of an existing
#WebKitContextMenu. The context menu of a #WebKitWebView is created by the view
and passed as an argument of #WebKitWebView::context-menu signal.
To add items to the menu use webkit_context_menu_prepend(),
webkit_context_menu_append() or webkit_context_menu_insert().
See also webkit_context_menu_new_with_items() to create a #WebKitContextMenu with
a list of initial items.
*/
func ContextMenuNew() (return__ *ContextMenu) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_context_menu_new()
	if __cgo__return__ != nil {
		return__ = NewContextMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitContextMenu object to be used as a submenu of an existing
#WebKitContextMenu with the given initial items.
See also webkit_context_menu_new()
*/
func ContextMenuNewWithItems(items *C.GList) (return__ *ContextMenu) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_context_menu_new_with_items(items)
	if __cgo__return__ != nil {
		return__ = NewContextMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitContextMenuItem for the given @action.
*/
func ContextMenuItemNew(action *C.GtkAction) (return__ *ContextMenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_context_menu_item_new(action)
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitContextMenuItem for the given stock action.
Stock actions are handled automatically by WebKit so that, for example,
when a menu item created with a %WEBKIT_CONTEXT_MENU_ACTION_STOP is
activated the action associated will be handled by WebKit and the current
load operation will be stopped. You can get the #GtkAction of a
#WebKitContextMenuItem created with a #WebKitContextMenuAction with
webkit_context_menu_item_get_action() and connect to #GtkAction::activate signal
to be notified when the item is activated. But you can't prevent the asociated
action from being performed.
*/
func ContextMenuItemNewFromStockAction(action C.WebKitContextMenuAction) (return__ *ContextMenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_context_menu_item_new_from_stock_action(action)
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitContextMenuItem for the given stock action using the given @label.
Stock actions have a predefined label, this method can be used to create a
#WebKitContextMenuItem for a #WebKitContextMenuAction but using a custom label.
*/
func ContextMenuItemNewFromStockActionWithLabel(action C.WebKitContextMenuAction, label string) (return__ *ContextMenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_context_menu_item_new_from_stock_action_with_label(action, __cgo__label)
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitContextMenuItem representing a separator.
*/
func ContextMenuItemNewSeparator() (return__ *ContextMenuItem) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_context_menu_item_new_separator()
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitContextMenuItem using the given @label with a submenu.
*/
func ContextMenuItemNewWithSubmenu(label string, submenu IsContextMenu) (return__ *ContextMenuItem) {
	__cgo__label := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_context_menu_item_new_with_submenu(__cgo__label, submenu.GetContextMenuPointer())
	C.free(unsafe.Pointer(__cgo__label))
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Create a new #WebKitPrintOperation to print @web_view contents.
*/
func PrintOperationNew(web_view IsWebView) (return__ *PrintOperation) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_print_operation_new(web_view.GetWebViewPointer())
	if __cgo__return__ != nil {
		return__ = NewPrintOperationFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitSettings instance with default values. It must
be manually attached to a #WebKitWebViewGroup.
See also webkit_settings_new_with_settings().
*/
func SettingsNew() (return__ *Settings) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_settings_new()
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

// webkit_settings_new_with_settings is not generated due to varargs

/*
Creates a new #WebKitURIRequest for the given URI.
*/
func URIRequestNew(uri string) (return__ *URIRequest) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_uri_request_new(__cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	if __cgo__return__ != nil {
		return__ = NewURIRequestFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitWebView with the default #WebKitWebContext and the
default #WebKitWebViewGroup.
See also webkit_web_view_new_with_context() and webkit_web_view_new_with_group().
*/
func WebViewNew() (return__ *WebView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_web_view_new()
	if __cgo__return__ != nil {
		return__ = NewWebViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitWebView with the given #WebKitWebContext and the
default #WebKitWebViewGroup.
See also webkit_web_view_new_with_group().
*/
func WebViewNewWithContext(context IsWebContext) (return__ *WebView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_web_view_new_with_context(context.GetWebContextPointer())
	if __cgo__return__ != nil {
		return__ = NewWebViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitWebView with the given #WebKitWebViewGroup.
The view will be part of @group and it will be affected by the
group properties like the settings.
*/
func WebViewNewWithGroup(group IsWebViewGroup) (return__ *WebView) {
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_web_view_new_with_group(group.GetWebViewGroupPointer())
	if __cgo__return__ != nil {
		return__ = NewWebViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Creates a new #WebKitWebViewGroup with the given @name.
If @name is %NULL a unique identifier name will be created
automatically.
The newly created #WebKitWebViewGroup doesn't contain any
#WebKitWebView, web views are added to the new group when created
with webkit_web_view_new_with_group() passing the group.
*/
func WebViewGroupNew(name string) (return__ *WebViewGroup) {
	__cgo__name := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	var __cgo__return__ interface{}
	__cgo__return__ = C.webkit_web_view_group_new(__cgo__name)
	C.free(unsafe.Pointer(__cgo__name))
	if __cgo__return__ != nil {
		return__ = NewWebViewGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}
