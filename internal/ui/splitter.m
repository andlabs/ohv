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

	// NSSplitView is fickle; under some circumstances this method will not take effect
	// see http://stackoverflow.com/questions/34574478/how-can-i-set-the-position-of-a-nssplitview-nowadays-setpositionofdivideratind for some details
	// stal in irc.freenode.net/#macdev suggested trying doing this first, and it works, so... thanks to him
	[sv setNeedsLayout:YES];
	[sv layoutSubtreeIfNeeded];

	[sv setPosition:pos ofDividerAtIndex:0];
}
