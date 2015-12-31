// 31 december 2015
#import <Cocoa/Cocoa.h>
#import "cocoa_darwin.h"
#import "_cgo_export.h"

@interface windowDelegate : NSObject<NSWindowDelegate>
@end

@implementation windowDelegate

- (BOOL)windowShouldClose:(id)sender
{
	return uiOnWindowClosing(sender) != 0;
}

@end

goid newWindow(void)
{
	NSWindow *w;

	w = [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, 1095, 546)
		styleMask:(NSTitledWindowMask | NSClosableWindowMask | NSMiniaturizableWindowMask | NSResizableWindowMask)
		backing:NSBackingStoreBuffered
		defer:YES];
	[w setTitle:@"ohv"];
	[w setReleasedWhenClosed:YES];
	[w setDelegate:[windowDelegate new]];
	// note: original code placed at (100, 100)
	[w center];
	return w;
}

goid newSearchField(void)
{
	NSSearchField *sf;

	sf = [[NSSearchField alloc] initWithFrame:NSZeroRect];
	// TODO configure a la Interface Builder
	[sf setFont:[NSFont systemFontOfSize:[NSFont systemFontSizeForControlSize:NSRegularControlSize]];
	return sf;
}

static void addConstraint(NSView *view, NSString *constraint, NSDictionary *views)
{
	NSArray *constraints;

	constraints = [NSLayoutConstraint constraintsWithVisualFormat:constraint
		options:0
		metrics:nil
		views:views];
	[view addConstraints:constraints];
}

void layoutWindow(goid window, goid search, goid navtree, goid browser)
{
	NSWindow *w = (NSWindow *) window;
	NSSearchField *sf = (NSSearchField *) search;
	NSScrollView *sv = (NSScrollView *) navtree;
	NSScrollView *page = (NSScrollView *) browser;
	NSView *contentView;
	NSDictionary *views;
	NSSplitView *splitView;
	NSView *leftside;

	leftside = [[NSView alloc] initWithFrame:NSZeroRect];
	[leftside addSubview:sf];
	[leftside addSubview:sv];

	splitView = [[NSSplitView alloc] initWithFrame:NSZeroRect];
	// TODO configure like Interface Builder
	[splitView setVertical:NO];
	[splitView addSubview:leftside];
	[splitView addSubview:page];

	contentView = [w contentView];
	[contentView addSubview:splitView];

	views = @{
		@"split":	splitView,
		@"left":	leftside,
		@"sf":	sf,
		@"sv":	sv,
		@"page":	page,		// no need to lay this out, but do run the enumerator below on this too
	};
	[views enumerateKeysAndObjectsUsingBlock:^(id key, id obj, BOOL *stop) {
		NSView *view = (NSView *) obj;

		[view setTranslatesAutoresizingMaskIntoConstraints:NO];
	}];

	addConstraint(contentView, @"H:|[split]|", views);
	addConstraint(contentView, @"V:|[split]|", views);
	addConstraint(leftside, @"H:|[sf]|", views);
	addConstraint(leftside, @"H:|[sf]|", views);
	addConstraint(leftside, @"V:|[sf][sv]|", views);

	// TODO is views autoreleased?
}
