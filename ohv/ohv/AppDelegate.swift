// 1 january 2016
import Cocoa
import WebKit

@NSApplicationMain
class AppDelegate : NSObject, NSApplicationDelegate {
	@IBOutlet weak var window: NSWindow!
	@IBOutlet weak var search: NSSearchField!
	@IBOutlet weak var navtree: NSOutlineView!
	@IBOutlet weak var webview: WebView!

	func applicationDidFinishLaunching(note: NSNotification) {
		// Insert code here to initialize your application
	}

	func applicationWillTerminate(note: NSNotification) {
		// Insert code here to tear down your application
	}
	
	func applicationShouldTerminateAfterLastWindowClosed(app: NSApplication) -> Bool {
		return true
	}
}
