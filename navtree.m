// 31 december 2015
#import <Cocoa/Cocoa.h>
#import "cocoa_darwin.h"
#import "_cgo_export.h"

indexArray newIndexArray(void)
{
	return (indexArray) [NSMutableArray new];
}

indexArray indexArrayCloneAppend(indexArray ia, intmax_t index)
{
	NSMutableArray *arr1 = (NSMutableArray *) ia;
	NSMutableArray *arr2;

	arr2 = [NSMutableArray new];
	if (arr1 != nil)
		[arr2 addObjectsFromArray:arr1];
	[arr2 addObject:[NSNumber numberWithInteger:index]];
	return (indexArray) arr2;
}

intmax_t indexArrayLen(indexArray ia)
{
	NSMutableArray *arr = (NSMutableArray *) ia;

	if (arr == nil)
		return 0;
	return (intmax_t) [arr count];
}

intmax_t indexArrayIndex(indexArray ia, intmax_t index)
{
	NSMutableArray *arr = (NSMutableArray *) ia;
	NSNumber *n;

	n = (NSNumber *) [arr objectAtIndex:index];
	return (intmax_t) [n integerValue];
}

void indexArrayPrepend(indexArray ia, intmax_t index)
{
	NSMutableArray *arr = (NSMutableArray *) ia;

	[arr insertObject:[NSNumber numberWithInteger:index] atIndex:0];
}

@interface navtreeDataSource : NSObject<NSOutlineViewDataSource>
@end

@implementation navtreeDataSource

- (id)outlineView:(NSOutlineView *)outlineView child:(NSInteger)index ofItem:(id)item
{
	return indexArrayCloneAppend((indexArray) item, (intmax_t) index);
}

- (BOOL)outlineView:(NSOutlineView *)outlineView isItemExpandable:(id)item
{
	return navtreeChildCount((indexArray) item) != 0;
}

- (NSInteger)outlineView:(NSOutlineView *)outlineView numberOfChildrenOfItem:(id)item
{
	if (item == nil)
		return (NSInteger) navtreeBookCount();
	return (NSInteger) navtreeChildCount((indexArray) item);
}

- (id)outlineView:(NSOutlineView *)outlineView objectValueForTableColumn:(NSTableColumn *)tableColumn byItem:(id)item
{
	char *text;
	NSString *str;

	text = navtreeItemName((indexArray) item);
	str = [NSString stringWithUTF8String:text];
	free(text);
	return str;
}

- (void)outlineView:(NSOutlineView *)outlineView setObjectValue:(id)object forTableColumn:(NSTableColumn *)tableColumn byItem:(id)item
{
	NSLog(@"TODO pop up warning");
	abort();
}

@end

@interface navtreeDelegate : NSObject<NSOutlineViewDelegate>
@end

@implementation navtreeDelegate

- (void)outlineViewSelectionDidChange:(NSNotification *)note
{
	// TODO
}

@end

goid newNavtree(void)
{
	NSScrollView *sv;
	NSOutlineView *ov;
	NSTableColumn *column;
	NSTextFieldCell *cell;

	// TODO verify ALL OF THIS against Interface Builder

	ov = [[NSOutlineView alloc] initWithFrame:NSZeroRect];

	// TODO disable persistance?

	[ov setAllowsColumnReordering:NO];
	[ov setAllowsColumnResizing:NO];
	[ov setAllowsMultipleSelection:NO];
	[ov setAllowsEmptySelection:NO];
	[ov setAllowsColumnSelection:NO];

	// TODO model what Xcode's viewer does
	[ov setSelectionHighlightStyle:NSTableViewSelectionHighlightStyleSourceList];

	// TODO set this property
	cell = [NSTextFieldCell new];
	[cell setFont:[NSFont systemFontOfSize:[NSFont systemFontSizeForControlSize:NSRegularControlSize]]];

	column = [[NSTableColumn alloc] initWithIdentifier:@"title"];
	[column setResizingMask:NSTableColumnAutoresizingMask];
	[column setEditable:NO];
	[column setDataCell:cell];
	[ov addTableColumn:column];
	[ov setOutlineTableColumn:column];

	// TODO type select?

	[ov setHeaderView:nil];		// drop the header

	// TODO autoresizing style?

	// don't set up the data source and delegate just yet; wait until the Go side is ready to process events
	// we actually do the setup in navtreeBegin()

	sv = [[NSScrollView alloc] initWithFrame:NSZeroRect];
	// TODO
	[sv setHasHorizontalScroller:YES];
	[sv setHasVerticalScroller:YES];

	[sv setDocumentView:ov];

	return sv;
}

void navtreeBegin(goid navtree)
{
	NSScrollView *sv = (NSScrollView *) navtree;
	NSOutlineView *ov;

	ov = [sv documentView];

	[ov setDataSource:[navtreeDataSource new]];
	[ov setDelegate:[navtreeDelegate new]];
}

void navtreeSelect(goid navtree, indexArray ia)
{
	NSScrollView *sv = (NSScrollView *) navtree;
	NSOutlineView *ov;
	id item = (id) ia;
	NSInteger row;

	ov = [sv documentView];

	[ov expandItem:item];
	row = [ov rowForItem:item];
	[ov scrollRowToVisible:row];
	[ov selectRowIndexes:[NSIndexSet indexSetWithIndex:row] byExtendingSelection:NO];
	// TODO invoke delegate manually?
}
