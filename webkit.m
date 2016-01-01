// 6 october 2014
#import <Cocoa/Cocoa.h>
// we can't use any WebKit classes directly; linker errors will arise
// see main.m for details
#import "cocoa_darwin.h"
#import "_cgo_export.h"

static Class webViewClass = Nil;

#define MSXHELP "ms-xhelp"

/*TODO
static void webViewLoadMSXHELP(WebKitURISchemeRequest *request, gpointer data)
{
	mainWindowDoFollowLink((void *) data, (char *) webkit_uri_scheme_request_get_uri(request));
}
*/

goid newWebView(void)
{
	NSScrollView *sv;
	id wv;

	if (webViewClass == nil)
		webViewClass = NSClassFromString(@"WebView");

	// TODO verify against Interface Builder
	wv = [[webViewClass alloc] initWithFrame:NSZeroRect frameName:nil groupName:nil];
	[wv setShouldCloseWithWindow:YES];

	// no cache
	// TODO disable the cache or switch to document viewer cache model

	// set up MSHI URIs
	// TODO set up MSHI URIs

	sv = [[NSScrollView alloc] initWithFrame:NSZeroRect];
	// TODO
	[sv setHasHorizontalScroller:YES];
	[sv setHasVerticalScroller:YES];

	[sv setDocumentView:wv];

	return sv;
}
