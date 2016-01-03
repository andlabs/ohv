// 2 january 2016
#import "uipriv.h"

@interface webViewFrameLoadDelegate : NSObject<WebFrameLoadDelegate>
@end

@implementation webViewFrameLoadDelegate

- (void)webView:(WebView *)wv didFailProvisionalLoadWithError:(NSError *)err forFrame:(WebFrame *)frame
{
	doWebViewLoadFailed(wv, err);
}

@end

id newWebView(void)
{
	WebView *wv;

	wv = [[WebView alloc] initWithFrame:NSZeroRect];
	// TODO set up like Interface Builder

	[wv setFrameLoadDelegate:[webViewFrameLoadDelegate new]];

	[wv setTranslatesAutoresizingMaskIntoConstraints:NO];

	return wv;
}

void webViewDestroy(id w)
{
	WebView *wv = (WebView *) w;
	webViewFrameLoadDelegate *fld;

	fld = (webViewFrameLoadDelegate *) [wv frameLoadDelegate];
	[wv setFrameLoadDelegate:nil];
	[fld release];

	[wv release];
}

void webViewNavigate(id w, id u)
{
	WebView *wv = (WebView *) w;
	NSURL *url = (NSURL *) u;
	NSURLRequest *req;

	req = [NSURLRequest requestWithURL:url
		cachePolicy:NSURLRequestReloadIgnoringCacheData
		timeoutInterval:INFINITY];		// TODO verify
	[[wv mainFrame] loadRequest:req];
}
