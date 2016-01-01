// 31 december 2015
#import <Cocoa/Cocoa.h>
#import <pthread.h>
#import <dlfcn.h>
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
static void *webkit;

void initUIThread(void)
{
	// TODO set to NULL?
	_CFRunLoopSetCurrent(CFRunLoopGetMain());
	// TODO is this part necessary?
	_CFMainPThread = pthread_self();

	// whoops, that isn't enough for WebKit
	// it keeps track of the main thread on its own, so having the loader load it for us will bypass this entirely
	// so unfortunately, we MUST load it manually :(
	webkit = dlopen("/System/Library/Frameworks/WebKit.framework/WebKit", RTLD_NOW);
	if (webkit == NULL) {
		NSLog(@"%s", dlerror());
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
