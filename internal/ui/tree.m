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
//	[sv setBorderType:NSBezelBorder];
	// TODO the above is for general use
	// this is for our use
	[sv setBorderType:NSNoBorder];
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

@interface treeDataSource : NSObject<NSOutlineViewDataSource>
@end

@implementation treeDataSource

- (id)outlineView:(NSOutlineView *)outlineView child:(NSInteger)index ofItem:(id)item
{
	return treeModelChild(self, (intmax_t) index, item);
}

- (BOOL)outlineView:(NSOutlineView *)outlineView isItemExpandable:(id)item
{
	return treeModelChildCount(self, item) != 0;
}

- (NSInteger)outlineView:(NSOutlineView *)outlineView numberOfChildrenOfItem:(id)item
{
	return (NSInteger) treeModelChildCount(self, item);
}

- (id)outlineView:(NSOutlineView *)outlineView objectValueForTableColumn:(NSTableColumn *)tableColumn byItem:(id)item
{
	char *str;
	NSString *s;

	str = treeModelNodeText(self, item);
	s = [NSString stringWithUTF8String:str];
	free(str);
	return s;
}

@end

id newTreeModel(void)
{
	return [treeDataSource new];
}

void treeModelDestroy(id model)
{
	treeDataSource *tds = (treeDataSource *) model;

	[tds release];
}

void treeSetModel(id tree, id model)
{
	NSScrollView *sv = (NSScrollView *) tree;
	NSOutlineView *ov;
	treeDataSource *tds = (treeDataSource *) model;

	ov = (NSOutlineView *) [sv documentView];
	[ov setDataSource:tds];
	[ov reloadData];
}

void treeInsertRow(id tree, id parent, intmax_t index)
{
	NSScrollView *sv = (NSScrollView *) tree;
	NSOutlineView *ov;
	NSIndexSet *indexSet;

	ov = (NSOutlineView *) [sv documentView];
	indexSet = [NSIndexSet indexSetWithIndex:((NSUInteger) index)];
	[ov beginUpdates];
	[ov insertItemsAtIndexes:indexSet inParent:parent withAnimation:NSTableViewAnimationEffectNone];
	[ov endUpdates];
}

void treeUpdateNode(id tree, id node)
{
	NSScrollView *sv = (NSScrollView *) tree;
	NSOutlineView *ov;

	ov = (NSOutlineView *) [sv documentView];
	[ov reloadItem:node];
}

void treeDeleteRow(id tree, id parent, intmax_t index)
{
	NSScrollView *sv = (NSScrollView *) tree;
	NSOutlineView *ov;
	NSIndexSet *indexSet;

	ov = (NSOutlineView *) [sv documentView];
	indexSet = [NSIndexSet indexSetWithIndex:((NSUInteger) index)];
	[ov beginUpdates];
	[ov removeItemsAtIndexes:indexSet inParent:parent withAnimation:NSTableViewAnimationEffectNone];
	[ov endUpdates];
}

@interface treeNodeObject : NSObject
@end

@implementation treeNodeObject
@end

id newTreeNode(void)
{
	return [treeNodeObject new];
}

void treeNodeDestroy(id node)
{
	treeNodeObject *obj = (treeNodeObject *) node;

	[obj release];
}
