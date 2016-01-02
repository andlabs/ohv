// 31 december 2015
#import <Cocoa/Cocoa.h>
#import <pthread.h>
#import "mach_override.h"
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

// bypass main thread restrictions
extern void _CFRunLoopSetCurrent(CFRunLoopRef);
extern pthread_t _CFMainPThread;
static int fakepthreadmain(void)
{
	return 1;
}

void initUIThread(void)
{
	// TODO set to NULL?
	_CFRunLoopSetCurrent(CFRunLoopGetMain());
	// TODO is this part necessary?
	_CFMainPThread = pthread_self();

	// whoops, that isn't enough for WebKit
	// it uses pthread_main_np(); let's override it
	if (mach_override_ptr(pthread_main_np, fakepthreadmain, NULL) != err_none) {
		NSLog(@"it failed");
		abort();
	}

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
