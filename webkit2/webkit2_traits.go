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

var _ = gobject.UnusedFix_
var _ = gtk.UnusedFix_

type TraitAuthenticationRequest struct {
	CPointer *C.WebKitAuthenticationRequest
}
type IsAuthenticationRequest interface {
	GetAuthenticationRequestPointer() *C.WebKitAuthenticationRequest
}

func (self *TraitAuthenticationRequest) GetAuthenticationRequestPointer() *C.WebKitAuthenticationRequest {
	return self.CPointer
}
func NewTraitAuthenticationRequest(p unsafe.Pointer) *TraitAuthenticationRequest {
	return &TraitAuthenticationRequest{(*C.WebKitAuthenticationRequest)(p)}
}

/*
Authenticate the #WebKitAuthenticationRequest using the #WebKitCredential
supplied. To continue without credentials, pass %NULL as @credential.
*/
func (self *TraitAuthenticationRequest) Authenticate(credential *C.WebKitCredential) {
	C.webkit_authentication_request_authenticate(self.CPointer, credential)
	return
}

/*
Determine whether the authentication method associated with this
#WebKitAuthenticationRequest should allow the storage of credentials.
This will return %FALSE if webkit doesn't support credential storing
or if private browsing is enabled.
*/
func (self *TraitAuthenticationRequest) CanSaveCredentials() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_authentication_request_can_save_credentials(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Cancel the authentication challenge. This will also cancel the page loading and result in a
#WebKitWebView::load-failed signal with a #WebKitNetworkError of type %WEBKIT_NETWORK_ERROR_CANCELLED being emitted.
*/
func (self *TraitAuthenticationRequest) Cancel() {
	C.webkit_authentication_request_cancel(self.CPointer)
	return
}

/*
Get the host that this authentication challenge is applicable to.
*/
func (self *TraitAuthenticationRequest) GetHost() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_authentication_request_get_host(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the port that this authentication challenge is applicable to.
*/
func (self *TraitAuthenticationRequest) GetPort() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_authentication_request_get_port(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Get the #WebKitCredential of the proposed authentication challenge that was
stored from a previous session. The client can use this directly for
authentication or construct their own #WebKitCredential.
*/
func (self *TraitAuthenticationRequest) GetProposedCredential() (return__ *C.WebKitCredential) {
	return__ = C.webkit_authentication_request_get_proposed_credential(self.CPointer)
	return
}

/*
Get the realm that this authentication challenge is applicable to.
*/
func (self *TraitAuthenticationRequest) GetRealm() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_authentication_request_get_realm(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the authentication scheme of the authentication challenge.
*/
func (self *TraitAuthenticationRequest) GetScheme() (return__ C.WebKitAuthenticationScheme) {
	return__ = C.webkit_authentication_request_get_scheme(self.CPointer)
	return
}

/*
Determine whether the authentication challenge is associated with a proxy server rather than an "origin" server.
*/
func (self *TraitAuthenticationRequest) IsForProxy() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_authentication_request_is_for_proxy(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determine whether this this is a first attempt or a retry for this authentication challenge.
*/
func (self *TraitAuthenticationRequest) IsRetry() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_authentication_request_is_retry(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitBackForwardList struct{ CPointer *C.WebKitBackForwardList }
type IsBackForwardList interface {
	GetBackForwardListPointer() *C.WebKitBackForwardList
}

func (self *TraitBackForwardList) GetBackForwardListPointer() *C.WebKitBackForwardList {
	return self.CPointer
}
func NewTraitBackForwardList(p unsafe.Pointer) *TraitBackForwardList {
	return &TraitBackForwardList{(*C.WebKitBackForwardList)(p)}
}

/*
Returns the item that precedes the current item.
*/
func (self *TraitBackForwardList) GetBackItem() (return__ *BackForwardListItem) {
	var __cgo__return__ *C.WebKitBackForwardListItem
	__cgo__return__ = C.webkit_back_forward_list_get_back_item(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewBackForwardListItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

func (self *TraitBackForwardList) GetBackList() (return__ *C.GList) {
	return__ = C.webkit_back_forward_list_get_back_list(self.CPointer)
	return
}

func (self *TraitBackForwardList) GetBackListWithLimit(limit uint) (return__ *C.GList) {
	return__ = C.webkit_back_forward_list_get_back_list_with_limit(self.CPointer, C.guint(limit))
	return
}

/*
Returns the current item in @back_forward_list.
*/
func (self *TraitBackForwardList) GetCurrentItem() (return__ *BackForwardListItem) {
	var __cgo__return__ *C.WebKitBackForwardListItem
	__cgo__return__ = C.webkit_back_forward_list_get_current_item(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewBackForwardListItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the item that follows the current item.
*/
func (self *TraitBackForwardList) GetForwardItem() (return__ *BackForwardListItem) {
	var __cgo__return__ *C.WebKitBackForwardListItem
	__cgo__return__ = C.webkit_back_forward_list_get_forward_item(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewBackForwardListItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

func (self *TraitBackForwardList) GetForwardList() (return__ *C.GList) {
	return__ = C.webkit_back_forward_list_get_forward_list(self.CPointer)
	return
}

func (self *TraitBackForwardList) GetForwardListWithLimit(limit uint) (return__ *C.GList) {
	return__ = C.webkit_back_forward_list_get_forward_list_with_limit(self.CPointer, C.guint(limit))
	return
}

func (self *TraitBackForwardList) GetLength() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_back_forward_list_get_length(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Returns the item at a given index relative to the current item.
*/
func (self *TraitBackForwardList) GetNthItem(index int) (return__ *BackForwardListItem) {
	var __cgo__return__ *C.WebKitBackForwardListItem
	__cgo__return__ = C.webkit_back_forward_list_get_nth_item(self.CPointer, C.gint(index))
	if __cgo__return__ != nil {
		return__ = NewBackForwardListItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

type TraitBackForwardListItem struct{ CPointer *C.WebKitBackForwardListItem }
type IsBackForwardListItem interface {
	GetBackForwardListItemPointer() *C.WebKitBackForwardListItem
}

func (self *TraitBackForwardListItem) GetBackForwardListItemPointer() *C.WebKitBackForwardListItem {
	return self.CPointer
}
func NewTraitBackForwardListItem(p unsafe.Pointer) *TraitBackForwardListItem {
	return &TraitBackForwardListItem{(*C.WebKitBackForwardListItem)(p)}
}

/*
See also webkit_back_forward_list_item_get_uri().
*/
func (self *TraitBackForwardListItem) GetOriginalUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_back_forward_list_item_get_original_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func (self *TraitBackForwardListItem) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_back_forward_list_item_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
This URI may differ from the original URI if the page was,
for example, redirected to a new location.
See also webkit_back_forward_list_item_get_original_uri().
*/
func (self *TraitBackForwardListItem) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_back_forward_list_item_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitContextMenu struct{ CPointer *C.WebKitContextMenu }
type IsContextMenu interface {
	GetContextMenuPointer() *C.WebKitContextMenu
}

func (self *TraitContextMenu) GetContextMenuPointer() *C.WebKitContextMenu {
	return self.CPointer
}
func NewTraitContextMenu(p unsafe.Pointer) *TraitContextMenu {
	return &TraitContextMenu{(*C.WebKitContextMenu)(p)}
}

/*
Adds @item at the end of the @menu.
*/
func (self *TraitContextMenu) Append(item IsContextMenuItem) {
	C.webkit_context_menu_append(self.CPointer, item.GetContextMenuItemPointer())
	return
}

/*
Gets the first item in the @menu.
*/
func (self *TraitContextMenu) First() (return__ *ContextMenuItem) {
	var __cgo__return__ *C.WebKitContextMenuItem
	__cgo__return__ = C.webkit_context_menu_first(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the item at the given position in the @menu.
*/
func (self *TraitContextMenu) GetItemAtPosition(position uint) (return__ *ContextMenuItem) {
	var __cgo__return__ *C.WebKitContextMenuItem
	__cgo__return__ = C.webkit_context_menu_get_item_at_position(self.CPointer, C.guint(position))
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the item list of @menu.
*/
func (self *TraitContextMenu) GetItems() (return__ *C.GList) {
	return__ = C.webkit_context_menu_get_items(self.CPointer)
	return
}

/*
Gets the length of the @menu.
*/
func (self *TraitContextMenu) GetNItems() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_context_menu_get_n_items(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Inserts @item into the @menu at the given position.
If @position is negative, or is larger than the number of items
in the #WebKitContextMenu, the item is added on to the end of
the @menu. The first position is 0.
*/
func (self *TraitContextMenu) Insert(item IsContextMenuItem, position int) {
	C.webkit_context_menu_insert(self.CPointer, item.GetContextMenuItemPointer(), C.gint(position))
	return
}

/*
Gets the last item in the @menu.
*/
func (self *TraitContextMenu) Last() (return__ *ContextMenuItem) {
	var __cgo__return__ *C.WebKitContextMenuItem
	__cgo__return__ = C.webkit_context_menu_last(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewContextMenuItemFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Moves @item to the given position in the @menu.
If @position is negative, or is larger than the number of items
in the #WebKitContextMenu, the item is added on to the end of
the @menu.
The first position is 0.
*/
func (self *TraitContextMenu) MoveItem(item IsContextMenuItem, position int) {
	C.webkit_context_menu_move_item(self.CPointer, item.GetContextMenuItemPointer(), C.gint(position))
	return
}

/*
Adds @item at the beginning of the @menu.
*/
func (self *TraitContextMenu) Prepend(item IsContextMenuItem) {
	C.webkit_context_menu_prepend(self.CPointer, item.GetContextMenuItemPointer())
	return
}

/*
Removes @item from the @menu.
See also webkit_context_menu_remove_all() to remove all items.
*/
func (self *TraitContextMenu) Remove(item IsContextMenuItem) {
	C.webkit_context_menu_remove(self.CPointer, item.GetContextMenuItemPointer())
	return
}

/*
Removes all items of the @menu.
*/
func (self *TraitContextMenu) RemoveAll() {
	C.webkit_context_menu_remove_all(self.CPointer)
	return
}

type TraitContextMenuItem struct{ CPointer *C.WebKitContextMenuItem }
type IsContextMenuItem interface {
	GetContextMenuItemPointer() *C.WebKitContextMenuItem
}

func (self *TraitContextMenuItem) GetContextMenuItemPointer() *C.WebKitContextMenuItem {
	return self.CPointer
}
func NewTraitContextMenuItem(p unsafe.Pointer) *TraitContextMenuItem {
	return &TraitContextMenuItem{(*C.WebKitContextMenuItem)(p)}
}

/*
Gets the action associated to @item.
*/
func (self *TraitContextMenuItem) GetAction() (return__ *C.GtkAction) {
	return__ = C.webkit_context_menu_item_get_action(self.CPointer)
	return
}

/*
Gets the #WebKitContextMenuAction of @item. If the #WebKitContextMenuItem was not
created for a stock action %WEBKIT_CONTEXT_MENU_ACTION_CUSTOM will be
returned. If the #WebKitContextMenuItem is a separator %WEBKIT_CONTEXT_MENU_ACTION_NO_ACTION
will be returned.
*/
func (self *TraitContextMenuItem) GetStockAction() (return__ C.WebKitContextMenuAction) {
	return__ = C.webkit_context_menu_item_get_stock_action(self.CPointer)
	return
}

/*
Gets the submenu of @item.
*/
func (self *TraitContextMenuItem) GetSubmenu() (return__ *ContextMenu) {
	var __cgo__return__ *C.WebKitContextMenu
	__cgo__return__ = C.webkit_context_menu_item_get_submenu(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewContextMenuFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Checks whether @item is a separator.
*/
func (self *TraitContextMenuItem) IsSeparator() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_context_menu_item_is_separator(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Sets or replaces the @item submenu. If @submenu is %NULL the current
submenu of @item is removed.
*/
func (self *TraitContextMenuItem) SetSubmenu(submenu IsContextMenu) {
	C.webkit_context_menu_item_set_submenu(self.CPointer, submenu.GetContextMenuPointer())
	return
}

type TraitCookieManager struct{ CPointer *C.WebKitCookieManager }
type IsCookieManager interface {
	GetCookieManagerPointer() *C.WebKitCookieManager
}

func (self *TraitCookieManager) GetCookieManagerPointer() *C.WebKitCookieManager {
	return self.CPointer
}
func NewTraitCookieManager(p unsafe.Pointer) *TraitCookieManager {
	return &TraitCookieManager{(*C.WebKitCookieManager)(p)}
}

/*
Delete all cookies of @cookie_manager
*/
func (self *TraitCookieManager) DeleteAllCookies() {
	C.webkit_cookie_manager_delete_all_cookies(self.CPointer)
	return
}

/*
Remove all cookies of @cookie_manager for the given @domain.
*/
func (self *TraitCookieManager) DeleteCookiesForDomain(domain string) {
	__cgo__domain := (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	C.webkit_cookie_manager_delete_cookies_for_domain(self.CPointer, __cgo__domain)
	C.free(unsafe.Pointer(__cgo__domain))
	return
}

/*
Asynchronously get the cookie acceptance policy of @cookie_manager.

When the operation is finished, @callback will be called. You can then call
webkit_cookie_manager_get_accept_policy_finish() to get the result of the operation.
*/
func (self *TraitCookieManager) GetAcceptPolicy(cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.webkit_cookie_manager_get_accept_policy(self.CPointer, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous operation started with webkit_cookie_manager_get_accept_policy().
*/
func (self *TraitCookieManager) GetAcceptPolicyFinish(result *C.GAsyncResult) (return__ C.WebKitCookieAcceptPolicy, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_cookie_manager_get_accept_policy_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously get the list of domains for which @cookie_manager contains cookies.

When the operation is finished, @callback will be called. You can then call
webkit_cookie_manager_get_domains_with_cookies_finish() to get the result of the operation.
*/
func (self *TraitCookieManager) GetDomainsWithCookies(cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.webkit_cookie_manager_get_domains_with_cookies(self.CPointer, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous operation started with webkit_cookie_manager_get_domains_with_cookies().
The return value is a %NULL terminated list of strings which should
be released with g_strfreev().
*/
func (self *TraitCookieManager) GetDomainsWithCookiesFinish(result *C.GAsyncResult) (return__ **C.gchar, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_cookie_manager_get_domains_with_cookies_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Set the cookie acceptance policy of @cookie_manager as @policy.
*/
func (self *TraitCookieManager) SetAcceptPolicy(policy C.WebKitCookieAcceptPolicy) {
	C.webkit_cookie_manager_set_accept_policy(self.CPointer, policy)
	return
}

/*
Set the @filename where non-session cookies are stored persistently using
@storage as the format to read/write the cookies.
Cookies are initially read from @filename to create an initial set of cookies.
Then, non-session cookies will be written to @filename when the WebKitCookieManager::changed
signal is emitted.
By default, @cookie_manager doesn't store the cookies persistenly, so you need to call this
method to keep cookies saved across sessions.
*/
func (self *TraitCookieManager) SetPersistentStorage(filename string, storage C.WebKitCookiePersistentStorage) {
	__cgo__filename := (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	C.webkit_cookie_manager_set_persistent_storage(self.CPointer, __cgo__filename, storage)
	C.free(unsafe.Pointer(__cgo__filename))
	return
}

type TraitDownload struct{ CPointer *C.WebKitDownload }
type IsDownload interface {
	GetDownloadPointer() *C.WebKitDownload
}

func (self *TraitDownload) GetDownloadPointer() *C.WebKitDownload {
	return self.CPointer
}
func NewTraitDownload(p unsafe.Pointer) *TraitDownload {
	return &TraitDownload{(*C.WebKitDownload)(p)}
}

/*
Cancels the download. When the ongoing download
operation is effectively cancelled the signal
#WebKitDownload::failed is emitted with
%WEBKIT_DOWNLOAD_ERROR_CANCELLED_BY_USER error.
*/
func (self *TraitDownload) Cancel() {
	C.webkit_download_cancel(self.CPointer)
	return
}

/*
Obtains the URI to which the downloaded file will be written. You
can connect to #WebKitDownload::created-destination to make
sure this method returns a valid destination.
*/
func (self *TraitDownload) GetDestination() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_download_get_destination(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the elapsed time in seconds, including any fractional part.
If the download finished, had an error or was cancelled this is
the time between its start and the event.
*/
func (self *TraitDownload) GetElapsedTime() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.webkit_download_get_elapsed_time(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the value of the #WebKitDownload:estimated-progress property.
You can monitor the estimated progress of the download operation by
connecting to the notify::estimated-progress signal of @download.
*/
func (self *TraitDownload) GetEstimatedProgress() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.webkit_download_get_estimated_progress(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Gets the length of the data already downloaded for @download
in bytes.
*/
func (self *TraitDownload) GetReceivedDataLength() (return__ uint64) {
	var __cgo__return__ C.guint64
	__cgo__return__ = C.webkit_download_get_received_data_length(self.CPointer)
	return__ = uint64(__cgo__return__)
	return
}

/*
Retrieves the #WebKitURIRequest object that backs the download
process.
*/
func (self *TraitDownload) GetRequest() (return__ *URIRequest) {
	var __cgo__return__ *C.WebKitURIRequest
	__cgo__return__ = C.webkit_download_get_request(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewURIRequestFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Retrieves the #WebKitURIResponse object that backs the download
process. This method returns %NULL if called before the response
is received from the server. You can connect to notify::response
signal to be notified when the response is received.
*/
func (self *TraitDownload) GetResponse() (return__ *URIResponse) {
	var __cgo__return__ *C.WebKitURIResponse
	__cgo__return__ = C.webkit_download_get_response(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewURIResponseFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the #WebKitWebView that initiated the download.
*/
func (self *TraitDownload) GetWebView() (return__ *WebView) {
	var __cgo__return__ *C.WebKitWebView
	__cgo__return__ = C.webkit_download_get_web_view(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Sets the URI to which the downloaded file will be written.
This method should be called before the download transfer
starts or it will not have any effect on the ongoing download
operation. To set the destination using the filename suggested
by the server connect to #WebKitDownload::decide-destination
signal and call webkit_download_set_destination(). If you want to
set a fixed destination URI that doesn't depend on the suggested
filename you can connect to notify::response signal and call
webkit_download_set_destination().
If #WebKitDownload::decide-destination signal is not handled
and destination URI is not set when the download tranfer starts,
the file will be saved with the filename suggested by the server in
%G_USER_DIRECTORY_DOWNLOAD directory.
*/
func (self *TraitDownload) SetDestination(uri string) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	C.webkit_download_set_destination(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return
}

type TraitFaviconDatabase struct{ CPointer *C.WebKitFaviconDatabase }
type IsFaviconDatabase interface {
	GetFaviconDatabasePointer() *C.WebKitFaviconDatabase
}

func (self *TraitFaviconDatabase) GetFaviconDatabasePointer() *C.WebKitFaviconDatabase {
	return self.CPointer
}
func NewTraitFaviconDatabase(p unsafe.Pointer) *TraitFaviconDatabase {
	return &TraitFaviconDatabase{(*C.WebKitFaviconDatabase)(p)}
}

/*
Clears all icons from the database.
*/
func (self *TraitFaviconDatabase) Clear() {
	C.webkit_favicon_database_clear(self.CPointer)
	return
}

/*
Asynchronously obtains a #cairo_surface_t of the favicon for the
given page URI. It returns the cached icon if it's in the database
asynchronously waiting for the icon to be read from the database.

This is an asynchronous method. When the operation is finished, callback will
be invoked. You can then call webkit_favicon_database_get_favicon_finish()
to get the result of the operation.
*/
func (self *TraitFaviconDatabase) GetFavicon(page_uri string, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__page_uri := (*C.gchar)(unsafe.Pointer(C.CString(page_uri)))
	C.webkit_favicon_database_get_favicon(self.CPointer, __cgo__page_uri, cancellable, callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__page_uri))
	return
}

/*
Finishes an operation started with webkit_favicon_database_get_favicon().
*/
func (self *TraitFaviconDatabase) GetFaviconFinish(result *C.GAsyncResult) (return__ *C.cairo_surface_t, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_favicon_database_get_favicon_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Obtains the URI of the favicon for the given @page_uri.
*/
func (self *TraitFaviconDatabase) GetFaviconUri(page_uri string) (return__ string) {
	__cgo__page_uri := (*C.gchar)(unsafe.Pointer(C.CString(page_uri)))
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_favicon_database_get_favicon_uri(self.CPointer, __cgo__page_uri)
	C.free(unsafe.Pointer(__cgo__page_uri))
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitFileChooserRequest struct{ CPointer *C.WebKitFileChooserRequest }
type IsFileChooserRequest interface {
	GetFileChooserRequestPointer() *C.WebKitFileChooserRequest
}

func (self *TraitFileChooserRequest) GetFileChooserRequestPointer() *C.WebKitFileChooserRequest {
	return self.CPointer
}
func NewTraitFileChooserRequest(p unsafe.Pointer) *TraitFileChooserRequest {
	return &TraitFileChooserRequest{(*C.WebKitFileChooserRequest)(p)}
}

/*
Ask WebKit to cancel the request. It's important to do this in case
no selection has been made in the client, otherwise the request
won't be properly completed and the browser will keep the request
pending forever, which might cause the browser to hang.
*/
func (self *TraitFileChooserRequest) Cancel() {
	C.webkit_file_chooser_request_cancel(self.CPointer)
	return
}

/*
Get the list of MIME types the file chooser dialog should handle,
in the format specified in RFC 2046 for "media types". Its contents
depend on the value of the 'accept' attribute for HTML input
elements. This function should normally be called before presenting
the file chooser dialog to the user, to decide whether to allow the
user to select multiple files at once or only one.
*/
func (self *TraitFileChooserRequest) GetMimeTypes() (return__ **C.gchar) {
	return__ = C.webkit_file_chooser_request_get_mime_types(self.CPointer)
	return
}

/*
Get the filter currently associated with the request, ready to be
used by #GtkFileChooser. This function should normally be called
before presenting the file chooser dialog to the user, to decide
whether to apply a filter so the user would not be allowed to
select files with other MIME types.

See webkit_file_chooser_request_get_mime_types() if you are
interested in getting the list of accepted MIME types.
*/
func (self *TraitFileChooserRequest) GetMimeTypesFilter() (return__ *C.GtkFileFilter) {
	return__ = C.webkit_file_chooser_request_get_mime_types_filter(self.CPointer)
	return
}

/*
Determine whether the file chooser associated to this
#WebKitFileChooserRequest should allow selecting multiple files,
which depends on the HTML input element having a 'multiple'
attribute defined.
*/
func (self *TraitFileChooserRequest) GetSelectMultiple() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_file_chooser_request_get_select_multiple(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the list of selected files currently associated to the
request. Initially, the return value of this method contains any
files selected in previous file chooser requests for this HTML
input element. Once webkit_file_chooser_request_select_files, the
value will reflect whatever files are given.

This function should normally be called only before presenting the
file chooser dialog to the user, to decide whether to perform some
extra action, like pre-selecting the files from a previous request.
*/
func (self *TraitFileChooserRequest) GetSelectedFiles() (return__ **C.gchar) {
	return__ = C.webkit_file_chooser_request_get_selected_files(self.CPointer)
	return
}

/*
Ask WebKit to select local files for upload and complete the
request.
*/
func (self *TraitFileChooserRequest) SelectFiles(files []string) {
	__header__files := (*reflect.SliceHeader)(unsafe.Pointer(&files))
	C.webkit_file_chooser_request_select_files(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__files.Data)))
	return
}

type TraitFindController struct{ CPointer *C.WebKitFindController }
type IsFindController interface {
	GetFindControllerPointer() *C.WebKitFindController
}

func (self *TraitFindController) GetFindControllerPointer() *C.WebKitFindController {
	return self.CPointer
}
func NewTraitFindController(p unsafe.Pointer) *TraitFindController {
	return &TraitFindController{(*C.WebKitFindController)(p)}
}

/*
Counts the number of matches for @search_text found in the
#WebKitWebView with the provided @find_options. The number of
matches will be provided by the
#WebKitFindController::counted-matches signal.
*/
func (self *TraitFindController) CountMatches(search_text string, find_options uint32, max_match_count uint) {
	__cgo__search_text := (*C.gchar)(unsafe.Pointer(C.CString(search_text)))
	C.webkit_find_controller_count_matches(self.CPointer, __cgo__search_text, C.guint32(find_options), C.guint(max_match_count))
	C.free(unsafe.Pointer(__cgo__search_text))
	return
}

/*
Gets the maximum number of matches to report during a text
lookup. This number is passed as the last argument of
webkit_find_controller_search() or
webkit_find_controller_count_matches().
*/
func (self *TraitFindController) GetMaxMatchCount() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_find_controller_get_max_match_count(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets a bitmask containing the #WebKitFindOptions associated with
the current search.
*/
func (self *TraitFindController) GetOptions() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.webkit_find_controller_get_options(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Gets the text that @find_controller is currently searching
for. This text is passed to either
webkit_find_controller_search() or
webkit_find_controller_count_matches().
*/
func (self *TraitFindController) GetSearchText() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_find_controller_get_search_text(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #WebKitWebView this find controller is associated to. Do
not unref the returned instance as it belongs to the
#WebKitFindController.
*/
func (self *TraitFindController) GetWebView() (return__ *WebView) {
	var __cgo__return__ *C.WebKitWebView
	__cgo__return__ = C.webkit_find_controller_get_web_view(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Looks for @search_text in the #WebKitWebView associated with
@find_controller since the beginning of the document highlighting
up to @max_match_count matches. The outcome of the search will be
asynchronously provided by the #WebKitFindController::found-text
and #WebKitFindController::failed-to-find-text signals.

To look for the next or previous occurrences of the same text
with the same find options use webkit_find_controller_search_next()
and/or webkit_find_controller_search_previous(). The
#WebKitFindController will use the same text and options for the
following searches unless they are modified by another call to this
method.

Note that if the number of matches is higher than @max_match_count
then #WebKitFindController::found-text will report %G_MAXUINT matches
instead of the actual number.

Callers should call webkit_find_controller_search_finish() to
finish the current search operation.
*/
func (self *TraitFindController) Search(search_text string, find_options uint32, max_match_count uint) {
	__cgo__search_text := (*C.gchar)(unsafe.Pointer(C.CString(search_text)))
	C.webkit_find_controller_search(self.CPointer, __cgo__search_text, C.guint32(find_options), C.guint(max_match_count))
	C.free(unsafe.Pointer(__cgo__search_text))
	return
}

/*
Finishes a find operation started by
webkit_find_controller_search(). It will basically unhighlight
every text match found.

This method will be typically called when the search UI is
closed/hidden by the client application.
*/
func (self *TraitFindController) SearchFinish() {
	C.webkit_find_controller_search_finish(self.CPointer)
	return
}

/*
Looks for the next occurrence of the search text.

Calling this method before webkit_find_controller_search() or
webkit_find_controller_count_matches() is a programming error.
*/
func (self *TraitFindController) SearchNext() {
	C.webkit_find_controller_search_next(self.CPointer)
	return
}

/*
Looks for the previous occurrence of the search text.

Calling this method before webkit_find_controller_search() or
webkit_find_controller_count_matches() is a programming error.
*/
func (self *TraitFindController) SearchPrevious() {
	C.webkit_find_controller_search_previous(self.CPointer)
	return
}

type TraitFormSubmissionRequest struct {
	CPointer *C.WebKitFormSubmissionRequest
}
type IsFormSubmissionRequest interface {
	GetFormSubmissionRequestPointer() *C.WebKitFormSubmissionRequest
}

func (self *TraitFormSubmissionRequest) GetFormSubmissionRequestPointer() *C.WebKitFormSubmissionRequest {
	return self.CPointer
}
func NewTraitFormSubmissionRequest(p unsafe.Pointer) *TraitFormSubmissionRequest {
	return &TraitFormSubmissionRequest{(*C.WebKitFormSubmissionRequest)(p)}
}

/*
Get a #GHashTable with the values of the text fields contained in the form
associated to @request.
*/
func (self *TraitFormSubmissionRequest) GetTextFields() (return__ *C.GHashTable) {
	return__ = C.webkit_form_submission_request_get_text_fields(self.CPointer)
	return
}

/*
Continue the form submission.
*/
func (self *TraitFormSubmissionRequest) Submit() {
	C.webkit_form_submission_request_submit(self.CPointer)
	return
}

type TraitGeolocationPermissionRequest struct {
	CPointer *C.WebKitGeolocationPermissionRequest
}
type IsGeolocationPermissionRequest interface {
	GetGeolocationPermissionRequestPointer() *C.WebKitGeolocationPermissionRequest
}

func (self *TraitGeolocationPermissionRequest) GetGeolocationPermissionRequestPointer() *C.WebKitGeolocationPermissionRequest {
	return self.CPointer
}
func NewTraitGeolocationPermissionRequest(p unsafe.Pointer) *TraitGeolocationPermissionRequest {
	return &TraitGeolocationPermissionRequest{(*C.WebKitGeolocationPermissionRequest)(p)}
}

type TraitHitTestResult struct{ CPointer *C.WebKitHitTestResult }
type IsHitTestResult interface {
	GetHitTestResultPointer() *C.WebKitHitTestResult
}

func (self *TraitHitTestResult) GetHitTestResultPointer() *C.WebKitHitTestResult {
	return self.CPointer
}
func NewTraitHitTestResult(p unsafe.Pointer) *TraitHitTestResult {
	return &TraitHitTestResult{(*C.WebKitHitTestResult)(p)}
}

/*
Gets whether %WEBKIT_HIT_TEST_RESULT_CONTEXT_EDITABLE flag is present in
#WebKitHitTestResult:context.
*/
func (self *TraitHitTestResult) ContextIsEditable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_hit_test_result_context_is_editable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether %WEBKIT_HIT_TEST_RESULT_CONTEXT_IMAGE flag is present in
#WebKitHitTestResult:context.
*/
func (self *TraitHitTestResult) ContextIsImage() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_hit_test_result_context_is_image(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether %WEBKIT_HIT_TEST_RESULT_CONTEXT_LINK flag is present in
#WebKitHitTestResult:context.
*/
func (self *TraitHitTestResult) ContextIsLink() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_hit_test_result_context_is_link(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether %WEBKIT_HIT_TEST_RESULT_CONTEXT_MEDIA flag is present in
#WebKitHitTestResult:context.
*/
func (self *TraitHitTestResult) ContextIsMedia() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_hit_test_result_context_is_media(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets whether %WEBKIT_HIT_TEST_RESULT_CONTEXT_SCROLLBAR flag is present in
#WebKitHitTestResult:context.
*/
func (self *TraitHitTestResult) ContextIsScrollbar() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_hit_test_result_context_is_scrollbar(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the value of the #WebKitHitTestResult:context property.
*/
func (self *TraitHitTestResult) GetContext() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_hit_test_result_get_context(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the value of the #WebKitHitTestResult:image-uri property.
*/
func (self *TraitHitTestResult) GetImageUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_hit_test_result_get_image_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of the #WebKitHitTestResult:link-label property.
*/
func (self *TraitHitTestResult) GetLinkLabel() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_hit_test_result_get_link_label(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of the #WebKitHitTestResult:link-title property.
*/
func (self *TraitHitTestResult) GetLinkTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_hit_test_result_get_link_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of the #WebKitHitTestResult:link-uri property.
*/
func (self *TraitHitTestResult) GetLinkUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_hit_test_result_get_link_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of the #WebKitHitTestResult:media-uri property.
*/
func (self *TraitHitTestResult) GetMediaUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_hit_test_result_get_media_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitNavigationPolicyDecision struct {
	CPointer *C.WebKitNavigationPolicyDecision
}
type IsNavigationPolicyDecision interface {
	GetNavigationPolicyDecisionPointer() *C.WebKitNavigationPolicyDecision
}

func (self *TraitNavigationPolicyDecision) GetNavigationPolicyDecisionPointer() *C.WebKitNavigationPolicyDecision {
	return self.CPointer
}
func NewTraitNavigationPolicyDecision(p unsafe.Pointer) *TraitNavigationPolicyDecision {
	return &TraitNavigationPolicyDecision{(*C.WebKitNavigationPolicyDecision)(p)}
}

/*
Gets the value of the #WebKitNavigationPolicyDecision:frame-name property.
*/
func (self *TraitNavigationPolicyDecision) GetFrameName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_navigation_policy_decision_get_frame_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of the #WebKitNavigationPolicyDecision:modifiers property.
*/
func (self *TraitNavigationPolicyDecision) GetModifiers() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_navigation_policy_decision_get_modifiers(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the value of the #WebKitNavigationPolicyDecision:mouse-button property.
*/
func (self *TraitNavigationPolicyDecision) GetMouseButton() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_navigation_policy_decision_get_mouse_button(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Gets the value of the #WebKitNavigationPolicyDecision:navigation-type property.
*/
func (self *TraitNavigationPolicyDecision) GetNavigationType() (return__ C.WebKitNavigationType) {
	return__ = C.webkit_navigation_policy_decision_get_navigation_type(self.CPointer)
	return
}

/*
Gets the value of the #WebKitNavigationPolicyDecision:request property.
*/
func (self *TraitNavigationPolicyDecision) GetRequest() (return__ *URIRequest) {
	var __cgo__return__ *C.WebKitURIRequest
	__cgo__return__ = C.webkit_navigation_policy_decision_get_request(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewURIRequestFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

type TraitPlugin struct{ CPointer *C.WebKitPlugin }
type IsPlugin interface {
	GetPluginPointer() *C.WebKitPlugin
}

func (self *TraitPlugin) GetPluginPointer() *C.WebKitPlugin {
	return self.CPointer
}
func NewTraitPlugin(p unsafe.Pointer) *TraitPlugin {
	return &TraitPlugin{(*C.WebKitPlugin)(p)}
}
func (self *TraitPlugin) GetDescription() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_plugin_get_description(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get information about MIME types handled by the plugin,
as a list of #WebKitMimeInfo.
*/
func (self *TraitPlugin) GetMimeInfoList() (return__ *C.GList) {
	return__ = C.webkit_plugin_get_mime_info_list(self.CPointer)
	return
}

func (self *TraitPlugin) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_plugin_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func (self *TraitPlugin) GetPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_plugin_get_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitPolicyDecision struct{ CPointer *C.WebKitPolicyDecision }
type IsPolicyDecision interface {
	GetPolicyDecisionPointer() *C.WebKitPolicyDecision
}

func (self *TraitPolicyDecision) GetPolicyDecisionPointer() *C.WebKitPolicyDecision {
	return self.CPointer
}
func NewTraitPolicyDecision(p unsafe.Pointer) *TraitPolicyDecision {
	return &TraitPolicyDecision{(*C.WebKitPolicyDecision)(p)}
}

/*
Spawn a download from this decision.
*/
func (self *TraitPolicyDecision) Download() {
	C.webkit_policy_decision_download(self.CPointer)
	return
}

/*
Ignore the action which triggerd this decision. For instance, for a
#WebKitResponsePolicyDecision, this would cancel the request.
*/
func (self *TraitPolicyDecision) Ignore() {
	C.webkit_policy_decision_ignore(self.CPointer)
	return
}

/*
Accept the action which triggerd this decision.
*/
func (self *TraitPolicyDecision) Use() {
	C.webkit_policy_decision_use(self.CPointer)
	return
}

type TraitPrintOperation struct{ CPointer *C.WebKitPrintOperation }
type IsPrintOperation interface {
	GetPrintOperationPointer() *C.WebKitPrintOperation
}

func (self *TraitPrintOperation) GetPrintOperationPointer() *C.WebKitPrintOperation {
	return self.CPointer
}
func NewTraitPrintOperation(p unsafe.Pointer) *TraitPrintOperation {
	return &TraitPrintOperation{(*C.WebKitPrintOperation)(p)}
}

/*
Return the current page setup of @print_operation. It returns %NULL until
either webkit_print_operation_set_print_settings() or webkit_print_operation_run_dialog()
have been called.
*/
func (self *TraitPrintOperation) GetPageSetup() (return__ *C.GtkPageSetup) {
	return__ = C.webkit_print_operation_get_page_setup(self.CPointer)
	return
}

/*
Return the current print settings of @print_operation. It returns %NULL until
either webkit_print_operation_set_print_settings() or webkit_print_operation_run_dialog()
have been called.
*/
func (self *TraitPrintOperation) GetPrintSettings() (return__ *C.GtkPrintSettings) {
	return__ = C.webkit_print_operation_get_print_settings(self.CPointer)
	return
}

/*
Start a print operation using current print settings and page setup
without showing the print dialog. If either print settings or page setup
are not set with webkit_print_operation_set_print_settings() and
webkit_print_operation_set_page_setup(), the default options will be used
and the print job will be sent to the default printer.
The #WebKitPrintOperation::finished signal is emitted when the printing
operation finishes. If an error occurs while printing the signal
#WebKitPrintOperation::failed is emitted before #WebKitPrintOperation::finished.
*/
func (self *TraitPrintOperation) Print() {
	C.webkit_print_operation_print(self.CPointer)
	return
}

/*
Run the print dialog and start printing using the options selected by
the user. This method returns when the print dialog is closed.
If the print dialog is cancelled %WEBKIT_PRINT_OPERATION_RESPONSE_CANCEL
is returned. If the user clicks on the print button, %WEBKIT_PRINT_OPERATION_RESPONSE_PRINT
is returned and the print operation starts. In this case, the #WebKitPrintOperation::finished
signal is emitted when the operation finishes. If an error occurs while printing, the signal
#WebKitPrintOperation::failed is emitted before #WebKitPrintOperation::finished.
If the print dialog is not cancelled current print settings and page setup of @print_operation
are updated with options selected by the user when Print button is pressed in print dialog.
You can get the updated print settings and page setup by calling
webkit_print_operation_get_print_settings() and webkit_print_operation_get_page_setup()
after this method.
*/
func (self *TraitPrintOperation) RunDialog(parent *C.GtkWindow) (return__ C.WebKitPrintOperationResponse) {
	return__ = C.webkit_print_operation_run_dialog(self.CPointer, parent)
	return
}

/*
Set the current page setup of @print_operation. Current page setup is used for the
initial values of the print dialog when webkit_print_operation_run_dialog() is called.
*/
func (self *TraitPrintOperation) SetPageSetup(page_setup *C.GtkPageSetup) {
	C.webkit_print_operation_set_page_setup(self.CPointer, page_setup)
	return
}

/*
Set the current print settings of @print_operation. Current print settings are used for
the initial values of the print dialog when webkit_print_operation_run_dialog() is called.
*/
func (self *TraitPrintOperation) SetPrintSettings(print_settings *C.GtkPrintSettings) {
	C.webkit_print_operation_set_print_settings(self.CPointer, print_settings)
	return
}

type TraitResponsePolicyDecision struct {
	CPointer *C.WebKitResponsePolicyDecision
}
type IsResponsePolicyDecision interface {
	GetResponsePolicyDecisionPointer() *C.WebKitResponsePolicyDecision
}

func (self *TraitResponsePolicyDecision) GetResponsePolicyDecisionPointer() *C.WebKitResponsePolicyDecision {
	return self.CPointer
}
func NewTraitResponsePolicyDecision(p unsafe.Pointer) *TraitResponsePolicyDecision {
	return &TraitResponsePolicyDecision{(*C.WebKitResponsePolicyDecision)(p)}
}

/*
Gets the value of the #WebKitResponsePolicyDecision:request property.
*/
func (self *TraitResponsePolicyDecision) GetRequest() (return__ *URIRequest) {
	var __cgo__return__ *C.WebKitURIRequest
	__cgo__return__ = C.webkit_response_policy_decision_get_request(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewURIRequestFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the value of the #WebKitResponsePolicyDecision:response property.
*/
func (self *TraitResponsePolicyDecision) GetResponse() (return__ *URIResponse) {
	var __cgo__return__ *C.WebKitURIResponse
	__cgo__return__ = C.webkit_response_policy_decision_get_response(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewURIResponseFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets whether the MIME type of the response can be displayed in the #WebKitWebView
that triggered this policy decision request. See also webkit_web_view_can_show_mime_type().
*/
func (self *TraitResponsePolicyDecision) IsMimeTypeSupported() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_response_policy_decision_is_mime_type_supported(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitSecurityManager struct{ CPointer *C.WebKitSecurityManager }
type IsSecurityManager interface {
	GetSecurityManagerPointer() *C.WebKitSecurityManager
}

func (self *TraitSecurityManager) GetSecurityManagerPointer() *C.WebKitSecurityManager {
	return self.CPointer
}
func NewTraitSecurityManager(p unsafe.Pointer) *TraitSecurityManager {
	return &TraitSecurityManager{(*C.WebKitSecurityManager)(p)}
}

/*
Register @scheme as a CORS (Cross-origin resource sharing) enabled scheme.
This means that CORS requests are allowed. See W3C CORS specification
http://www.w3.org/TR/cors/.
*/
func (self *TraitSecurityManager) RegisterUriSchemeAsCorsEnabled(scheme string) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.webkit_security_manager_register_uri_scheme_as_cors_enabled(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

/*
Register @scheme as a display isolated scheme. This means that pages cannot
display these URIs unless they are from the same scheme.
*/
func (self *TraitSecurityManager) RegisterUriSchemeAsDisplayIsolated(scheme string) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.webkit_security_manager_register_uri_scheme_as_display_isolated(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

/*
Register @scheme as an empty document scheme. This means that
they are allowd to commit synchronously.
*/
func (self *TraitSecurityManager) RegisterUriSchemeAsEmptyDocument(scheme string) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.webkit_security_manager_register_uri_scheme_as_empty_document(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

/*
Register @scheme as a local scheme. This means that other non-local pages
cannot link to or access URIs of this scheme.
*/
func (self *TraitSecurityManager) RegisterUriSchemeAsLocal(scheme string) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.webkit_security_manager_register_uri_scheme_as_local(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

/*
Register @scheme as a no-access scheme. This means that pages loaded
with this URI scheme cannot access pages loaded with any other URI scheme.
*/
func (self *TraitSecurityManager) RegisterUriSchemeAsNoAccess(scheme string) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.webkit_security_manager_register_uri_scheme_as_no_access(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

/*
Register @scheme as a secure scheme. This means that mixed
content warnings won't be generated for this scheme when
included by an HTTPS page.
*/
func (self *TraitSecurityManager) RegisterUriSchemeAsSecure(scheme string) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.webkit_security_manager_register_uri_scheme_as_secure(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

/*
Whether @scheme is considered as a CORS enabled scheme.
See also webkit_security_manager_register_uri_scheme_as_cors_enabled().
*/
func (self *TraitSecurityManager) UriSchemeIsCorsEnabled(scheme string) (return__ bool) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_security_manager_uri_scheme_is_cors_enabled(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether @scheme is considered as a display isolated scheme.
See also webkit_security_manager_register_uri_scheme_as_display_isolated().
*/
func (self *TraitSecurityManager) UriSchemeIsDisplayIsolated(scheme string) (return__ bool) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_security_manager_uri_scheme_is_display_isolated(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether @scheme is considered as an empty document scheme.
See also webkit_security_manager_register_uri_scheme_as_empty_document().
*/
func (self *TraitSecurityManager) UriSchemeIsEmptyDocument(scheme string) (return__ bool) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_security_manager_uri_scheme_is_empty_document(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether @scheme is considered as a local scheme.
See also webkit_security_manager_register_uri_scheme_as_local().
*/
func (self *TraitSecurityManager) UriSchemeIsLocal(scheme string) (return__ bool) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_security_manager_uri_scheme_is_local(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether @scheme is considered as a no-access scheme.
See also webkit_security_manager_register_uri_scheme_as_no_access().
*/
func (self *TraitSecurityManager) UriSchemeIsNoAccess(scheme string) (return__ bool) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_security_manager_uri_scheme_is_no_access(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether @scheme is considered as a secure scheme.
See also webkit_security_manager_register_uri_scheme_as_secure().
*/
func (self *TraitSecurityManager) UriSchemeIsSecure(scheme string) (return__ bool) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_security_manager_uri_scheme_is_secure(self.CPointer, __cgo__scheme)
	C.free(unsafe.Pointer(__cgo__scheme))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

type TraitSettings struct{ CPointer *C.WebKitSettings }
type IsSettings interface {
	GetSettingsPointer() *C.WebKitSettings
}

func (self *TraitSettings) GetSettingsPointer() *C.WebKitSettings {
	return self.CPointer
}
func NewTraitSettings(p unsafe.Pointer) *TraitSettings {
	return &TraitSettings{(*C.WebKitSettings)(p)}
}

/*
Get the #WebKitSettings:allow-modal-dialogs property.
*/
func (self *TraitSettings) GetAllowModalDialogs() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_allow_modal_dialogs(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:auto-load-images property.
*/
func (self *TraitSettings) GetAutoLoadImages() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_auto_load_images(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the #WebKitSettings:cursive-font-family property.
*/
func (self *TraitSettings) GetCursiveFontFamily() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_cursive_font_family(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #WebKitSettings:default-charset property.
*/
func (self *TraitSettings) GetDefaultCharset() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_default_charset(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func (self *TraitSettings) GetDefaultFontFamily() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_default_font_family(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #WebKitSettings:default-font-size property.
*/
func (self *TraitSettings) GetDefaultFontSize() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.webkit_settings_get_default_font_size(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Gets the #WebKitSettings:default-monospace-font-size property.
*/
func (self *TraitSettings) GetDefaultMonospaceFontSize() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.webkit_settings_get_default_monospace_font_size(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Get the #WebKitSettings:draw-compositing-indicators property.
*/
func (self *TraitSettings) GetDrawCompositingIndicators() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_draw_compositing_indicators(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-accelerated-2d-canvas property.
*/
func (self *TraitSettings) GetEnableAccelerated2dCanvas() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_accelerated_2d_canvas(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-caret-browsing property.
*/
func (self *TraitSettings) GetEnableCaretBrowsing() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_caret_browsing(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-developer-extras property.
*/
func (self *TraitSettings) GetEnableDeveloperExtras() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_developer_extras(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-dns-prefetching property.
*/
func (self *TraitSettings) GetEnableDnsPrefetching() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_dns_prefetching(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-frame-flattening property.
*/
func (self *TraitSettings) GetEnableFrameFlattening() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_frame_flattening(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-fullscreen property.
*/
func (self *TraitSettings) GetEnableFullscreen() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_fullscreen(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-html5-database property.
*/
func (self *TraitSettings) GetEnableHtml5Database() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_html5_database(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-html5-local-storage property.
*/
func (self *TraitSettings) GetEnableHtml5LocalStorage() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_html5_local_storage(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-hyperlink-auditing property.
*/
func (self *TraitSettings) GetEnableHyperlinkAuditing() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_hyperlink_auditing(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-java property.
*/
func (self *TraitSettings) GetEnableJava() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_java(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-javascript property.
*/
func (self *TraitSettings) GetEnableJavascript() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_javascript(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-media-stream property.
*/
func (self *TraitSettings) GetEnableMediaStream() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_media_stream(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-mediasource property.
*/
func (self *TraitSettings) GetEnableMediasource() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_mediasource(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-offline-web-application-cache property.
*/
func (self *TraitSettings) GetEnableOfflineWebApplicationCache() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_offline_web_application_cache(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-page-cache property.
*/
func (self *TraitSettings) GetEnablePageCache() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_page_cache(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-plugins property.
*/
func (self *TraitSettings) GetEnablePlugins() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_plugins(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-private-browsing property.
*/
func (self *TraitSettings) GetEnablePrivateBrowsing() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_private_browsing(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-resizable-text-areas property.
*/
func (self *TraitSettings) GetEnableResizableTextAreas() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_resizable_text_areas(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-site-specific-quirks property.
*/
func (self *TraitSettings) GetEnableSiteSpecificQuirks() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_site_specific_quirks(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-smooth-scrolling property.
*/
func (self *TraitSettings) GetEnableSmoothScrolling() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_smooth_scrolling(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-spatial-navigation property.
*/
func (self *TraitSettings) GetEnableSpatialNavigation() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_spatial_navigation(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-tabs-to-links property.
*/
func (self *TraitSettings) GetEnableTabsToLinks() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_tabs_to_links(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-webaudio property.
*/
func (self *TraitSettings) GetEnableWebaudio() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_webaudio(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-webgl property.
*/
func (self *TraitSettings) GetEnableWebgl() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_webgl(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-write-console-messages-to-stdout property.
*/
func (self *TraitSettings) GetEnableWriteConsoleMessagesToStdout() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_write_console_messages_to_stdout(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:enable-xss-auditor property.
*/
func (self *TraitSettings) GetEnableXssAuditor() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_enable_xss_auditor(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the #WebKitSettings:fantasy-font-family property.
*/
func (self *TraitSettings) GetFantasyFontFamily() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_fantasy_font_family(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the #WebKitSettings:javascript-can-access-clipboard property.
*/
func (self *TraitSettings) GetJavascriptCanAccessClipboard() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_javascript_can_access_clipboard(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:javascript-can-open-windows-automatically property.
*/
func (self *TraitSettings) GetJavascriptCanOpenWindowsAutomatically() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_javascript_can_open_windows_automatically(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:load-icons-ignoring-image-load-setting property.
*/
func (self *TraitSettings) GetLoadIconsIgnoringImageLoadSetting() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_load_icons_ignoring_image_load_setting(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:media-playback-allows-inline property.
*/
func (self *TraitSettings) GetMediaPlaybackAllowsInline() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_media_playback_allows_inline(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the #WebKitSettings:media-playback-requires-user-gesture property.
*/
func (self *TraitSettings) GetMediaPlaybackRequiresUserGesture() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_media_playback_requires_user_gesture(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the #WebKitSettings:minimum-font-size property.
*/
func (self *TraitSettings) GetMinimumFontSize() (return__ uint32) {
	var __cgo__return__ C.guint32
	__cgo__return__ = C.webkit_settings_get_minimum_font_size(self.CPointer)
	return__ = uint32(__cgo__return__)
	return
}

/*
Gets the #WebKitSettings:monospace-font-family property.
*/
func (self *TraitSettings) GetMonospaceFontFamily() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_monospace_font_family(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #WebKitSettings:pictograph-font-family property.
*/
func (self *TraitSettings) GetPictographFontFamily() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_pictograph_font_family(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the #WebKitSettings:print-backgrounds property.
*/
func (self *TraitSettings) GetPrintBackgrounds() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_print_backgrounds(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Gets the #WebKitSettings:sans-serif-font-family property.
*/
func (self *TraitSettings) GetSansSerifFontFamily() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_sans_serif_font_family(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #WebKitSettings:serif-font-family property.
*/
func (self *TraitSettings) GetSerifFontFamily() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_serif_font_family(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the #WebKitSettings:user-agent property.
*/
func (self *TraitSettings) GetUserAgent() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_settings_get_user_agent(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the #WebKitSettings:zoom-text-only property.
*/
func (self *TraitSettings) GetZoomTextOnly() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_settings_get_zoom_text_only(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Set the #WebKitSettings:allow-modal-dialogs property.
*/
func (self *TraitSettings) SetAllowModalDialogs(allowed bool) {
	__cgo__allowed := C.gboolean(0)
	if allowed {
		__cgo__allowed = C.gboolean(1)
	}
	C.webkit_settings_set_allow_modal_dialogs(self.CPointer, __cgo__allowed)
	return
}

/*
Set the #WebKitSettings:auto-load-images property.
*/
func (self *TraitSettings) SetAutoLoadImages(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_auto_load_images(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:cursive-font-family property.
*/
func (self *TraitSettings) SetCursiveFontFamily(cursive_font_family string) {
	__cgo__cursive_font_family := (*C.gchar)(unsafe.Pointer(C.CString(cursive_font_family)))
	C.webkit_settings_set_cursive_font_family(self.CPointer, __cgo__cursive_font_family)
	C.free(unsafe.Pointer(__cgo__cursive_font_family))
	return
}

/*
Set the #WebKitSettings:default-charset property.
*/
func (self *TraitSettings) SetDefaultCharset(default_charset string) {
	__cgo__default_charset := (*C.gchar)(unsafe.Pointer(C.CString(default_charset)))
	C.webkit_settings_set_default_charset(self.CPointer, __cgo__default_charset)
	C.free(unsafe.Pointer(__cgo__default_charset))
	return
}

/*
Set the #WebKitSettings:default-font-family property.
*/
func (self *TraitSettings) SetDefaultFontFamily(default_font_family string) {
	__cgo__default_font_family := (*C.gchar)(unsafe.Pointer(C.CString(default_font_family)))
	C.webkit_settings_set_default_font_family(self.CPointer, __cgo__default_font_family)
	C.free(unsafe.Pointer(__cgo__default_font_family))
	return
}

/*
Set the #WebKitSettings:default-font-size property.
*/
func (self *TraitSettings) SetDefaultFontSize(font_size uint32) {
	C.webkit_settings_set_default_font_size(self.CPointer, C.guint32(font_size))
	return
}

/*
Set the #WebKitSettings:default-monospace-font-size property.
*/
func (self *TraitSettings) SetDefaultMonospaceFontSize(font_size uint32) {
	C.webkit_settings_set_default_monospace_font_size(self.CPointer, C.guint32(font_size))
	return
}

/*
Set the #WebKitSettings:draw-compositing-indicators property.
*/
func (self *TraitSettings) SetDrawCompositingIndicators(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_draw_compositing_indicators(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-accelerated-2d-canvas property.
*/
func (self *TraitSettings) SetEnableAccelerated2dCanvas(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_accelerated_2d_canvas(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-caret-browsing property.
*/
func (self *TraitSettings) SetEnableCaretBrowsing(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_caret_browsing(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-developer-extras property.
*/
func (self *TraitSettings) SetEnableDeveloperExtras(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_developer_extras(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-dns-prefetching property.
*/
func (self *TraitSettings) SetEnableDnsPrefetching(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_dns_prefetching(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-frame-flattening property.
*/
func (self *TraitSettings) SetEnableFrameFlattening(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_frame_flattening(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-fullscreen property.
*/
func (self *TraitSettings) SetEnableFullscreen(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_fullscreen(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-html5-database property.
*/
func (self *TraitSettings) SetEnableHtml5Database(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_html5_database(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-html5-local-storage property.
*/
func (self *TraitSettings) SetEnableHtml5LocalStorage(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_html5_local_storage(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-hyperlink-auditing property.
*/
func (self *TraitSettings) SetEnableHyperlinkAuditing(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_hyperlink_auditing(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-java property.
*/
func (self *TraitSettings) SetEnableJava(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_java(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-javascript property.
*/
func (self *TraitSettings) SetEnableJavascript(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_javascript(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-media-stream property.
*/
func (self *TraitSettings) SetEnableMediaStream(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_media_stream(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-mediasource property.
*/
func (self *TraitSettings) SetEnableMediasource(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_mediasource(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-offline-web-application-cache property.
*/
func (self *TraitSettings) SetEnableOfflineWebApplicationCache(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_offline_web_application_cache(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-page-cache property.
*/
func (self *TraitSettings) SetEnablePageCache(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_page_cache(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-plugins property.
*/
func (self *TraitSettings) SetEnablePlugins(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_plugins(self.CPointer, __cgo__enabled)
	return
}

func (self *TraitSettings) SetEnablePrivateBrowsing(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_private_browsing(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-resizable-text-areas property.
*/
func (self *TraitSettings) SetEnableResizableTextAreas(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_resizable_text_areas(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-site-specific-quirks property.
*/
func (self *TraitSettings) SetEnableSiteSpecificQuirks(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_site_specific_quirks(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-smooth-scrolling property.
*/
func (self *TraitSettings) SetEnableSmoothScrolling(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_smooth_scrolling(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-spatial-navigation property.
*/
func (self *TraitSettings) SetEnableSpatialNavigation(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_spatial_navigation(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-tabs-to-links property.
*/
func (self *TraitSettings) SetEnableTabsToLinks(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_tabs_to_links(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-webaudio property.
*/
func (self *TraitSettings) SetEnableWebaudio(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_webaudio(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-webgl property.
*/
func (self *TraitSettings) SetEnableWebgl(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_webgl(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-write-console-messages-to-stdout property.
*/
func (self *TraitSettings) SetEnableWriteConsoleMessagesToStdout(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_write_console_messages_to_stdout(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:enable-xss-auditor property.
*/
func (self *TraitSettings) SetEnableXssAuditor(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_enable_xss_auditor(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:fantasy-font-family property.
*/
func (self *TraitSettings) SetFantasyFontFamily(fantasy_font_family string) {
	__cgo__fantasy_font_family := (*C.gchar)(unsafe.Pointer(C.CString(fantasy_font_family)))
	C.webkit_settings_set_fantasy_font_family(self.CPointer, __cgo__fantasy_font_family)
	C.free(unsafe.Pointer(__cgo__fantasy_font_family))
	return
}

/*
Set the #WebKitSettings:javascript-can-access-clipboard property.
*/
func (self *TraitSettings) SetJavascriptCanAccessClipboard(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_javascript_can_access_clipboard(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:javascript-can-open-windows-automatically property.
*/
func (self *TraitSettings) SetJavascriptCanOpenWindowsAutomatically(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_javascript_can_open_windows_automatically(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:load-icons-ignoring-image-load-setting property.
*/
func (self *TraitSettings) SetLoadIconsIgnoringImageLoadSetting(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_load_icons_ignoring_image_load_setting(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:media-playback-allows-inline property.
*/
func (self *TraitSettings) SetMediaPlaybackAllowsInline(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_media_playback_allows_inline(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:media-playback-requires-user-gesture property.
*/
func (self *TraitSettings) SetMediaPlaybackRequiresUserGesture(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_settings_set_media_playback_requires_user_gesture(self.CPointer, __cgo__enabled)
	return
}

/*
Set the #WebKitSettings:minimum-font-size property.
*/
func (self *TraitSettings) SetMinimumFontSize(font_size uint32) {
	C.webkit_settings_set_minimum_font_size(self.CPointer, C.guint32(font_size))
	return
}

/*
Set the #WebKitSettings:monospace-font-family property.
*/
func (self *TraitSettings) SetMonospaceFontFamily(monospace_font_family string) {
	__cgo__monospace_font_family := (*C.gchar)(unsafe.Pointer(C.CString(monospace_font_family)))
	C.webkit_settings_set_monospace_font_family(self.CPointer, __cgo__monospace_font_family)
	C.free(unsafe.Pointer(__cgo__monospace_font_family))
	return
}

/*
Set the #WebKitSettings:pictograph-font-family property.
*/
func (self *TraitSettings) SetPictographFontFamily(pictograph_font_family string) {
	__cgo__pictograph_font_family := (*C.gchar)(unsafe.Pointer(C.CString(pictograph_font_family)))
	C.webkit_settings_set_pictograph_font_family(self.CPointer, __cgo__pictograph_font_family)
	C.free(unsafe.Pointer(__cgo__pictograph_font_family))
	return
}

/*
Set the #WebKitSettings:print-backgrounds property.
*/
func (self *TraitSettings) SetPrintBackgrounds(print_backgrounds bool) {
	__cgo__print_backgrounds := C.gboolean(0)
	if print_backgrounds {
		__cgo__print_backgrounds = C.gboolean(1)
	}
	C.webkit_settings_set_print_backgrounds(self.CPointer, __cgo__print_backgrounds)
	return
}

/*
Set the #WebKitSettings:sans-serif-font-family property.
*/
func (self *TraitSettings) SetSansSerifFontFamily(sans_serif_font_family string) {
	__cgo__sans_serif_font_family := (*C.gchar)(unsafe.Pointer(C.CString(sans_serif_font_family)))
	C.webkit_settings_set_sans_serif_font_family(self.CPointer, __cgo__sans_serif_font_family)
	C.free(unsafe.Pointer(__cgo__sans_serif_font_family))
	return
}

/*
Set the #WebKitSettings:serif-font-family property.
*/
func (self *TraitSettings) SetSerifFontFamily(serif_font_family string) {
	__cgo__serif_font_family := (*C.gchar)(unsafe.Pointer(C.CString(serif_font_family)))
	C.webkit_settings_set_serif_font_family(self.CPointer, __cgo__serif_font_family)
	C.free(unsafe.Pointer(__cgo__serif_font_family))
	return
}

/*
Set the #WebKitSettings:user-agent property.
*/
func (self *TraitSettings) SetUserAgent(user_agent string) {
	__cgo__user_agent := (*C.gchar)(unsafe.Pointer(C.CString(user_agent)))
	C.webkit_settings_set_user_agent(self.CPointer, __cgo__user_agent)
	C.free(unsafe.Pointer(__cgo__user_agent))
	return
}

/*
Set the #WebKitSettings:user-agent property by appending the application details to the default user
agent. If no application name or version is given, the default user agent used will be used. If only
the version is given, the default engine version is used with the given application name.
*/
func (self *TraitSettings) SetUserAgentWithApplicationDetails(application_name string, application_version string) {
	__cgo__application_name := (*C.gchar)(unsafe.Pointer(C.CString(application_name)))
	__cgo__application_version := (*C.gchar)(unsafe.Pointer(C.CString(application_version)))
	C.webkit_settings_set_user_agent_with_application_details(self.CPointer, __cgo__application_name, __cgo__application_version)
	C.free(unsafe.Pointer(__cgo__application_name))
	C.free(unsafe.Pointer(__cgo__application_version))
	return
}

/*
Set the #WebKitSettings:zoom-text-only property.
*/
func (self *TraitSettings) SetZoomTextOnly(zoom_text_only bool) {
	__cgo__zoom_text_only := C.gboolean(0)
	if zoom_text_only {
		__cgo__zoom_text_only = C.gboolean(1)
	}
	C.webkit_settings_set_zoom_text_only(self.CPointer, __cgo__zoom_text_only)
	return
}

type TraitURIRequest struct{ CPointer *C.WebKitURIRequest }
type IsURIRequest interface {
	GetURIRequestPointer() *C.WebKitURIRequest
}

func (self *TraitURIRequest) GetURIRequestPointer() *C.WebKitURIRequest {
	return self.CPointer
}
func NewTraitURIRequest(p unsafe.Pointer) *TraitURIRequest {
	return &TraitURIRequest{(*C.WebKitURIRequest)(p)}
}

/*
Get the HTTP headers of a #WebKitURIRequest as a #SoupMessageHeaders.
*/
func (self *TraitURIRequest) GetHttpHeaders() (return__ *C.SoupMessageHeaders) {
	return__ = C.webkit_uri_request_get_http_headers(self.CPointer)
	return
}

func (self *TraitURIRequest) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_uri_request_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Set the URI of @request
*/
func (self *TraitURIRequest) SetUri(uri string) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	C.webkit_uri_request_set_uri(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return
}

type TraitURIResponse struct{ CPointer *C.WebKitURIResponse }
type IsURIResponse interface {
	GetURIResponsePointer() *C.WebKitURIResponse
}

func (self *TraitURIResponse) GetURIResponsePointer() *C.WebKitURIResponse {
	return self.CPointer
}
func NewTraitURIResponse(p unsafe.Pointer) *TraitURIResponse {
	return &TraitURIResponse{(*C.WebKitURIResponse)(p)}
}

/*
Get the expected content length of the #WebKitURIResponse. It can
be 0 if the server provided an incorrect or missing Content-Length.
*/
func (self *TraitURIResponse) GetContentLength() (return__ uint64) {
	var __cgo__return__ C.guint64
	__cgo__return__ = C.webkit_uri_response_get_content_length(self.CPointer)
	return__ = uint64(__cgo__return__)
	return
}

func (self *TraitURIResponse) GetMimeType() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_uri_response_get_mime_type(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the status code of the #WebKitURIResponse as returned by
the server. It will normally be a #SoupKnownStatusCode, for
example %SOUP_STATUS_OK, though the server can respond with any
unsigned integer.
*/
func (self *TraitURIResponse) GetStatusCode() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_uri_response_get_status_code(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Get the suggested filename for @response, as specified by
the 'Content-Disposition' HTTP header, or %NULL if it's not
present.
*/
func (self *TraitURIResponse) GetSuggestedFilename() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_uri_response_get_suggested_filename(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

func (self *TraitURIResponse) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_uri_response_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitURISchemeRequest struct{ CPointer *C.WebKitURISchemeRequest }
type IsURISchemeRequest interface {
	GetURISchemeRequestPointer() *C.WebKitURISchemeRequest
}

func (self *TraitURISchemeRequest) GetURISchemeRequestPointer() *C.WebKitURISchemeRequest {
	return self.CPointer
}
func NewTraitURISchemeRequest(p unsafe.Pointer) *TraitURISchemeRequest {
	return &TraitURISchemeRequest{(*C.WebKitURISchemeRequest)(p)}
}

/*
Finish a #WebKitURISchemeRequest by setting the contents of the request and its mime type.
*/
func (self *TraitURISchemeRequest) Finish(stream *C.GInputStream, stream_length int64, mime_type string) {
	__cgo__mime_type := (*C.gchar)(unsafe.Pointer(C.CString(mime_type)))
	C.webkit_uri_scheme_request_finish(self.CPointer, stream, C.gint64(stream_length), __cgo__mime_type)
	C.free(unsafe.Pointer(__cgo__mime_type))
	return
}

/*
Finish a #WebKitURISchemeRequest with a #GError.
*/
func (self *TraitURISchemeRequest) FinishError(error_ *C.GError) {
	C.webkit_uri_scheme_request_finish_error(self.CPointer, error_)
	return
}

/*
Get the URI path of @request
*/
func (self *TraitURISchemeRequest) GetPath() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_uri_scheme_request_get_path(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the URI scheme of @request
*/
func (self *TraitURISchemeRequest) GetScheme() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_uri_scheme_request_get_scheme(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the URI of @request
*/
func (self *TraitURISchemeRequest) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_uri_scheme_request_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the #WebKitWebView that initiated the request.
*/
func (self *TraitURISchemeRequest) GetWebView() (return__ *WebView) {
	var __cgo__return__ *C.WebKitWebView
	__cgo__return__ = C.webkit_uri_scheme_request_get_web_view(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebViewFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

type TraitWebContext struct{ CPointer *C.WebKitWebContext }
type IsWebContext interface {
	GetWebContextPointer() *C.WebKitWebContext
}

func (self *TraitWebContext) GetWebContextPointer() *C.WebKitWebContext {
	return self.CPointer
}
func NewTraitWebContext(p unsafe.Pointer) *TraitWebContext {
	return &TraitWebContext{(*C.WebKitWebContext)(p)}
}

/*
Ignore further TLS errors on the @host for the certificate present in @info.
*/
func (self *TraitWebContext) AllowTlsCertificateForHost(info *C.WebKitCertificateInfo, host string) {
	__cgo__host := (*C.gchar)(unsafe.Pointer(C.CString(host)))
	C.webkit_web_context_allow_tls_certificate_for_host(self.CPointer, info, __cgo__host)
	C.free(unsafe.Pointer(__cgo__host))
	return
}

/*
Clears all resources currently cached.
See also webkit_web_context_set_cache_model().
*/
func (self *TraitWebContext) ClearCache() {
	C.webkit_web_context_clear_cache(self.CPointer)
	return
}

/*
Requests downloading of the specified URI string. The download operation
will not be associated to any #WebKitWebView, if you are interested in
starting a download from a particular #WebKitWebView use
webkit_web_view_download_uri() instead.
*/
func (self *TraitWebContext) DownloadUri(uri string) (return__ *Download) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	var __cgo__return__ *C.WebKitDownload
	__cgo__return__ = C.webkit_web_context_download_uri(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	if __cgo__return__ != nil {
		return__ = NewDownloadFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the current cache model. For more information about this
value check the documentation of the function
webkit_web_context_set_cache_model().
*/
func (self *TraitWebContext) GetCacheModel() (return__ C.WebKitCacheModel) {
	return__ = C.webkit_web_context_get_cache_model(self.CPointer)
	return
}

/*
Get the #WebKitCookieManager of @context.
*/
func (self *TraitWebContext) GetCookieManager() (return__ *CookieManager) {
	var __cgo__return__ *C.WebKitCookieManager
	__cgo__return__ = C.webkit_web_context_get_cookie_manager(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewCookieManagerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the #WebKitFaviconDatabase associated with @context.

To initialize the database you need to call
webkit_web_context_set_favicon_database_directory().
*/
func (self *TraitWebContext) GetFaviconDatabase() (return__ *FaviconDatabase) {
	var __cgo__return__ *C.WebKitFaviconDatabase
	__cgo__return__ = C.webkit_web_context_get_favicon_database(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewFaviconDatabaseFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the directory path being used to store the favicons database
for @context, or %NULL if
webkit_web_context_set_favicon_database_directory() hasn't been
called yet.

This function will always return the same path after having called
webkit_web_context_set_favicon_database_directory() for the first
time.
*/
func (self *TraitWebContext) GetFaviconDatabaseDirectory() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_web_context_get_favicon_database_directory(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Asynchronously get the list of installed plugins.

When the operation is finished, @callback will be called. You can then call
webkit_web_context_get_plugins_finish() to get the result of the operation.
*/
func (self *TraitWebContext) GetPlugins(cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.webkit_web_context_get_plugins(self.CPointer, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous operation started with webkit_web_context_get_plugins.
*/
func (self *TraitWebContext) GetPluginsFinish(result *C.GAsyncResult) (return__ *C.GList, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_web_context_get_plugins_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Returns the current process model. For more information about this value
see webkit_web_context_set_process_model().
*/
func (self *TraitWebContext) GetProcessModel() (return__ C.WebKitProcessModel) {
	return__ = C.webkit_web_context_get_process_model(self.CPointer)
	return
}

/*
Get the #WebKitSecurityManager of @context.
*/
func (self *TraitWebContext) GetSecurityManager() (return__ *SecurityManager) {
	var __cgo__return__ *C.WebKitSecurityManager
	__cgo__return__ = C.webkit_web_context_get_security_manager(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewSecurityManagerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get whether spell checking feature is currently enabled.
*/
func (self *TraitWebContext) GetSpellCheckingEnabled() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_context_get_spell_checking_enabled(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the the list of spell checking languages associated with
@context, or %NULL if no languages have been previously set.

See webkit_web_context_set_spell_checking_languages() for more
details on the format of the languages in the list.
*/
func (self *TraitWebContext) GetSpellCheckingLanguages() (return__ **C.gchar) {
	return__ = C.webkit_web_context_get_spell_checking_languages(self.CPointer)
	return
}

/*
Get the TLS errors policy of @context
*/
func (self *TraitWebContext) GetTlsErrorsPolicy() (return__ C.WebKitTLSErrorsPolicy) {
	return__ = C.webkit_web_context_get_tls_errors_policy(self.CPointer)
	return
}

/*
Resolve the domain name of the given @hostname in advance, so that if a URI
of @hostname is requested the load will be performed more quickly.
*/
func (self *TraitWebContext) PrefetchDns(hostname string) {
	__cgo__hostname := (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	C.webkit_web_context_prefetch_dns(self.CPointer, __cgo__hostname)
	C.free(unsafe.Pointer(__cgo__hostname))
	return
}

/*
Register @scheme in @context, so that when an URI request with @scheme is made in the
#WebKitWebContext, the #WebKitURISchemeRequestCallback registered will be called with a
#WebKitURISchemeRequest.
It is possible to handle URI scheme requests asynchronously, by calling g_object_ref() on the
#WebKitURISchemeRequest and calling webkit_uri_scheme_request_finish() later
when the data of the request is available or
webkit_uri_scheme_request_finish_error() in case of error.

<informalexample><programlisting>
static void
about_uri_scheme_request_cb (WebKitURISchemeRequest *request,
                             gpointer                user_data)
{
    GInputStream *stream;
    gsize         stream_length;
    const gchar  *path;

    path = webkit_uri_scheme_request_get_path (request);
    if (!g_strcmp0 (path, "plugins")) {
        /<!-- -->* Create a GInputStream with the contents of plugins about page, and set its length to stream_length *<!-- -->/
    } else if (!g_strcmp0 (path, "memory")) {
        /<!-- -->* Create a GInputStream with the contents of memory about page, and set its length to stream_length *<!-- -->/
    } else if (!g_strcmp0 (path, "applications")) {
        /<!-- -->* Create a GInputStream with the contents of applications about page, and set its length to stream_length *<!-- -->/
    } else if (!g_strcmp0 (path, "example")) {
        gchar *contents;

        contents = g_strdup_printf ("&lt;html&gt;&lt;body&gt;&lt;p&gt;Example about page&lt;/p&gt;&lt;/body&gt;&lt;/html&gt;");
        stream_length = strlen (contents);
        stream = g_memory_input_stream_new_from_data (contents, stream_length, g_free);
    } else {
        GError *error;

        error = g_error_new (ABOUT_HANDLER_ERROR, ABOUT_HANDLER_ERROR_INVALID, "Invalid about:%s page.", path);
        webkit_uri_scheme_request_finish_error (request, error);
        g_error_free (error);
        return;
    }
    webkit_uri_scheme_request_finish (request, stream, stream_length, "text/html");
    g_object_unref (stream);
}
</programlisting></informalexample>
*/
func (self *TraitWebContext) RegisterUriScheme(scheme string, callback C.WebKitURISchemeRequestCallback, user_data unsafe.Pointer, user_data_destroy_func C.GDestroyNotify) {
	__cgo__scheme := (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	C.webkit_web_context_register_uri_scheme(self.CPointer, __cgo__scheme, callback, (C.gpointer)(user_data), user_data_destroy_func)
	C.free(unsafe.Pointer(__cgo__scheme))
	return
}

/*
Set an additional directory where WebKit will look for plugins.
*/
func (self *TraitWebContext) SetAdditionalPluginsDirectory(directory string) {
	__cgo__directory := (*C.gchar)(unsafe.Pointer(C.CString(directory)))
	C.webkit_web_context_set_additional_plugins_directory(self.CPointer, __cgo__directory)
	C.free(unsafe.Pointer(__cgo__directory))
	return
}

/*
Specifies a usage model for WebViews, which WebKit will use to
determine its caching behavior. All web views follow the cache
model. This cache model determines the RAM and disk space to use
for caching previously viewed content .

Research indicates that users tend to browse within clusters of
documents that hold resources in common, and to revisit previously
visited documents. WebKit and the frameworks below it include
built-in caches that take advantage of these patterns,
substantially improving document load speed in browsing
situations. The WebKit cache model controls the behaviors of all of
these caches, including various WebCore caches.

Browsers can improve document load speed substantially by
specifying %WEBKIT_CACHE_MODEL_WEB_BROWSER. Applications without a
browsing interface can reduce memory usage substantially by
specifying %WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER. The default value is
%WEBKIT_CACHE_MODEL_WEB_BROWSER.
*/
func (self *TraitWebContext) SetCacheModel(cache_model C.WebKitCacheModel) {
	C.webkit_web_context_set_cache_model(self.CPointer, cache_model)
	return
}

/*
Set the directory where disk cache files will be stored
This method must be called before loading anything in this context, otherwise
it will not have any effect.
*/
func (self *TraitWebContext) SetDiskCacheDirectory(directory string) {
	__cgo__directory := (*C.gchar)(unsafe.Pointer(C.CString(directory)))
	C.webkit_web_context_set_disk_cache_directory(self.CPointer, __cgo__directory)
	C.free(unsafe.Pointer(__cgo__directory))
	return
}

/*
Set the directory path to be used to store the favicons database
for @context on disk. Passing %NULL as @path means using the
default directory for the platform (see g_get_user_data_dir()).

Calling this method also means enabling the favicons database for
its use from the applications, so that's why it's expected to be
called only once. Further calls for the same instance of
#WebKitWebContext won't cause any effect.
*/
func (self *TraitWebContext) SetFaviconDatabaseDirectory(path string) {
	__cgo__path := (*C.gchar)(unsafe.Pointer(C.CString(path)))
	C.webkit_web_context_set_favicon_database_directory(self.CPointer, __cgo__path)
	C.free(unsafe.Pointer(__cgo__path))
	return
}

/*
Set the list of preferred languages, sorted from most desirable
to least desirable. The list will be used to build the "Accept-Language"
header that will be included in the network requests started by
the #WebKitWebContext.
*/
func (self *TraitWebContext) SetPreferredLanguages(languages []string) {
	__header__languages := (*reflect.SliceHeader)(unsafe.Pointer(&languages))
	C.webkit_web_context_set_preferred_languages(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__languages.Data)))
	return
}

/*
Specifies a process model for WebViews, which WebKit will use to
determine how auxiliary processes are handled. The default setting
(%WEBKIT_PROCESS_MODEL_SHARED_SECONDARY_PROCESS) is suitable for most
applications which embed a small amount of WebViews, or are used to
display documents which are considered safe — like local files.

Applications which may potentially use a large amount of WebViews
—for example a multi-tabbed web browser— may want to use
%WEBKIT_PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES, which will use
one process per view most of the time, while still allowing for web
views to share a process when needed (for example when different
views interact with each other). Using this model, when a process
hangs or crashes, only the WebViews using it stop working, while
the rest of the WebViews in the application will still function
normally.

This method **must be called before any other functions**,
as early as possible in your application. Calling it later will make
your application crash.
*/
func (self *TraitWebContext) SetProcessModel(process_model C.WebKitProcessModel) {
	C.webkit_web_context_set_process_model(self.CPointer, process_model)
	return
}

/*
Enable or disable the spell checking feature.
*/
func (self *TraitWebContext) SetSpellCheckingEnabled(enabled bool) {
	__cgo__enabled := C.gboolean(0)
	if enabled {
		__cgo__enabled = C.gboolean(1)
	}
	C.webkit_web_context_set_spell_checking_enabled(self.CPointer, __cgo__enabled)
	return
}

/*
Set the list of spell checking languages to be used for spell
checking.

The locale string typically is in the form lang_COUNTRY, where lang
is an ISO-639 language code, and COUNTRY is an ISO-3166 country code.
For instance, sv_FI for Swedish as written in Finland or pt_BR
for Portuguese as written in Brazil.

You need to call this function with a valid list of languages at
least once in order to properly enable the spell checking feature
in WebKit.
*/
func (self *TraitWebContext) SetSpellCheckingLanguages(languages []string) {
	__header__languages := (*reflect.SliceHeader)(unsafe.Pointer(&languages))
	C.webkit_web_context_set_spell_checking_languages(self.CPointer, (**C.gchar)(unsafe.Pointer(__header__languages.Data)))
	return
}

/*
Set the TLS errors policy of @context as @policy
*/
func (self *TraitWebContext) SetTlsErrorsPolicy(policy C.WebKitTLSErrorsPolicy) {
	C.webkit_web_context_set_tls_errors_policy(self.CPointer, policy)
	return
}

/*
Set the directory where WebKit will look for Web Extensions.
This method must be called before loading anything in this context,
otherwise it will not have any effect. You can connect to
#WebKitWebContext::initialize-web-extensions to call this method
before anything is loaded.
*/
func (self *TraitWebContext) SetWebExtensionsDirectory(directory string) {
	__cgo__directory := (*C.gchar)(unsafe.Pointer(C.CString(directory)))
	C.webkit_web_context_set_web_extensions_directory(self.CPointer, __cgo__directory)
	C.free(unsafe.Pointer(__cgo__directory))
	return
}

/*
Set user data to be passed to Web Extensions on initialization.
The data will be passed to the
#WebKitWebExtensionInitializeWithUserDataFunction.
This method must be called before loading anything in this context,
otherwise it will not have any effect. You can connect to
#WebKitWebContext::initialize-web-extensions to call this method
before anything is loaded.
*/
func (self *TraitWebContext) SetWebExtensionsInitializationUserData(user_data *C.GVariant) {
	C.webkit_web_context_set_web_extensions_initialization_user_data(self.CPointer, user_data)
	return
}

type TraitWebInspector struct{ CPointer *C.WebKitWebInspector }
type IsWebInspector interface {
	GetWebInspectorPointer() *C.WebKitWebInspector
}

func (self *TraitWebInspector) GetWebInspectorPointer() *C.WebKitWebInspector {
	return self.CPointer
}
func NewTraitWebInspector(p unsafe.Pointer) *TraitWebInspector {
	return &TraitWebInspector{(*C.WebKitWebInspector)(p)}
}

/*
Request @inspector to be attached. The signal #WebKitWebInspector::attach
will be emitted. If the inspector is already attached it does nothing.
*/
func (self *TraitWebInspector) Attach() {
	C.webkit_web_inspector_attach(self.CPointer)
	return
}

/*
Request @inspector to be closed.
*/
func (self *TraitWebInspector) Close() {
	C.webkit_web_inspector_close(self.CPointer)
	return
}

/*
Request @inspector to be detached. The signal #WebKitWebInspector::detach
will be emitted. If the inspector is already detached it does nothing.
*/
func (self *TraitWebInspector) Detach() {
	C.webkit_web_inspector_detach(self.CPointer)
	return
}

/*
Get the height that the inspector view should have when
it's attached. If the inspector view is not attached this
returns 0.
*/
func (self *TraitWebInspector) GetAttachedHeight() (return__ uint) {
	var __cgo__return__ C.guint
	__cgo__return__ = C.webkit_web_inspector_get_attached_height(self.CPointer)
	return__ = uint(__cgo__return__)
	return
}

/*
Get the URI that is currently being inspected. This can be %NULL if
nothing has been loaded yet in the inspected view, if the inspector
has been closed or when inspected view was loaded from a HTML string
instead of a URI.
*/
func (self *TraitWebInspector) GetInspectedUri() (return__ string) {
	var __cgo__return__ *C.char
	__cgo__return__ = C.webkit_web_inspector_get_inspected_uri(self.CPointer)
	return__ = C.GoString(__cgo__return__)
	return
}

/*
Get the #WebKitWebViewBase used to display the inspector.
This might be %NULL if the inspector hasn't been loaded yet,
or it has been closed.
*/
func (self *TraitWebInspector) GetWebView() (return__ *WebViewBase) {
	var __cgo__return__ *C.WebKitWebViewBase
	__cgo__return__ = C.webkit_web_inspector_get_web_view(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebViewBaseFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Whether the @inspector view is currently attached to the same window that contains
the inspected view.
*/
func (self *TraitWebInspector) IsAttached() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_inspector_is_attached(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Request @inspector to be shown.
*/
func (self *TraitWebInspector) Show() {
	C.webkit_web_inspector_show(self.CPointer)
	return
}

type TraitWebResource struct{ CPointer *C.WebKitWebResource }
type IsWebResource interface {
	GetWebResourcePointer() *C.WebKitWebResource
}

func (self *TraitWebResource) GetWebResourcePointer() *C.WebKitWebResource {
	return self.CPointer
}
func NewTraitWebResource(p unsafe.Pointer) *TraitWebResource {
	return &TraitWebResource{(*C.WebKitWebResource)(p)}
}

/*
Asynchronously get the raw data for @resource.

When the operation is finished, @callback will be called. You can then call
webkit_web_resource_get_data_finish() to get the result of the operation.
*/
func (self *TraitWebResource) GetData(cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.webkit_web_resource_get_data(self.CPointer, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous operation started with webkit_web_resource_get_data().
*/
func (self *TraitWebResource) GetDataFinish(result *C.GAsyncResult) (length int64, return__ []byte, __err__ error) {
	var __cgo__length C.gsize
	var __cgo_error__ *C.GError
	var __cgo__return__ *C.guchar
	__cgo__return__ = C.webkit_web_resource_get_data_finish(self.CPointer, result, &__cgo__length, &__cgo_error__)
	length = int64(__cgo__length)
	defer func() { return__ = C.GoBytes(unsafe.Pointer(__cgo__return__), C.int(length)) }()
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Retrieves the #WebKitURIResponse of the resource load operation.
This method returns %NULL if called before the response
is received from the server. You can connect to notify::response
signal to be notified when the response is received.
*/
func (self *TraitWebResource) GetResponse() (return__ *URIResponse) {
	var __cgo__return__ *C.WebKitURIResponse
	__cgo__return__ = C.webkit_web_resource_get_response(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewURIResponseFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the current active URI of @resource. The active URI might change during
a load operation:

<orderedlist>
<listitem><para>
  When the resource load starts, the active URI is the requested URI
</para></listitem>
<listitem><para>
  When the initial request is sent to the server, #WebKitWebResource::sent-request
  signal is emitted without a redirected response, the active URI is the URI of
  the request sent to the server.
</para></listitem>
<listitem><para>
  In case of a server redirection, #WebKitWebResource::sent-request signal
  is emitted again with a redirected response, the active URI is the URI the request
  was redirected to.
</para></listitem>
<listitem><para>
  When the response is received from the server, the active URI is the final
  one and it will not change again.
</para></listitem>
</orderedlist>

You can monitor the active URI by connecting to the notify::uri
signal of @resource.
*/
func (self *TraitWebResource) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_web_resource_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

type TraitWebView struct{ CPointer *C.WebKitWebView }
type IsWebView interface {
	GetWebViewPointer() *C.WebKitWebView
}

func (self *TraitWebView) GetWebViewPointer() *C.WebKitWebView {
	return self.CPointer
}
func NewTraitWebView(p unsafe.Pointer) *TraitWebView {
	return &TraitWebView{(*C.WebKitWebView)(p)}
}

/*
Asynchronously execute the given editing command.

When the operation is finished, @callback will be called. You can then call
webkit_web_view_can_execute_editing_command_finish() to get the result of the operation.
*/
func (self *TraitWebView) CanExecuteEditingCommand(command string, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__command := (*C.gchar)(unsafe.Pointer(C.CString(command)))
	C.webkit_web_view_can_execute_editing_command(self.CPointer, __cgo__command, cancellable, callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__command))
	return
}

/*
Finish an asynchronous operation started with webkit_web_view_can_execute_editing_command().
*/
func (self *TraitWebView) CanExecuteEditingCommandFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_view_can_execute_editing_command_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Determines whether @web_view has a previous history item.
*/
func (self *TraitWebView) CanGoBack() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_view_can_go_back(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Determines whether @web_view has a next history item.
*/
func (self *TraitWebView) CanGoForward() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_view_can_go_forward(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Whether or not a MIME type can be displayed in @web_view.
*/
func (self *TraitWebView) CanShowMimeType(mime_type string) (return__ bool) {
	__cgo__mime_type := (*C.gchar)(unsafe.Pointer(C.CString(mime_type)))
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_view_can_show_mime_type(self.CPointer, __cgo__mime_type)
	C.free(unsafe.Pointer(__cgo__mime_type))
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Requests downloading of the specified URI string for @web_view.
*/
func (self *TraitWebView) DownloadUri(uri string) (return__ *Download) {
	__cgo__uri := C.CString(uri)
	var __cgo__return__ *C.WebKitDownload
	__cgo__return__ = C.webkit_web_view_download_uri(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	if __cgo__return__ != nil {
		return__ = NewDownloadFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Request to execute the given @command for @web_view. You can use
webkit_web_view_can_execute_editing_command() to check whether
it's possible to execute the command.
*/
func (self *TraitWebView) ExecuteEditingCommand(command string) {
	__cgo__command := (*C.gchar)(unsafe.Pointer(C.CString(command)))
	C.webkit_web_view_execute_editing_command(self.CPointer, __cgo__command)
	C.free(unsafe.Pointer(__cgo__command))
	return
}

/*
Obtains the #WebKitBackForwardList associated with the given #WebKitWebView. The
#WebKitBackForwardList is owned by the #WebKitWebView.
*/
func (self *TraitWebView) GetBackForwardList() (return__ *BackForwardList) {
	var __cgo__return__ *C.WebKitBackForwardList
	__cgo__return__ = C.webkit_web_view_get_back_forward_list(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewBackForwardListFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the web context of @web_view.
*/
func (self *TraitWebView) GetContext() (return__ *WebContext) {
	var __cgo__return__ *C.WebKitWebContext
	__cgo__return__ = C.webkit_web_view_get_context(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebContextFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Returns the current custom character encoding name of @web_view.
*/
func (self *TraitWebView) GetCustomCharset() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_web_view_get_custom_charset(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the value of the #WebKitWebView:estimated-load-progress property.
You can monitor the estimated progress of a load operation by
connecting to the notify::estimated-load-progress signal of @web_view.
*/
func (self *TraitWebView) GetEstimatedLoadProgress() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.webkit_web_view_get_estimated_load_progress(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Returns favicon currently associated to @web_view, if any. You can
connect to notify::favicon signal of @web_view to be notified when
the favicon is available.
*/
func (self *TraitWebView) GetFavicon() (return__ *C.cairo_surface_t) {
	return__ = C.webkit_web_view_get_favicon(self.CPointer)
	return
}

/*
Gets the #WebKitFindController that will allow the caller to query
the #WebKitWebView for the text to look for.
*/
func (self *TraitWebView) GetFindController() (return__ *FindController) {
	var __cgo__return__ *C.WebKitFindController
	__cgo__return__ = C.webkit_web_view_get_find_controller(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewFindControllerFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Gets the group @web_view belongs to.
*/
func (self *TraitWebView) GetGroup() (return__ *WebViewGroup) {
	var __cgo__return__ *C.WebKitWebViewGroup
	__cgo__return__ = C.webkit_web_view_get_group(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebViewGroupFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the #WebKitWebInspector associated to @web_view
*/
func (self *TraitWebView) GetInspector() (return__ *WebInspector) {
	var __cgo__return__ *C.WebKitWebInspector
	__cgo__return__ = C.webkit_web_view_get_inspector(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebInspectorFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the global JavaScript context used by @web_view to deserialize the
result values of scripts executed with webkit_web_view_run_javascript().
*/
func (self *TraitWebView) GetJavascriptGlobalContext() (return__ C.JSGlobalContextRef) {
	return__ = C.webkit_web_view_get_javascript_global_context(self.CPointer)
	return
}

/*
Return the main resource of @web_view.
*/
func (self *TraitWebView) GetMainResource() (return__ *WebResource) {
	var __cgo__return__ *C.WebKitWebResource
	__cgo__return__ = C.webkit_web_view_get_main_resource(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWebResourceFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the identifier of the #WebKitWebPage corresponding to
the #WebKitWebView
*/
func (self *TraitWebView) GetPageId() (return__ uint64) {
	var __cgo__return__ C.guint64
	__cgo__return__ = C.webkit_web_view_get_page_id(self.CPointer)
	return__ = uint64(__cgo__return__)
	return
}

/*
Gets the #WebKitSettings currently applied to @web_view.
This is a convenient method to get the settings of the
#WebKitWebViewGroup @web_view belongs to.
#WebKitSettings objects are shared by all the #WebKitWebView<!-- -->s
in the same #WebKitWebViewGroup, so modifying
the settings of a #WebKitWebView would affect other
#WebKitWebView<!-- -->s of the same group.
See also webkit_web_view_group_get_settings().
*/
func (self *TraitWebView) GetSettings() (return__ *Settings) {
	var __cgo__return__ *C.WebKitSettings
	__cgo__return__ = C.webkit_web_view_get_settings(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Asynchronously retrieves a snapshot of @web_view for @region.
@options specifies how the snapshot should be rendered.

When the operation is finished, @callback will be called. You must
call webkit_web_view_get_snapshot_finish() to get the result of the
operation.
*/
func (self *TraitWebView) GetSnapshot(region C.WebKitSnapshotRegion, options C.WebKitSnapshotOptions, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.webkit_web_view_get_snapshot(self.CPointer, region, options, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finishes an asynchronous operation started with webkit_web_view_get_snapshot().
*/
func (self *TraitWebView) GetSnapshotFinish(result *C.GAsyncResult) (return__ *C.cairo_surface_t, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_web_view_get_snapshot_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Gets the value of the #WebKitWebView:title property.
You can connect to notify::title signal of @web_view to
be notified when the title has been received.
*/
func (self *TraitWebView) GetTitle() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_web_view_get_title(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Retrieves the #GTlsCertificate associated with the @web_view connection,
and the #GTlsCertificateFlags showing what problems, if any, have been found
with that certificate.
If the connection is not HTTPS, this function returns %FALSE.
This function should be called after a response has been received from the
server, so you can connect to #WebKitWebView::load-changed and call this function
when it's emitted with %WEBKIT_LOAD_COMMITTED event.
*/
func (self *TraitWebView) GetTlsInfo() (certificate *C.GTlsCertificate, errors C.GTlsCertificateFlags, return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_view_get_tls_info(self.CPointer, &certificate, &errors)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Returns the current active URI of @web_view. The active URI might change during
a load operation:

<orderedlist>
<listitem><para>
  When nothing has been loaded yet on @web_view the active URI is %NULL.
</para></listitem>
<listitem><para>
  When a new load operation starts the active URI is the requested URI:
  <itemizedlist>
  <listitem><para>
    If the load operation was started by webkit_web_view_load_uri(),
    the requested URI is the given one.
  </para></listitem>
  <listitem><para>
    If the load operation was started by webkit_web_view_load_html(),
    the requested URI is "about:blank".
  </para></listitem>
  <listitem><para>
    If the load operation was started by webkit_web_view_load_alternate_html(),
    the requested URI is content URI provided.
  </para></listitem>
  <listitem><para>
    If the load operation was started by webkit_web_view_go_back() or
    webkit_web_view_go_forward(), the requested URI is the original URI
    of the previous/next item in the #WebKitBackForwardList of @web_view.
  </para></listitem>
  <listitem><para>
    If the load operation was started by
    webkit_web_view_go_to_back_forward_list_item(), the requested URI
    is the opriginal URI of the given #WebKitBackForwardListItem.
  </para></listitem>
  </itemizedlist>
</para></listitem>
<listitem><para>
  If there is a server redirection during the load operation,
  the active URI is the redirected URI. When the signal
  #WebKitWebView::load-changed is emitted with %WEBKIT_LOAD_REDIRECTED
  event, the active URI is already updated to the redirected URI.
</para></listitem>
<listitem><para>
  When the signal #WebKitWebView::load-changed is emitted
  with %WEBKIT_LOAD_COMMITTED event, the active URI is the final
  one and it will not change unless a new load operation is started
  or a navigation action within the same page is performed.
</para></listitem>
</orderedlist>

You can monitor the active URI by connecting to the notify::uri
signal of @web_view.
*/
func (self *TraitWebView) GetUri() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_web_view_get_uri(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Get the view mode of @web_view.
*/
func (self *TraitWebView) GetViewMode() (return__ C.WebKitViewMode) {
	return__ = C.webkit_web_view_get_view_mode(self.CPointer)
	return
}

/*
Get the #WebKitWindowProperties object containing the properties
that the window containing @web_view should have.
*/
func (self *TraitWebView) GetWindowProperties() (return__ *WindowProperties) {
	var __cgo__return__ *C.WebKitWindowProperties
	__cgo__return__ = C.webkit_web_view_get_window_properties(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewWindowPropertiesFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Get the zoom level of @web_view, i.e. the factor by which the
view contents are scaled with respect to their original size.
*/
func (self *TraitWebView) GetZoomLevel() (return__ float64) {
	var __cgo__return__ C.gdouble
	__cgo__return__ = C.webkit_web_view_get_zoom_level(self.CPointer)
	return__ = float64(__cgo__return__)
	return
}

/*
Loads the previous history item.
You can monitor the load operation by connecting to
#WebKitWebView::load-changed signal.
*/
func (self *TraitWebView) GoBack() {
	C.webkit_web_view_go_back(self.CPointer)
	return
}

/*
Loads the next history item.
You can monitor the load operation by connecting to
#WebKitWebView::load-changed signal.
*/
func (self *TraitWebView) GoForward() {
	C.webkit_web_view_go_forward(self.CPointer)
	return
}

/*
Loads the specific history item @list_item.
You can monitor the load operation by connecting to
#WebKitWebView::load-changed signal.
*/
func (self *TraitWebView) GoToBackForwardListItem(list_item IsBackForwardListItem) {
	C.webkit_web_view_go_to_back_forward_list_item(self.CPointer, list_item.GetBackForwardListItemPointer())
	return
}

/*
Gets the value of the #WebKitWebView:is-loading property.
You can monitor when a #WebKitWebView is loading a page by connecting to
notify::is-loading signal of @web_view. This is useful when you are
interesting in knowing when the view is loding something but not in the
details about the status of the load operation, for example to start a spinner
when the view is loading a page and stop it when it finishes.
*/
func (self *TraitWebView) IsLoading() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_view_is_loading(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Load the given @content string for the URI @content_uri.
This allows clients to display page-loading errors in the #WebKitWebView itself.
When this method is called from #WebKitWebView::load-failed signal to show an
error page, the the back-forward list is maintained appropriately.
For everything else this method works the same way as webkit_web_view_load_html().
*/
func (self *TraitWebView) LoadAlternateHtml(content string, content_uri string, base_uri string) {
	__cgo__content := (*C.gchar)(unsafe.Pointer(C.CString(content)))
	__cgo__content_uri := (*C.gchar)(unsafe.Pointer(C.CString(content_uri)))
	__cgo__base_uri := (*C.gchar)(unsafe.Pointer(C.CString(base_uri)))
	C.webkit_web_view_load_alternate_html(self.CPointer, __cgo__content, __cgo__content_uri, __cgo__base_uri)
	C.free(unsafe.Pointer(__cgo__content))
	C.free(unsafe.Pointer(__cgo__content_uri))
	C.free(unsafe.Pointer(__cgo__base_uri))
	return
}

/*
Load the given @content string with the specified @base_uri.
If @base_uri is not %NULL, relative URLs in the @content will be
resolved against @base_uri and absolute local paths must be children of the @base_uri.
For security reasons absolute local paths that are not children of @base_uri
will cause the web process to terminate.
If you need to include URLs in @content that are local paths in a different
directory than @base_uri you can build a data URI for them. When @base_uri is %NULL,
it defaults to "about:blank". The mime type of the document will be "text/html".
You can monitor the load operation by connecting to #WebKitWebView::load-changed signal.
*/
func (self *TraitWebView) LoadHtml(content string, base_uri string) {
	__cgo__content := (*C.gchar)(unsafe.Pointer(C.CString(content)))
	__cgo__base_uri := (*C.gchar)(unsafe.Pointer(C.CString(base_uri)))
	C.webkit_web_view_load_html(self.CPointer, __cgo__content, __cgo__base_uri)
	C.free(unsafe.Pointer(__cgo__content))
	C.free(unsafe.Pointer(__cgo__base_uri))
	return
}

/*
Load the specified @plain_text string into @web_view. The mime type of
document will be "text/plain". You can monitor the load
operation by connecting to #WebKitWebView::load-changed signal.
*/
func (self *TraitWebView) LoadPlainText(plain_text string) {
	__cgo__plain_text := (*C.gchar)(unsafe.Pointer(C.CString(plain_text)))
	C.webkit_web_view_load_plain_text(self.CPointer, __cgo__plain_text)
	C.free(unsafe.Pointer(__cgo__plain_text))
	return
}

/*
Requests loading of the specified #WebKitURIRequest.
You can monitor the load operation by connecting to
#WebKitWebView::load-changed signal.
*/
func (self *TraitWebView) LoadRequest(request IsURIRequest) {
	C.webkit_web_view_load_request(self.CPointer, request.GetURIRequestPointer())
	return
}

/*
Requests loading of the specified URI string.
You can monitor the load operation by connecting to
#WebKitWebView::load-changed signal.
*/
func (self *TraitWebView) LoadUri(uri string) {
	__cgo__uri := (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	C.webkit_web_view_load_uri(self.CPointer, __cgo__uri)
	C.free(unsafe.Pointer(__cgo__uri))
	return
}

/*
Creates a new #WebKitWebView sharing the same web process with @web_view.
This method doesn't have any effect when %WEBKIT_PROCESS_MODEL_SHARED_SECONDARY_PROCESS
process model is used, because a single web process is shared for all the web views in the
same #WebKitWebContext. When using %WEBKIT_PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES process model,
this method should always be used when creating the #WebKitWebView in the #WebKitWebView::create signal.
You can also use this method to implement other process models based on %WEBKIT_PROCESS_MODEL_MULTIPLE_SECONDARY_PROCESSES,
like for example, sharing the same web process for all the views in the same security domain.
*/
func (self *TraitWebView) NewWithRelatedView() (return__ *C.GtkWidget) {
	return__ = C.webkit_web_view_new_with_related_view(self.CPointer)
	return
}

/*
Reloads the current contents of @web_view.
See also webkit_web_view_reload_bypass_cache().
*/
func (self *TraitWebView) Reload() {
	C.webkit_web_view_reload(self.CPointer)
	return
}

/*
Reloads the current contents of @web_view without
using any cached data.
*/
func (self *TraitWebView) ReloadBypassCache() {
	C.webkit_web_view_reload_bypass_cache(self.CPointer)
	return
}

/*
Asynchronously run @script in the context of the current page in @web_view. If
WebKitWebSettings:enable-javascript is FALSE, this method will do nothing.

When the operation is finished, @callback will be called. You can then call
webkit_web_view_run_javascript_finish() to get the result of the operation.
*/
func (self *TraitWebView) RunJavascript(script string, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__script := (*C.gchar)(unsafe.Pointer(C.CString(script)))
	C.webkit_web_view_run_javascript(self.CPointer, __cgo__script, cancellable, callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__script))
	return
}

/*
Finish an asynchronous operation started with webkit_web_view_run_javascript().

This is an example of using webkit_web_view_run_javascript() with a script returning
a string:

<informalexample><programlisting>
static void
web_view_javascript_finished (GObject      *object,
                              GAsyncResult *result,
                              gpointer      user_data)
{
    WebKitJavascriptResult *js_result;
    JSValueRef              value;
    JSGlobalContextRef      context;
    GError                 *error = NULL;

    js_result = webkit_web_view_run_javascript_finish (WEBKIT_WEB_VIEW (object), result, &error);
    if (!js_result) {
        g_warning ("Error running javascript: %s", error->message);
        g_error_free (error);
        return;
    }

    context = webkit_javascript_result_get_global_context (js_result);
    value = webkit_javascript_result_get_value (js_result);
    if (JSValueIsString (context, value)) {
        JSStringRef js_str_value;
        gchar      *str_value;
        gsize       str_length;

        js_str_value = JSValueToStringCopy (context, value, NULL);
        str_length = JSStringGetMaximumUTF8CStringSize (js_str_value);
        str_value = (gchar *)g_malloc (str_length);
        JSStringGetUTF8CString (js_str_value, str_value, str_length);
        JSStringRelease (js_str_value);
        g_print ("Script result: %s\n", str_value);
        g_free (str_value);
    } else {
        g_warning ("Error running javascript: unexpected return value");
    }
    webkit_javascript_result_unref (js_result);
}

static void
web_view_get_link_url (WebKitWebView *web_view,
                       const gchar   *link_id)
{
    gchar *script;

    script = g_strdup_printf ("window.document.getElementById('%s').href;", link_id);
    webkit_web_view_run_javascript (web_view, script, NULL, web_view_javascript_finished, NULL);
    g_free (script);
}
</programlisting></informalexample>
*/
func (self *TraitWebView) RunJavascriptFinish(result *C.GAsyncResult) (return__ *C.WebKitJavascriptResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_web_view_run_javascript_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously run the script from @resource in the context of the
current page in @web_view.

When the operation is finished, @callback will be called. You can
then call webkit_web_view_run_javascript_from_gresource_finish() to get the result
of the operation.
*/
func (self *TraitWebView) RunJavascriptFromGresource(resource string, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	__cgo__resource := (*C.gchar)(unsafe.Pointer(C.CString(resource)))
	C.webkit_web_view_run_javascript_from_gresource(self.CPointer, __cgo__resource, cancellable, callback, (C.gpointer)(user_data))
	C.free(unsafe.Pointer(__cgo__resource))
	return
}

/*
Finish an asynchronous operation started with webkit_web_view_run_javascript_from_gresource().

Check webkit_web_view_run_javascript_finish() for a usage example.
*/
func (self *TraitWebView) RunJavascriptFromGresourceFinish(result *C.GAsyncResult) (return__ *C.WebKitJavascriptResult, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_web_view_run_javascript_from_gresource_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously save the current web page associated to the
#WebKitWebView into a self-contained format using the mode
specified in @save_mode.

When the operation is finished, @callback will be called. You can
then call webkit_web_view_save_finish() to get the result of the
operation.
*/
func (self *TraitWebView) Save(save_mode C.WebKitSaveMode, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.webkit_web_view_save(self.CPointer, save_mode, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous operation started with webkit_web_view_save().
*/
func (self *TraitWebView) SaveFinish(result *C.GAsyncResult) (return__ *C.GInputStream, __err__ error) {
	var __cgo_error__ *C.GError
	return__ = C.webkit_web_view_save_finish(self.CPointer, result, &__cgo_error__)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Asynchronously save the current web page associated to the
#WebKitWebView into a self-contained format using the mode
specified in @save_mode and writing it to @file.

When the operation is finished, @callback will be called. You can
then call webkit_web_view_save_to_file_finish() to get the result of the
operation.
*/
func (self *TraitWebView) SaveToFile(file *C.GFile, save_mode C.WebKitSaveMode, cancellable *C.GCancellable, callback C.GAsyncReadyCallback, user_data unsafe.Pointer) {
	C.webkit_web_view_save_to_file(self.CPointer, file, save_mode, cancellable, callback, (C.gpointer)(user_data))
	return
}

/*
Finish an asynchronous operation started with webkit_web_view_save_to_file().
*/
func (self *TraitWebView) SaveToFileFinish(result *C.GAsyncResult) (return__ bool, __err__ error) {
	var __cgo_error__ *C.GError
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_web_view_save_to_file_finish(self.CPointer, result, &__cgo_error__)
	return__ = __cgo__return__ == C.gboolean(1)
	if __cgo_error__ != nil {
		__err__ = errors.New(C.GoString((*C.char)(unsafe.Pointer(__cgo_error__.message))))
	}
	return
}

/*
Sets the current custom character encoding override of @web_view. The custom
character encoding will override any text encoding detected via HTTP headers or
META tags. Calling this method will stop any current load operation and reload the
current page. Setting the custom character encoding to %NULL removes the character
encoding override.
*/
func (self *TraitWebView) SetCustomCharset(charset string) {
	__cgo__charset := (*C.gchar)(unsafe.Pointer(C.CString(charset)))
	C.webkit_web_view_set_custom_charset(self.CPointer, __cgo__charset)
	C.free(unsafe.Pointer(__cgo__charset))
	return
}

/*
Sets the #WebKitSettings to be applied to @web_view.
This is a convenient method to set new settings to the
#WebKitWebViewGroup @web_view belongs to.
New settings are applied immediately on all #WebKitWebView<!-- -->s
in the @web_view group.
See also webkit_web_view_group_set_settings().
*/
func (self *TraitWebView) SetSettings(settings IsSettings) {
	C.webkit_web_view_set_settings(self.CPointer, settings.GetSettingsPointer())
	return
}

/*
Set the view mode of @web_view to @view_mode. This method should be called
before loading new contents on @web_view so that the new #WebKitViewMode will
be applied to the new contents.
*/
func (self *TraitWebView) SetViewMode(view_mode C.WebKitViewMode) {
	C.webkit_web_view_set_view_mode(self.CPointer, view_mode)
	return
}

/*
Set the zoom level of @web_view, i.e. the factor by which the
view contents are scaled with respect to their original size.
*/
func (self *TraitWebView) SetZoomLevel(zoom_level float64) {
	C.webkit_web_view_set_zoom_level(self.CPointer, C.gdouble(zoom_level))
	return
}

/*
Stops any ongoing loading operation in @web_view.
This method does nothing if no content is being loaded.
If there is a loading operation in progress, it will be cancelled and
#WebKitWebView::load-failed signal will be emitted with
%WEBKIT_NETWORK_ERROR_CANCELLED error.
*/
func (self *TraitWebView) StopLoading() {
	C.webkit_web_view_stop_loading(self.CPointer)
	return
}

type TraitWebViewBase struct{ CPointer *C.WebKitWebViewBase }
type IsWebViewBase interface {
	GetWebViewBasePointer() *C.WebKitWebViewBase
}

func (self *TraitWebViewBase) GetWebViewBasePointer() *C.WebKitWebViewBase {
	return self.CPointer
}
func NewTraitWebViewBase(p unsafe.Pointer) *TraitWebViewBase {
	return &TraitWebViewBase{(*C.WebKitWebViewBase)(p)}
}

type TraitWebViewGroup struct{ CPointer *C.WebKitWebViewGroup }
type IsWebViewGroup interface {
	GetWebViewGroupPointer() *C.WebKitWebViewGroup
}

func (self *TraitWebViewGroup) GetWebViewGroupPointer() *C.WebKitWebViewGroup {
	return self.CPointer
}
func NewTraitWebViewGroup(p unsafe.Pointer) *TraitWebViewGroup {
	return &TraitWebViewGroup{(*C.WebKitWebViewGroup)(p)}
}

/*
Inject an external style sheet into pages. It is possible to only apply the style sheet
to some URIs by passing non-null values for @whitelist or @blacklist. Passing a %NULL
whitelist implies that all URIs are on the whitelist. The style sheet is applied if a URI matches
the whitelist and not the blacklist. URI patterns must be of the form [protocol]://[host]/[path]
where the host and path components can contain the wildcard character ('*') to represent zero
or more other characters.
*/
func (self *TraitWebViewGroup) AddUserStyleSheet(source string, base_uri string, whitelist []string, blacklist []string, injected_frames C.WebKitInjectedContentFrames) {
	__cgo__source := (*C.gchar)(unsafe.Pointer(C.CString(source)))
	__cgo__base_uri := (*C.gchar)(unsafe.Pointer(C.CString(base_uri)))
	__header__whitelist := (*reflect.SliceHeader)(unsafe.Pointer(&whitelist))
	__header__blacklist := (*reflect.SliceHeader)(unsafe.Pointer(&blacklist))
	C.webkit_web_view_group_add_user_style_sheet(self.CPointer, __cgo__source, __cgo__base_uri, (**C.gchar)(unsafe.Pointer(__header__whitelist.Data)), (**C.gchar)(unsafe.Pointer(__header__blacklist.Data)), injected_frames)
	C.free(unsafe.Pointer(__cgo__source))
	C.free(unsafe.Pointer(__cgo__base_uri))
	return
}

/*
Gets the name that uniquely identifies the #WebKitWebViewGroup.
*/
func (self *TraitWebViewGroup) GetName() (return__ string) {
	var __cgo__return__ *C.gchar
	__cgo__return__ = C.webkit_web_view_group_get_name(self.CPointer)
	return__ = C.GoString((*C.char)(unsafe.Pointer(__cgo__return__)))
	return
}

/*
Gets the #WebKitSettings of the #WebKitWebViewGroup.
*/
func (self *TraitWebViewGroup) GetSettings() (return__ *Settings) {
	var __cgo__return__ *C.WebKitSettings
	__cgo__return__ = C.webkit_web_view_group_get_settings(self.CPointer)
	if __cgo__return__ != nil {
		return__ = NewSettingsFromCPointer(unsafe.Pointer(reflect.ValueOf(__cgo__return__).Pointer()))
	}
	return
}

/*
Remove all style sheets previously injected into this #WebKitWebViewGroup
via webkit_web_view_group_add_user_style_sheet().
*/
func (self *TraitWebViewGroup) RemoveAllUserStyleSheets() {
	C.webkit_web_view_group_remove_all_user_style_sheets(self.CPointer)
	return
}

/*
Sets a new #WebKitSettings for the #WebKitWebViewGroup. The settings will
affect to all the #WebKitWebView<!-- -->s of the group.
#WebKitWebViewGroup<!-- -->s always have a #WebKitSettings so if you just want to
modify a setting you can use webkit_web_view_group_get_settings() and modify the
returned #WebKitSettings instead.
Setting the same #WebKitSettings multiple times doesn't have any effect.
You can monitor the settings of a #WebKitWebViewGroup by connecting to the
notify::settings signal of @group.
*/
func (self *TraitWebViewGroup) SetSettings(settings IsSettings) {
	C.webkit_web_view_group_set_settings(self.CPointer, settings.GetSettingsPointer())
	return
}

type TraitWindowProperties struct{ CPointer *C.WebKitWindowProperties }
type IsWindowProperties interface {
	GetWindowPropertiesPointer() *C.WebKitWindowProperties
}

func (self *TraitWindowProperties) GetWindowPropertiesPointer() *C.WebKitWindowProperties {
	return self.CPointer
}
func NewTraitWindowProperties(p unsafe.Pointer) *TraitWindowProperties {
	return &TraitWindowProperties{(*C.WebKitWindowProperties)(p)}
}

/*
Get whether the window should be shown in fullscreen state or not.
*/
func (self *TraitWindowProperties) GetFullscreen() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_window_properties_get_fullscreen(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get the geometry the window should have on the screen when shown.
*/
func (self *TraitWindowProperties) GetGeometry() (geometry C.GdkRectangle) {
	C.webkit_window_properties_get_geometry(self.CPointer, &geometry)
	return
}

/*
Get whether the window should have the locationbar visible or not.
*/
func (self *TraitWindowProperties) GetLocationbarVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_window_properties_get_locationbar_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get whether the window should have the menubar visible or not.
*/
func (self *TraitWindowProperties) GetMenubarVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_window_properties_get_menubar_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get whether the window should be resizable by the user or not.
*/
func (self *TraitWindowProperties) GetResizable() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_window_properties_get_resizable(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get whether the window should have the scrollbars visible or not.
*/
func (self *TraitWindowProperties) GetScrollbarsVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_window_properties_get_scrollbars_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get whether the window should have the statusbar visible or not.
*/
func (self *TraitWindowProperties) GetStatusbarVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_window_properties_get_statusbar_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}

/*
Get whether the window should have the toolbar visible or not.
*/
func (self *TraitWindowProperties) GetToolbarVisible() (return__ bool) {
	var __cgo__return__ C.gboolean
	__cgo__return__ = C.webkit_window_properties_get_toolbar_visible(self.CPointer)
	return__ = __cgo__return__ == C.gboolean(1)
	return
}
