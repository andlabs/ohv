// 3 january 2016
#import "uipriv.h"

id newSplitter(id left, id right)
{
	NSSplitView *sv;
	NSView *lv = (NSView *) left;
	NSView *rv = (NSView *) right;

	sv = [[NSSplitView alloc] initWithFrame:NSZeroRect];
	// TODO initialize like Interface Builder
	[sv setVertical:YES];
	[sv setDividerStyle:NSSplitViewDividerStyleThin];

	[sv addSubview:lv];
	[sv addSubview:rv];

	[sv setTranslatesAutoresizingMaskIntoConstraints:NO];

	return sv;
}

void splitterDestroy(id ss)
{
	NSSplitView *sv = (NSSplitView *) ss;

	[sv release];
}

void splitterSetPosition(id ss, intmax_t pos)
{
	NSSplitView *sv = (NSSplitView *) ss;

	[sv setPosition:pos ofDividerAtIndex:0];
}
