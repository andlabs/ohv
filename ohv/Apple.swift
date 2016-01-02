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
	var Z_PK: Int64 = 0
	var ZKNAME: String = ""
	var ZPATH: String = ""
	var ZANCHOR: String = ""
	var HasParent: Bool = false
	var ZPARENT: Int64 = 0
	var ZORDER: Int64 = 0	// 1-based!
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

	class func Open(dir: String) throws -> Apple {
		let a = Apple()
		a.dir = dir
		a.dircr = NSString.pathWithComponents([a.dir, "Contents", "Resources"])
		
		let dbname = NSString.pathWithComponents([a.dircr, "docSet.dsidx"])
		let db = FMDatabase(path: dbname)
		if !db.openWithFlags(SQLITE_OPEN_READONLY) {
			throw db.lastError()
		}
		defer { db.close() }
		
		let results = db.executeQuery(appleQuery, withArgumentsInArray: nil)
		defer { results.close() }
		while results.next() {
			var node = appleNode()
			node.Z_PK = results.longLongIntForColumnIndex(0)
			node.ZKNAME = results.stringForColumnIndex(1)
			if !results.columnIndexIsNull(2) {
				// TODO what happens if the column IS null?
				node.ZPATH = results.stringForColumnIndex(2)
			}
			if !results.columnIndexIsNull(3) {
				node.ZANCHOR = results.stringForColumnIndex(3)
			}
			if !results.columnIndexIsNull(4) {
				node.HasParent = true
				node.ZPARENT = results.longLongIntForColumnIndex(4)
			}
			node.ZORDER = results.longLongIntForColumnIndex(5)
			a.nodes.append(node)
		}
		
		// now create the topics
		for node in a.nodes {
			var topic: AppleTopic

			// TODO make this into a function and note that .keys is inefficient
			if a.topics[node.Z_PK] != nil {
				throw LibraryError.HasDuplicates
			}
			topic = AppleTopic(a)
			topic.dir = a.dircr				// note: not a.dir
			topic.Name = node.ZKNAME
			topic.hasParent = node.HasParent
			topic.pptr = node.ZPARENT
			topic.path = node.ZPATH
			topic.anchor = node.ZANCHOR
			topic.order = node.ZORDER - 1	// remember that it's 1-based
			a.topics[node.Z_PK] = topic
		}

		// now hook up parents and children and sort into books and orphans
		// TODO make this use a method of AppleTopic and make the above all in the constructor
		// TODO and make all the variables of AppleTopic properties
		for (_, topic) in a.topics {
			if !topic.hasParent {
				a.books.append(topic)
				continue
			}
			if a.topics[topic.pptr] != nil {
				let parent = a.topics[topic.pptr]!
				parent.Children.append(topic)
				topic.Parent = parent
				continue
			}
			// otherwise this is an orphan
			a.orphans.append(topic)
		}
		
		// finally get everything into sorted order
		for (_, topic) in a.topics {
			topic.Children = SortBooks(topic.Children)
		}
		
		return a
	}
	
	var Name: String {
		get {
			// TODO
//			return "Apple documentation collection " + filepath.Base(a.dir)
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

	func Lookup(url: NSURL) -> Topic? {
		// TODO
		return nil
	}
}

class AppleTopic : NSObject, Topic {
	var dir: String = ""
	var hasParent: Bool = false
	var pptr: Int64 = 0
	var path: String = ""
	var anchor: String = ""
	var order: Int64 = 0
	
	var Name: String = ""
	
	func Prepare() throws -> Prepared {
		let p = Prepared()
		let cssdir = NSString.pathWithComponents([self.dir, "Contents", "Resources", "1014", "CSS"]) as String
		p.Path = NSString.pathWithComponents([self.dir, "Documents", self.path]) as String
		p.CSSPath = NSString.pathWithComponents([cssdir, "xcode5.css"]) as String
		p.CSSBaseDir = cssdir
		return p
	}
	
	var Parent: Topic? = nil
	
	var Children: [Topic] = []
	
	var Source: HelpSource

	func Less(b: Topic) -> Bool {
		return self.order < (b as! AppleTopic).order
	}
	
	init(_ helpSource: HelpSource) {
		self.Source = helpSource
		super.init()
	}
}
