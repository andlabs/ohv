// 1 january 2016
import Cocoa

protocol HelpSource : NSObjectProtocol {
	var Name: String { get }
	var Books: [Topic] { get }
	var Orphans: [Topic] { get }
//	var SharedAssets: [Asset] { get }
	func Lookup(url: NSURL) -> Topic?
}

protocol Topic : NSObjectProtocol {
	var Name: String { get }
	func Prepare() throws -> Prepared
	var Parent: Topic? { get }
	var Children: [Topic] { get }
	var Source: HelpSource { get }
	func Less(t: Topic) -> Bool
}

class Prepared : NSObject {
	var Path: String = ""
	var CSSPath: String = ""
	var CSSBaseDir: String = ""
}

enum LibraryError : ErrorType {
	case HasDuplicates
}

var Library: [Topic] = []

func SortBooks(books: [Topic]) -> [Topic] {
	return books.sort({(a, b) -> Bool in
		return a.Less(b)
	})
}

func LoadLibraries() throws {
	var installedOSXDocs: String

	installedOSXDocs = ("~/Library/Developer/Shared/Documentation/DocSets/com.apple.adc.documentation.OSX.docset/" as NSString).stringByExpandingTildeInPath
	if NSFileManager.defaultManager().fileExistsAtPath(installedOSXDocs) {
		Library.appendContentsOf((try Apple.Open(installedOSXDocs)).Books)
	}
}
