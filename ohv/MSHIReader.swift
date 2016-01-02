// 2 january 2016
import Cocoa

class MSHIReader : NSObject {
	private var bytes: [UInt8]
	private var pos: Int
	private var skip: Int
	private var tailData: TailData
	
	private struct TailData {
		var Version: String = ""
		var FileVersion: String = ""
		var OffsetInterval: UInt32 = 0
		var Product: String = ""
		var ProductVersion: String = ""
		var ProductLocale: String = ""
		var AssetEntryOffset: CLS = CLS()
		var AssetEntryData: CLS = CLS()
		var SearchTerm: ODL = ODL()
		var TOCRootData: CLS = CLS()
		var ContentType: ODL = ODL()
		var ContentFilter: ODL = ODL()
		var Category: ODL = ODL()
		var Parent: ODL = ODL()
		var Content: ODL = ODL()
		var Keyword: ODL = ODL()
		var F1: ODL = ODL()
		var AssetID: ODL = ODL()
		var AssetDataOffset: CLS = CLS()
		var AssetDataData: CLS = CLS()
		var ContainerPathOffset: CLS = CLS()		// TODO rename
		var ContainerPathData: CLS = CLS()
	}
	
	struct ODL {
		var Offset: CLS = CLS()
		var Data: CLS = CLS()
		var Location: CLS = CLS()
	}
	
	struct CLS {
		var Count: UInt32 = 0
		var Length: UInt32 = 0
		var StartPos: UInt32 = 0
	}

	init(_ bytes: [UInt8]) {
		self.bytes = bytes
		self.pos = 0
		self.skip = 0
		self.tailData = TailData()
		super.init()
		self.readTailData()
	}

	private func readu16() -> UInt16 {
		var w: UInt16 = 0
		w |= UInt16(self.bytes[self.pos])
		w |= UInt16(self.bytes[self.pos + 1]) << 8
		self.pos += 2
		self.skip -= 2
		return w
	}
	
	private func readu32() -> UInt32 {
		var l: UInt32 = 0
		l |= UInt32(self.bytes[self.pos])
		l |= UInt32(self.bytes[self.pos + 1]) << 8
		l |= UInt32(self.bytes[self.pos + 2]) << 16
		l |= UInt32(self.bytes[self.pos + 3]) << 24
		self.pos += 4
		self.skip -= 4
		return l
	}
	
	private func readSkip() {
		self.skip = 0
		if self.versionGreaterEqual("1.1.0.0") {
			self.skip = Int(self.read7BitEncodedInt())
		}
	}
	
	private func doSkip() {
		if self.skip > 0 {
			self.pos += self.skip
		}
		self.skip = 0
	}
	
	private func readCLS() -> CLS {
		var cls = CLS()
		cls.Count = self.readu32()
		cls.Length = self.readu32()
		cls.StartPos = self.readu32()
		return cls
	}
	
	private func readODL() -> ODL {
		var odl = ODL()
		odl.Offset = self.readCLS()
		odl.Data = self.readCLS()
		odl.Location = self.readCLS()
		return odl
	}
	
	private func readString() -> String {
		let length = Int(self.read7BitEncodedInt())
		let str = String(bytes: self.bytes[self.pos..<self.pos + length],
			encoding: NSUTF8StringEncoding)!
		self.pos += length
		self.skip -= length
		return str
	}
	
	private func read7BitEncodedInt() -> UInt32 {
		var out: UInt32 = 0
		var shift: UInt32 = 0
		while true {
			let b = self.bytes[self.pos]
			self.pos++
			self.skip--
			out |= UInt32(b & 0x7F) << shift
			shift += 7
			if (b & 0x80) == 0 {
				break
			}
		}
		return out
	}
	
	private func versionGreaterEqual(against: String) -> Bool {
		let version = self.tailData.FileVersion
		let vparts = version.componentsSeparatedByString(",")
		let aparts = against.componentsSeparatedByString(",")
		if vparts.count != aparts.count {
			// TODO add error
			return false
		}
		// TODO error if any fail
		var vints = [Int](count: vparts.count, repeatedValue: 0)
		for (i, p) in vparts.enumerate() {
			vints[i] = Int(p)!
		}
		var aints = [Int](count: aparts.count, repeatedValue: 0)
		for (i, p) in aparts.enumerate() {
			aints[i] = Int(p)!
		}
		for i in 0..<vints.count {
			if vints[i] > aints[i] {
				// versions are greater
				return true
			} else if vints[i] == aints[i] {
				// parts are the same so far...
				continue
			} else {
				// versions are less
				return false
			}
		}
		// versions are the same
		return true
	}
	
	private func readOffsetArray(offset: CLS, _ data: CLS) -> [UInt32] {
		if data.Count == 0 {
			return []
		}
		// TODO HOW WAS THIS SUPPOSED TO WORK
		let n = data.Count// / self.tailData.OffsetInterval
//		if data.Count % self.tailData.OffsetInterval != 0 {
//			n++
//		}
		var list = [UInt32](count: Int(n), repeatedValue: 0)
		self.pos = Int(offset.StartPos)
		let postv1 = self.versionGreaterEqual("1.1.0.0")
		for i in 0..<Int(n) {
			if postv1 {
				self.readSkip()
				list[i] = self.read7BitEncodedInt()
				self.doSkip()
			} else {
				list[i] = self.readu32()
			}
		}
		return list
	}
	
