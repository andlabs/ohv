// 2 january 2016
#import "uipriv.h"

id newTree(void)
{
	NSScrollView *sv;
	NSOutlineView *ov;
	NSTableColumn *column;
	NSTextFieldCell *cell;

	cell = [[NSTextFieldCell alloc] initTextCell:@"Text Cell"];
	[cell setTextColor:[NSColor controlTextColor]];
	// TODO bezel style
	[cell setBackgroundColor:[NSColor controlBackgroundColor]];
	[cell setDrawsBackground:NO];
	[cell setEditable:NO];
	[cell setFont:[NSFont systemFontOfSize:[NSFont systemFontSizeForControlSize:NSRegularControlSize]]];
	[cell setScrollable:NO];
	[cell setLineBreakMode:NSLineBreakByTruncatingTail];
	[cell setTruncatesLastVisibleLine:NO];
	[cell setWraps:NO];
	[cell setUsesSingleLineMode:NO];

	column = [[NSTableColumn alloc] initWithIdentifier:nil];
	[column setResizingMask:(NSTableColumnAutoresizingMask | NSTableColumnUserResizingMask)];
	[column setEditable:NO];
	[column setDataCell:cell];

	ov = [[NSOutlineView alloc] initWithFrame:NSZeroRect];
	[ov addTableColumn:column];
	[ov setOutlineTableColumn:column];
	[ov setAutoresizesOutlineColumn:YES];
	[ov setIndentationPerLevel:14];
	[ov setIndentationMarkerFollowsCell:YES];
	[ov setAllowsColumnReordering:NO];
	[ov setAllowsColumnResizing:NO];		// by the user
	[ov setAllowsMultipleSelection:NO];
	[ov setAllowsEmptySelection:YES];		// TODO NO?
	[ov setAllowsColumnSelection:NO];
	[ov setSelectionHighlightStyle:NSTableViewSelectionHighlightStyleSourceList];
	// TODO other style properties?
	[ov setAllowsTypeSelect:YES];
	[ov setHeaderView:nil];				// no header
	[ov setColumnAutoresizingStyle:NSTableViewLastColumnOnlyAutoresizingStyle];

	sv = [[NSScrollView alloc] initWithFrame:NSZeroRect];
	[sv setDrawsBackground:NO];
	[sv setBorderType:NSBezelBorder];
	[sv setHasHorizontalScroller:YES];
	[sv setHasVerticalScroller:YES];
	[sv setAutohidesScrollers:YES];
	[sv setHasHorizontalRuler:NO];
	[sv setHasVerticalRuler:NO];
	[sv setAutomaticallyAdjustsContentInsets:YES];
	[sv setScrollerKnobStyle:NSScrollerKnobStyleDefault];
	// scroller style is documented as being set automatically at runtime based on user preferences; don't touch
	[sv setFindBarPosition:NSScrollViewFindBarPositionAboveContent];
	[sv setUsesPredominantAxisScrolling:NO];
	[sv setHorizontalScrollElasticity:NSScrollElasticityAutomatic];
	[sv setVerticalScrollElasticity:NSScrollElasticityAutomatic];
	[sv setAllowsMagnification:NO];

	[sv setDocumentView:ov];

	[sv setTranslatesAutoresizingMaskIntoConstraints:NO];

	return sv;
}

void treeDestroy(id tree)
{
	NSScrollView *sv = (NSScrollView *) tree;

	[sv release];
}

id treeOutlineView(id tree)
{
	NSScrollView *sv = (NSScrollView *) tree;

	return [sv documentView];
}
