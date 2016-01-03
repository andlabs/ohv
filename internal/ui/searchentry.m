// 3 january 2016
#import "uipriv.h"

@interface searchEntryDelegate : NSObject
- (IBAction)onChanged:(id)sender;
@end

@implementation searchEntryDelegate

- (IBAction)onChanged:(id)sender
{
	doSearchEntryChanged(sender);
}

@end

id newSearchEntry(void)
{
	NSSearchField *sf;

	sf = [[NSSearchField alloc] initWithFrame:NSZeroRect];
	// TODO set up like Interface Builder
	[sf setFont:[NSFont systemFontOfSize:[NSFont systemFontSizeForControlSize:NSRegularControlSize]]];
	[sf setSendsWholeSearchString:NO];
	[sf setSendsSearchStringImmediately:NO];

	[sf setTarget:[searchEntryDelegate new]];
	[sf setAction:@selector(onChanged:)];

	[sf setTranslatesAutoresizingMaskIntoConstraints:NO];

	return sf;
}

void searchEntryDestroy(id ss)
{
	NSSearchField *sf = (NSSearchField *) ss;
	searchEntryDelegate *delegate;

	delegate = (searchEntryDelegate *) [sf target];
	[sf setTarget:nil];
	[sf setAction:NULL];
	[delegate release];
	[sf release];
}

char *searchEntryText(id ss)
{
	NSSearchField *sf = (NSSearchField *) ss;
	NSString *s;

	s = [sf stringValue];
	return strdup([s UTF8String]);
}
