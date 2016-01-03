// 3 january 2016
#import "uipriv.h"

id newBox(id view1, id view2)
{
	NSStackView *sv;

	sv = [NSStackView stackViewWithViews:[NSArray arrayWithObject:view1]];
	// TODO arrange as in interface builder?
	[sv setOrientation:NSUserInterfaceLayoutOrientationVertical];
	[sv setViews:[NSArray arrayWithObject:view2] inGravity:NSStackViewGravityBottom];
	[sv setSpacing:0];
	return sv;
}

void boxDestroy(id box)
{
	NSStackView *sv = (NSStackView *) box;

	[sv release];
}
