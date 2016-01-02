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
		strongDelegate.appDelegate = self
		do {
			try LoadLibraries()
		} catch let err {
			fatalError("\(err)")
		}
		navtree.setDataSource(self.strongDataSource)
		navtree.setDelegate(self.strongDelegate)
		navtree.reloadData()
	}

	func navigate(topic: Topic) {
		do {
			let prepared = try topic.Prepare()
			self.webview.mainFrame.loadRequest(NSURLRequest(
				URL: NSURL.fileURLWithPath(prepared.Path),
				cachePolicy: NSURLRequestCachePolicy.ReloadIgnoringCacheData,
				timeoutInterval: NSTimeInterval.infinity))
		} catch let err as NSError {
			// TODO customize this
			let alert = NSAlert(error: err)
			alert.beginSheetModalForWindow(self.window, completionHandler: {(response) -> Void in
				// do nothing
			})
		}
	}
	
	func applicationWillTerminate(note: NSNotification) {
		// Insert code here to tear down your application
	}
	
	func applicationShouldTerminateAfterLastWindowClosed(app: NSApplication) -> Bool {
		return true
	}
}
