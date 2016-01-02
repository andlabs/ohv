// 1 january 2016
import Cocoa
import WebKit

@NSApplicationMain
class AppDelegate : NSObject, NSApplicationDelegate {
	@IBOutlet weak var window: NSWindow!
	@IBOutlet weak var search: NSSearchField!
	@IBOutlet weak var navtree: NSOutlineView!
	@IBOutlet weak var webview: WebView!

	// keep strong to keep alive
	var strongDataSource = navtreeDataSource()
	var strongDelegate = navtreeDelegate()
	
	func applicationDidFinishLaunching(note: NSNotification) {
		do {
			try LoadLibraries()
		} catch let err {
			fatalError("\(err)")
		}
		navtree.setDataSource(self.strongDataSource)
		navtree.setDelegate(self.strongDelegate)
		navtree.reloadData()
	}

	func applicationWillTerminate(note: NSNotification) {
		// Insert code here to tear down your application
	}
	
	func applicationShouldTerminateAfterLastWindowClosed(app: NSApplication) -> Bool {
		return true
	}
}
