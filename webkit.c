// 6 october 2014

#include "gtk_unix.h"
#include "_cgo_export.h"

#define MSXHELP "ms-xhelp"

static void webViewLoadMSXHELP(WebKitURISchemeRequest *request, gpointer data)
{
	mainWindowDoFollowLink((void *) data, (char *) webkit_uri_scheme_request_get_uri(request));
}

WebKitWebView *newWebView(void *gomw)
{
	WebKitWebView *wv;

	wv = WEBKIT_WEB_VIEW(webkit_web_view_new());

	// no cache
	webkit_web_context_set_cache_model(webkit_web_view_get_context(wv), WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER);

	// set up MSHI URIs
	webkit_web_context_register_uri_scheme(webkit_web_view_get_context(wv), MSXHELP,
		webViewLoadMSXHELP, (gpointer) gomw, NULL);

	return wv;
}
