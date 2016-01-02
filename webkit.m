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

@interface webviewDelegate : NSObject<WebFrameLoadDelegate>
@end

@implementation webviewDelegate

- (void)webView:(WebView *)wv didFailProvisionalLoadWithError:(NSError *)err forFrame:(WebFrame *)frame
{
	NSAlert *alert;

	alert = [NSAlert alertWithError:err];
	[alert runModal];
}

@end

//#define get(x) ((WebView *) [((NSScrollView *) x) documentView])
#define get(x) ((WebView *) (x))

void webviewNavigate(goid webview, char *path)
{
	WebView *wv = get(webview);
	NSString *spath;
	NSURL *url;
	NSURLRequest *req;

	// TODO clear all CSS
	spath = [NSString stringWithUTF8String:path];
	free(path);
	url = [NSURL fileURLWithPath:spath];
	req = [NSURLRequest requestWithURL:url
		cachePolicy:NSURLRequestReloadIgnoringCacheData
		timeoutInterval:INFINITY];			// TODO verify
	[[wv mainFrame] loadRequest:req];
}

goid newWebView(void)
{
	WebView *wv;

	// TODO verify against Interface Builder
	wv = [[WebView alloc] initWithFrame:NSZeroRect frameName:nil groupName:nil];
	[wv setShouldCloseWithWindow:YES];

	// no cache
	// TODO disable the cache or switch to document viewer cache model

	[wv setFrameLoadDelegate:[webviewDelegate new]];

	// set up MSHI URIs
	// TODO set up MSHI URIs

	return wv;
}
