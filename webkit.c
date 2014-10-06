// 6 october 2014

#include "gtk_unix.h"
#include "_cgo_export.h"

static void webViewLoadResource(WebKitWebView *wv, WebKitWebResource *resource, WebKitURIRequest *request, gpointer data)
{
	printf("%s\n", webkit_uri_request_get_uri(request));
}

//static xxx webViewShowContextMenu(xxx)
//{
//}

WebKitWebView *newWebView(void *gomw)
{
	WebKitWebView *wv;

	wv = WEBKIT_WEB_VIEW(webkit_web_view_new());
	// no cache
	webkit_web_context_set_cache_model(webkit_web_view_get_context(wv), WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER);
	g_signal_connect(wv, "resource-load-started", G_CALLBACK(webViewLoadResource), (gpointer) gomw);
//	g_signal_connect(wv, "context-menu", G_CALLBACK(webViewShowContextMenu), (gpointer) gomw);
	return wv;
}
