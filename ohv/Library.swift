// 1 january 2016
import Cocoa

protocol HelpSource : NSObjectProtocol {
	var Name: String { get }
	var Books: [Topic] { get }
	var Orphans: [Topic] { get }
//	var SharedAssets: [Asset] { get }
	func Lookup(url: NSURL) -> Topic
}

protocol Topic : NSObjectProtocol {
	var Name: String { get }
	func Prepare() -> (Prepared, String)
	var Parent: Topic { get }
	var Children: [Topic] { get }
	var Source: HelpSource { get }
	func Less(t: Topic) -> Bool
}

class Prepared : NSObject {
	var Path: String = ""
	var CSSPath: String = ""
	var CSSBaseDir: String = ""
}

var Library: [Topic] = []

func SortBooks(books: [Topic]) -> [Topic] {
	return books.sort({(a, b) -> Bool in
		return a.Less(b)
	})
}

func LoadLibraries() {
	var installedOSXDocs: NSString
	var add: [Topic]
	var err: String

	installedOSXDocs = "~/Library/Developer/Shared/Documentation/DocSets/com.apple.adc.documentation.OSX.docset/"
	installedOSXDocs = installedOSXDocs.stringByExpandingTildeInPath
	if NSFileManager.defaultManager().fileExistsAtPath(installedOSXDocs as String) {
		(add, err) = LoadAppleLibraries(installedOSXDocs as String)
		if err != "" {
			fatalError(err)
		}
		Library.appendContentsOf(add)
	}
}
