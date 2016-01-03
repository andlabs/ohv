// 2 january 2016
#import "uipriv.h"

id newWebView(void)
{
	WebView *wv;

	wv = [[WebView alloc] initWithFrame:NSZeroRect];

	// TODO set up like Interface Builder

	[wv setTranslatesAutoresizingMaskIntoConstraints:NO];

	return wv;
}

void webViewDestroy(id w)
{
	WebView *wv = (WebView *) w;

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
