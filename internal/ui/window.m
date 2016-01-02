// 2 january 2016
#import <Cocoa/Cocoa.h>
#import "cocoa_darwin.h"
#import "_cgo_export.h"

@interface windowDelegate : NSObject<NSWindowDelegate>
@end

@implementation windowDelegate

- (BOOL)windowShouldClose:(id)sender
{
	return doWindowClosing(sender) != 0;
}

@end

id newWindow(char *title, int width, int height)
{
	NSWindow *w;

	w = [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, width, height)
		styleMask:(NSTitledWindowMask | NSClosableWindowMask | NSMiniaturizableWindowMask | NSResizableWindowMask)
		backing:NSBackingStoreBuffered
		defer:YES];
	[w setTitle:[NSString stringWithUTF8String:title]];
	free(title);
	[w setReleasedWhenClosed:YES];
	[w setDelegate:[windowDelegate new]];
	return w;
}

void windowDestroy(id ww)
{
	NSWindow *w = (NSWindow *) ww;
	windowDelegate *delegate;

	[w orderOut:nil];
	delegate = (windowDelegate *) [w delegate];
	[w setDelegate:nil];
	[delegate release];
	[w release];
}

void windowMove(id ww, int x, int y)
{
	NSWindow *w = (NSWindow *) ww;

	[w cascadeTopLeftFromPoint:NSMakePoint(x, y)];
}

void windowCenter(id ww)
{
	NSWindow *w = (NSWindow *) ww;

	[w center];
}

void windowShow(id ww)
{
	NSWindow *w = (NSWindow *) ww;

	[w makeKeyAndOrderFront:nil];
}

void windowSetChild(id ww, id cc)
{
	NSWindow *w = (NSWindow *) ww;
	NSView *child = (NSView *) cc;
	NSView *contentView;

	contentView = [w contentView];
	// TODO
}

void windowUnsetChild(id ww, id cc)
{
	NSWindow *w = (NSWindow *) ww;
	NSView *c = (NSView *) cc;
	NSView *contentView;

	contentView = [w contentView];
	[contentView removeAllConstraints];
	[c removeFromSuperview];
}
