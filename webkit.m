// 6 october 2014
#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>
#import "cocoa_darwin.h"
#import "_cgo_export.h"

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
	WebView *wv;

	// TODO verify against Interface Builder
	wv = [[WebView alloc] initWithFrame:NSZeroRect frameName:nil groupName:nil];
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
