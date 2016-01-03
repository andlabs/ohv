// 3 january 2016
#import "uipriv.h"

// TODO this uses NSVisualEffectView because wrapping it around the NSStackView didn't work
// TODO does NSVisualEffectView have its own border?

id newMargin(id child)
{
	NSVisualEffectView *view;
	NSDictionary *views;
	NSString *constraint;

	view = [[NSVisualEffectView alloc] initWithFrame:NSZeroRect];
	// TODO interface builder?
	[view setMaterial:NSVisualEffectMaterialSidebar];
	[view setBlendingMode:NSVisualEffectBlendingModeBehindWindow];
	// TODO state?

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
