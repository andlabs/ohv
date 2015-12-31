// 31 december 2015
#import <Cocoa/Cocoa.h>
#import "cocoa_darwin.h"

@interface ohvApplication : NSApplication
@end

@implementation ohvApplication

// bypass automatic exit on quit
- (void)terminate:(id)sender
{
}

@end

static ohvApplication *app;

void initUIThread(void)
{
	app = [ohvApplication sharedApplication];
	[app setActivationPolicy:NSApplicationActivationPolicyRegular];
}

void runUIThread(void)
{
	[app run];
}

void stopUIThread(void)
{
	NSEvent *e;

	[app stop:app];
	// stop: won't register until another event has passed; let's synthesize one
	e = [NSEvent otherEventWithType:NSApplicationDefined
		location:NSZeroPoint
		modifierFlags:0
		timestamp:[[NSProcessInfo processInfo] systemUptime]
		windowNumber:0
		context:[NSGraphicsContext currentContext]
		subtype:0
		data1:0
		data2:0];
	[app postEvent:e atStart:NO];
}
