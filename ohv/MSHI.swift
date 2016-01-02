// 2 january 2016
import Cocoa

class MSHI : NSObject, HelpSource {
	private var dir: String = ""
	private var containers: [[MSHIReader.ContainerPath]] = []
	private var assets: [[MSHIReader.AssetData]] = []
	private var topics: [String: AppleTopic] = [:]			// maps asset IDs; these are case-insensitive, so are stored here lowercase
	private var books: [Topic] = []
	private var orphans: [Topic] = []
	
	private override init() {
		super.init()
	}
	
	class func Open(dir: String) throws -> MSHI {
		let m = MSHI()
		m.dir = dir
		
		// first read all indices
		try m.readAllIndices()
		
		// then produce a map that goes from ID -> topic
		m.collectTopics()
		
		// now figure out the hierarchy, books, and orphans
		m.buildHierarchy()
		
		// now sort all children
		m.sortAllChildren()
		
		return m
	}
	
	private func addMSHI(path: String) throws {
		let data = try NSData(contentsOfFile: path, options: [])
		var bytes = [UInt8](count: data.length, repeatedValue: 0)
		data.getBytes(&bytes, length: data.length)
		let r = MSHIReader(bytes)
		self.containers.append(r.ReadContainerPaths())
		self.assets.append(r.ReadAssetDatas())
	}
	
	private func readAllIndices() throws {
		let walk = (NSFileManager.defaultManager().enumeratorAtPath(self.dir))!
		for (_, file) in walk.enumerate() {
			let name = file as! String
			if !name.lowercaseString.hasSuffix(".mshi") {
				continue
			}
			try self.addMSHI(name)
		}
	}
	
	var Name: String {
		get {
			return "All Microsoft Help Viewer-format documentation"
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
		// TODO split into a function
		var id = ""
		let queryParts = (NSURLComponents(URL: url, resolvingAgainstBaseURL: true)?.queryItems)!
		for part in queryParts {
			if part.name == "Id" {
				id = part.value!
				break
			}
		}
		id = id.lowercaseString
		return self.topics[id]
	}
}

class MSHITopic : NSObject, Topic {
	
}