	func realOffset(offsets: [UInt32], _ cls: CLS) -> UInt32 {
		if !self.versionGreaterEqual("1.1.0.0") {
			return offsets[0]
		}
		return offsets[0] + cls.StartPos
	}
	
	private func readTailData() {
		var td = TailData()
		
		self.pos = self.bytes.count - 8
		self.pos = Int(self.readu32())
		td.Version = self.readString()
		td.FileVersion = self.readString()
		self.readSkip()
		td.OffsetInterval = self.readu32()
		td.Product = self.readString()
		td.ProductVersion = self.readString()
		td.ProductLocale = self.readString()
		self.doSkip()
		self.readu32()				// skip length of stats array
		if self.versionGreaterEqual("2.0.0.0") {
			td.AssetEntryOffset = self.readCLS()
			td.AssetEntryData = self.readCLS()
		}
		if self.versionGreaterEqual("1.0.21123.0") {
			td.SearchTerm = self.readODL()
		}
		td.TOCRootData = self.readCLS()
		td.ContentType = self.readODL()
		td.ContentFilter = self.readODL()
		td.Category = self.readODL()
		td.Parent = self.readODL()
		td.Content = self.readODL()
		td.Keyword = self.readODL()
		td.F1 = self.readODL()
		td.AssetID = self.readODL()
		td.AssetDataOffset = self.readCLS()
		td.AssetDataData = self.readCLS()
		td.ContainerPathOffset = self.readCLS()
		td.ContainerPathData = self.readCLS()
		self.tailData = td
	}
	
	struct ContainerPath {
		var Index: Int = 0
		var Filename: String = ""
		var Path: String = ""
		var Vendor: String = ""
	}
	
	private func readContainerPath(index: Int) -> ContainerPath {
		var c = ContainerPath()
		self.readSkip()
		c.Index = index
		c.Filename = self.readString()
		c.Path = self.readString()
		if self.versionGreaterEqual("1.0.0.20430") {
			c.Vendor = self.readString()
		}
		self.doSkip()
		return c
	}
	
	func ReadContainerPaths() -> [ContainerPath] {
		let list = self.readOffsetArray(self.tailData.ContainerPathOffset, self.tailData.ContainerPathData)
		let realOffset = self.realOffset(list, self.tailData.ContainerPathData)
		self.pos = Int(realOffset)
		var c = [ContainerPath](count: list.count, repeatedValue: ContainerPath())
		for i in 0..<c.count {
			c[i] = self.readContainerPath(i)
		}
		return c
	}
	
	struct AssetData {
		var Index: Int = 0
		var ContainerPath: UInt32 = 0
		var EntryName: String = ""
		var CompressedDataOffset: UInt32 = 0
		var CompressedSize: UInt32 = 0
		var UncompressedSize: UInt32 = 0
		var CompressionMethod: UInt16 = 0
		var Locale: String = ""
		var Title: String = ""
		var ID: String = ""
		var Version: UInt32 = 0
		var ParentID: String = ""
		var Order: UInt32 = 0
		var ContentFilter: String = ""
		var ContentType: String = ""
		var Description: String = ""
		var Category: String = ""
		var DisplayVersion: String = ""
		var ParentVersion: UInt32 = 0
		var ParentLocale: String = ""
		var F1KeywordCount: UInt32 = 0
	}
	
	private func readAssetData(index: Int) -> AssetData {
		var a = AssetData()
		self.readSkip()
		a.Index = index
		a.ContainerPath = self.read7BitEncodedInt()
		a.EntryName = self.readString()
		a.CompressedDataOffset = self.read7BitEncodedInt()
		a.CompressedSize = self.read7BitEncodedInt()
		a.UncompressedSize = self.read7BitEncodedInt()
		a.CompressionMethod = self.readu16()
		a.Locale = self.readString()
		a.Title = self.readString()
		a.ID = self.readString()
		a.Version = self.read7BitEncodedInt()
		a.ParentID = self.readString()
		a.Order = self.read7BitEncodedInt()
		a.ContentFilter = self.readString()
		a.ContentType = self.readString()
		a.Description = self.readString()
		if self.versionGreaterEqual("2.0.0.0") {
			a.Category = self.readString()
		}
		if self.versionGreaterEqual("2.0.0.3") {
			a.DisplayVersion = self.readString()
		}
		if self.versionGreaterEqual("2.0.0.5") {
			a.ParentVersion = self.read7BitEncodedInt()
			a.ParentLocale = self.readString()
		}
		if self.versionGreaterEqual("2.0.0.6") {
			a.F1KeywordCount = self.read7BitEncodedInt()
		}
		self.doSkip()
		return a
	}
	
	func ReadAssetDatas() -> [AssetData] {
		let list = self.readOffsetArray(self.tailData.AssetDataOffset, self.tailData.AssetDataData)
		let realOffset = self.realOffset(list, self.tailData.AssetDataData)
		self.pos = Int(realOffset)
		var a = [AssetData](count: list.count, repeatedValue: AssetData())
		for i in 0..<a.count {
			a[i] = self.readAssetData(i)
		}
		return a
	}
}
