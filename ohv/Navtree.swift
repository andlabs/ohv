// 1 january 2016
import Cocoa

class navtreeDataSource : NSObject, NSOutlineViewDataSource {
	func outlineView(outlineView: NSOutlineView, child index: Int, ofItem item: AnyObject?) -> AnyObject {
		if item == nil {
			return Library[index]
		}
		let parent = item as! Topic
		return parent.Children[index]
	}
	
	func outlineView(outlineView: NSOutlineView, isItemExpandable item: AnyObject) -> Bool {
		let topic = item as! Topic
		return topic.Children.count != 0
	}
	
	func outlineView(outlineView: NSOutlineView, numberOfChildrenOfItem item: AnyObject?) -> Int {
		if item == nil {
			return Library.count
		}
		let topic = item as! Topic
		return topic.Children.count
	}
	
	func outlineView(outlineView: NSOutlineView, objectValueForTableColumn tableColumn: NSTableColumn?, byItem item: AnyObject?) -> AnyObject? {
		let result = outlineView.makeViewWithIdentifier("TopicNameView", owner: nil) as! NSTableCellView
		let textfield = result.textField!
		let topic = item as! Topic
		textfield.stringValue = topic.Name
debugPrint(topic.Name)
		return result
	}
}

class navtreeDelegate : NSObject, NSOutlineViewDelegate {
	func outlineViewSelectionDidChange(note: NSNotification) {
		// TODO
	}
}
