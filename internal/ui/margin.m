// 3 january 2016
#import "uipriv.h"

id newMargin(id child)
{
	NSView *view;
	NSDictionary *views;
	NSString *constraint;

	view = [[NSView alloc] initWithFrame:NSZeroRect];

	[view addSubview:child];

	views = NSDictionaryOfVariableBindings(child);

	// TODO if we move this to libui, manually halve the constraints
	constraint = @"H:|-4-[child]-4-|";
	addConstraint(view, constraint, nil, views);
	constraint = @"V:|-4-[child]-4-|";
	addConstraint(view, constraint, nil, views);

	[view setTranslatesAutoresizingMaskIntoConstraints:NO];

	return view;
}

void marginDestroy(id mm)
{
	NSView *view = (NSView *) mm;

	[view release];
}
