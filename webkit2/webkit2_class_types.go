package webkit2

/*
#include <webkit2/webkit2.h>
#include <stdlib.h>
#include <glib-object.h>
#cgo pkg-config: gobject-2.0
*/
import "C"
import "unsafe"
import "github.com/reusee/ggir/gobject"
import "github.com/reusee/ggir/gtk"
import "runtime"

func init() {
	_ = unsafe.Pointer(nil)
	_ = runtime.Compiler
}

var _ = gobject.UnusedFix_
var _ = gtk.UnusedFix_

type AuthenticationRequest struct {
	*TraitAuthenticationRequest
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewAuthenticationRequestFromCPointer(p unsafe.Pointer) *AuthenticationRequest {
	ret := &AuthenticationRequest{
		NewTraitAuthenticationRequest(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *AuthenticationRequest) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type BackForwardList struct {
	*TraitBackForwardList
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBackForwardListFromCPointer(p unsafe.Pointer) *BackForwardList {
	ret := &BackForwardList{
		NewTraitBackForwardList(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *BackForwardList) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type BackForwardListItem struct {
	*TraitBackForwardListItem
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewBackForwardListItemFromCPointer(p unsafe.Pointer) *BackForwardListItem {
	ret := &BackForwardListItem{
		NewTraitBackForwardListItem(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *BackForwardListItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ContextMenu struct {
	*TraitContextMenu
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewContextMenuFromCPointer(p unsafe.Pointer) *ContextMenu {
	ret := &ContextMenu{
		NewTraitContextMenu(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ContextMenu) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ContextMenuItem struct {
	*TraitContextMenuItem
	*gobject.TraitInitiallyUnowned
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewContextMenuItemFromCPointer(p unsafe.Pointer) *ContextMenuItem {
	ret := &ContextMenuItem{
		NewTraitContextMenuItem(p),
		gobject.NewTraitInitiallyUnowned(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ContextMenuItem) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type CookieManager struct {
	*TraitCookieManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewCookieManagerFromCPointer(p unsafe.Pointer) *CookieManager {
	ret := &CookieManager{
		NewTraitCookieManager(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *CookieManager) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Download struct {
	*TraitDownload
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewDownloadFromCPointer(p unsafe.Pointer) *Download {
	ret := &Download{
		NewTraitDownload(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Download) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FaviconDatabase struct {
	*TraitFaviconDatabase
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFaviconDatabaseFromCPointer(p unsafe.Pointer) *FaviconDatabase {
	ret := &FaviconDatabase{
		NewTraitFaviconDatabase(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FaviconDatabase) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FileChooserRequest struct {
	*TraitFileChooserRequest
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFileChooserRequestFromCPointer(p unsafe.Pointer) *FileChooserRequest {
	ret := &FileChooserRequest{
		NewTraitFileChooserRequest(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FileChooserRequest) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FindController struct {
	*TraitFindController
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFindControllerFromCPointer(p unsafe.Pointer) *FindController {
	ret := &FindController{
		NewTraitFindController(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FindController) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type FormSubmissionRequest struct {
	*TraitFormSubmissionRequest
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewFormSubmissionRequestFromCPointer(p unsafe.Pointer) *FormSubmissionRequest {
	ret := &FormSubmissionRequest{
		NewTraitFormSubmissionRequest(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *FormSubmissionRequest) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type GeolocationPermissionRequest struct {
	*TraitGeolocationPermissionRequest
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewGeolocationPermissionRequestFromCPointer(p unsafe.Pointer) *GeolocationPermissionRequest {
	ret := &GeolocationPermissionRequest{
		NewTraitGeolocationPermissionRequest(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *GeolocationPermissionRequest) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type HitTestResult struct {
	*TraitHitTestResult
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewHitTestResultFromCPointer(p unsafe.Pointer) *HitTestResult {
	ret := &HitTestResult{
		NewTraitHitTestResult(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *HitTestResult) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type NavigationPolicyDecision struct {
	*TraitNavigationPolicyDecision
	*TraitPolicyDecision
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewNavigationPolicyDecisionFromCPointer(p unsafe.Pointer) *NavigationPolicyDecision {
	ret := &NavigationPolicyDecision{
		NewTraitNavigationPolicyDecision(p),
		NewTraitPolicyDecision(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *NavigationPolicyDecision) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Plugin struct {
	*TraitPlugin
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPluginFromCPointer(p unsafe.Pointer) *Plugin {
	ret := &Plugin{
		NewTraitPlugin(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Plugin) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PolicyDecision struct {
	*TraitPolicyDecision
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPolicyDecisionFromCPointer(p unsafe.Pointer) *PolicyDecision {
	ret := &PolicyDecision{
		NewTraitPolicyDecision(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PolicyDecision) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type PrintOperation struct {
	*TraitPrintOperation
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewPrintOperationFromCPointer(p unsafe.Pointer) *PrintOperation {
	ret := &PrintOperation{
		NewTraitPrintOperation(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *PrintOperation) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type ResponsePolicyDecision struct {
	*TraitResponsePolicyDecision
	*TraitPolicyDecision
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewResponsePolicyDecisionFromCPointer(p unsafe.Pointer) *ResponsePolicyDecision {
	ret := &ResponsePolicyDecision{
		NewTraitResponsePolicyDecision(p),
		NewTraitPolicyDecision(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *ResponsePolicyDecision) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type SecurityManager struct {
	*TraitSecurityManager
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSecurityManagerFromCPointer(p unsafe.Pointer) *SecurityManager {
	ret := &SecurityManager{
		NewTraitSecurityManager(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *SecurityManager) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type Settings struct {
	*TraitSettings
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewSettingsFromCPointer(p unsafe.Pointer) *Settings {
	ret := &Settings{
		NewTraitSettings(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *Settings) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type URIRequest struct {
	*TraitURIRequest
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewURIRequestFromCPointer(p unsafe.Pointer) *URIRequest {
	ret := &URIRequest{
		NewTraitURIRequest(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *URIRequest) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type URIResponse struct {
	*TraitURIResponse
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewURIResponseFromCPointer(p unsafe.Pointer) *URIResponse {
	ret := &URIResponse{
		NewTraitURIResponse(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *URIResponse) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type URISchemeRequest struct {
	*TraitURISchemeRequest
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewURISchemeRequestFromCPointer(p unsafe.Pointer) *URISchemeRequest {
	ret := &URISchemeRequest{
		NewTraitURISchemeRequest(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *URISchemeRequest) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WebContext struct {
	*TraitWebContext
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWebContextFromCPointer(p unsafe.Pointer) *WebContext {
	ret := &WebContext{
		NewTraitWebContext(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WebContext) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WebInspector struct {
	*TraitWebInspector
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWebInspectorFromCPointer(p unsafe.Pointer) *WebInspector {
	ret := &WebInspector{
		NewTraitWebInspector(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WebInspector) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WebResource struct {
	*TraitWebResource
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWebResourceFromCPointer(p unsafe.Pointer) *WebResource {
	ret := &WebResource{
		NewTraitWebResource(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WebResource) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WebView struct {
	*TraitWebView
	*TraitWebViewBase
	*gtk.TraitContainer
	*gtk.TraitWidget
	CPointer unsafe.Pointer
}

func NewWebViewFromCPointer(p unsafe.Pointer) *WebView {
	ret := &WebView{
		NewTraitWebView(p),
		NewTraitWebViewBase(p),
		gtk.NewTraitContainer(p),
		gtk.NewTraitWidget(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WebView) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WebViewBase struct {
	*TraitWebViewBase
	*gtk.TraitContainer
	*gtk.TraitWidget
	CPointer unsafe.Pointer
}

func NewWebViewBaseFromCPointer(p unsafe.Pointer) *WebViewBase {
	ret := &WebViewBase{
		NewTraitWebViewBase(p),
		gtk.NewTraitContainer(p),
		gtk.NewTraitWidget(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WebViewBase) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WebViewGroup struct {
	*TraitWebViewGroup
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWebViewGroupFromCPointer(p unsafe.Pointer) *WebViewGroup {
	ret := &WebViewGroup{
		NewTraitWebViewGroup(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WebViewGroup) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}

type WindowProperties struct {
	*TraitWindowProperties
	*gobject.TraitObject
	CPointer unsafe.Pointer
}

func NewWindowPropertiesFromCPointer(p unsafe.Pointer) *WindowProperties {
	ret := &WindowProperties{
		NewTraitWindowProperties(p),
		gobject.NewTraitObject(p),
		p,
	}
	C.g_object_ref_sink(C.gpointer(p))
	runtime.SetFinalizer(ret, func(p *WindowProperties) {
		C.g_object_unref(C.gpointer(unsafe.Pointer(p.CPointer)))
	})
	return ret
}
