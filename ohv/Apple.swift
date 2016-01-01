// 1 january 2016
import Cocoa

// DOUBLE BOO
// a) no raw string literals
// b) no hard tabs in strings; have to use \t
let appleQuery =
	"SELECT znode.z_pk, znode.zkname, znodeurl.zpath, znodeurl.zanchor, zorderedsubnode.zparent, zorderedsubnode.zorder" +
	"\n\tFROM znode" +
	"\n\t\tLEFT JOIN znodeurl" +
	"\n\t\t\tON znode.z_pk = znodeurl.znode" +
	"\n\t\tLEFT JOIN zorderedsubnode" +
	"\n\t\t\tON znode.z_pk = zorderedsubnode.znode" +
	"\n\t\t\tAND znode.zprimaryparent = zorderedsubnode.zparent;"

struct appleNode {
	var Z_PK: Int64
	var ZKNAME: String
	var ZPATH: String
	var ZANCHOR: String
	var ZPARENT: Int64
	var ZORDER: Int64	// 1-based!
}

class Apple : NSObject, HelpSource {
	private var dir: String = ""
	private var dircr: String = ""
	private var nodes: [appleNode] = []
	private var topics: [Int64: AppleTopic] = [:]
	private var books: [Topic] = []
	private var orphans: [Topic] = []

	private override init() {
		super.init()
	}

	class func Open(dir: String) -> (Apple, String) {
		let a = Apple()
		a.dir = dir
		a.dircr = NSString.pathWithComponents([a.dir, "Contents", "Resources"])
		
		let dbname = NSString.pathWithComponents([a.dircr, "docSet.dsidx"])
	}
	
	var Name: String {
		get {
			return "Apple"
		}
	}
	
	var Books: [Topic] {
		get {
			return self.books
		}
	}
	
	var Orphans: [Topic] {
		get {
			return self.orphans
		}
	}
	
//	var SharedAssets: [Asset] { get }

	func Lookup(url: NSURL) -> Topic {
		// TODO
	}
}

class AppleTopic : NSObject, Topic {
	
}
