// 2 january 2016
#import "uipriv.h"

// TODOs
// - make sure the cache is well and truly disabled

@interface webViewFrameLoadDelegate : NSObject<WebFrameLoadDelegate>
@end

@implementation webViewFrameLoadDelegate

- (void)webView:(WebView *)wv didFailProvisionalLoadWithError:(NSError *)err forFrame:(WebFrame *)frame
{
	doWebViewLoadFailed(wv, err);
}

// TODO what's the difference between the two?
- (void)webView:(WebView *)wv didFailLoadWithError:(NSError *)err forFrame:(WebFrame *)frame
{
	doWebViewLoadFailed(wv, err);
}

@end

@interface webViewPolicyDelegate : NSObject<WebPolicyDelegate>
@end

@implementation webViewPolicyDelegate

- (void)webView:(WebView *)wv decidePolicyForNavigationAction:(NSDictionary *)info request:(NSURLRequest *)req frame:(WebFrame *)frame decisionListener:(id<WebPolicyDecisionListener>)listener
{
	NSNumber *navtype;

	navtype = (NSNumber *) [info objectForKey:WebActionNavigationTypeKey];
	if ([navtype integerValue] == WebNavigationTypeLinkClicked)
		if (doWebViewLinkClicked(wv, [req URL]) == 0)
			[listener ignore];
		// else fall through
	[listener use];
}

@end

@interface webViewResourceLoadDelegate : NSObject<WebResourceLoadDelegate>
@end

@implementation webViewResourceLoadDelegate

- (NSURLRequest *)webView:(WebView *)sender resource:(id)identifier willSendRequest:(NSURLRequest *)request redirectResponse:(NSURLResponse *)redirectResponse fromDataSource:(WebDataSource *)dataSource
{
	// TODO differentiate between the page asset and images
	// TODO implement this
	return request;
}

@end

id newWebView(void)
{
	WebView *wv;

	wv = [[WebView alloc] initWithFrame:NSZeroRect];
	// TODO set up like Interface Builder

	[wv setFrameLoadDelegate:[webViewFrameLoadDelegate new]];
	[wv setPolicyDelegate:[webViewPolicyDelegate new]];
	[wv setResourceLoadDelegate:[webViewResourceLoadDelegate new]];

	[wv setTranslatesAutoresizingMaskIntoConstraints:NO];

	return wv;
}

void webViewDestroy(id w)
{
	WebView *wv = (WebView *) w;
	webViewFrameLoadDelegate *fld;
	webViewPolicyDelegate *pd;
	webViewResourceLoadDelegate *rld;

	rld = (webViewResourceLoadDelegate *) [wv resourceLoadDelegate];
	[wv setResourceLoadDelegate:nil];
	[rld release];

	pd = (webViewPolicyDelegate *) [wv policyDelegate];
	[wv setPolicyDelegate:nil];
	[pd release];

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
