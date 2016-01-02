// 31 december 2015
#import "uipriv.h"

@interface ohvApplication : NSApplication
@end

@implementation ohvApplication

// bypass automatic exit on quit
- (void)terminate:(id)sender
{
	NSEvent *e;

	[self stop:self];
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
	[self postEvent:e atStart:NO];
}

@end

@interface appDelegate : NSObject<NSApplicationDelegate>
@end

@implementation appDelegate

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)app
{
	if (shouldQuit())
		// this will call terminate:, which is the same as uiQuit()
		return NSTerminateNow;
	return NSTerminateCancel;
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

const char *initUIThread(void)
{
	mach_error_t err;

	// TODO set to NULL?
	_CFRunLoopSetCurrent(CFRunLoopGetMain());
	// TODO is this part necessary?
	_CFMainPThread = pthread_self();
	// whoops, that isn't enough for WebKit
	// it uses pthread_main_np(); let's override it
	err = mach_override_ptr(pthread_main_np, fakepthreadmain, NULL);
	if (err != err_none) {
		if (err == err_cannot_override)
			return "pthread_main_np() cannot be overrided; we must do so for WebKit — report to andlabs";
		return mach_error_string(err);
	}

	app = [ohvApplication sharedApplication];
	[app setActivationPolicy:NSApplicationActivationPolicyRegular];
	[app setDelegate:[appDelegate new]];

	return NULL;
}

void runUIThread(void)
{
	[app run];
}

void stopUIThread(void)
{
	[app terminate:app];
}

void queueUIThread(uintptr_t n)
{
	// dispatch_get_main_queue() is a serial queue so it will not execute multiple uiQueueMain() functions concurrently
	dispatch_async(dispatch_get_main_queue(), ^{
		doQueued(n);
	});
}
