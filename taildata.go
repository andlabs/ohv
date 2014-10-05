// 4 october 2014
package main

type TailData struct {
	Version			string
	FileVersion		string
	OffsetInterval		uint32
	Product			string
	ProductVersion		string
	ProductLocale		string
	AssetEntryOffset	CLS
	AssetEntryData		CLS
	SearchTerm		ODL
	TOCRootData		CLS
	ContentType		ODL
	ContentFilter		ODL
	Category			ODL
	Parent			ODL
	Content			ODL
	Keyword			ODL
	F1				ODL
	AssetID			ODL
	AssetDataOffset	CLS
	AssetDataData		CLS
	ContainerPathOffset	CLS		// TODO rename
	ContainerPathData	CLS
}

type ODL struct {
	Offset	CLS
	Data		CLS
	Location	CLS
}

type CLS struct {
	Count	uint32
	Length	uint32
	StartPos	uint32
}

// TODO adorn error messages?
func (f *File) readTailData() (td *TailData) {
	var tdoff, i uint32

	td = new(TailData)
	f.seek(-8, 2)
	tdoff = f.readu32()
	f.seek(tdoff, 0)
	td.Version = f.readString()
	td.FileVersion = f.readString()
	f.version = td.FileVersion		//  needed by f.readSkip() and some things below
	f.readSkip()
	td.OffsetInterval = f.readu32()
	td.Product = f.readString()
	td.ProductVersion = f.readString()
	td.ProductLocale = f.readString()
	f.doSkip()
	f.readu32()	// skip length of stats array
	if f.versionGreaterEqual("2.0.0.0") {
		td.AssetEntryOffset = f.readCLS()
		td.AssetEntryData = f.readCLS()
	}
	if f.versionGreaterEqual("1.0.21123.0") {
		td.SearchTerm = f.readODL()
	}
	td.TOCRootData = f.readCLS()
	td.ContentType = f.readODL()
	td.ContentFilter = f.readODL()
	td.Category = f.readODL()
	td.Parent = f.readODL()
	td.Content = f.readODL()
	td.Keyword = f.readODL()
	td.F1 = f.readODL()
	td.AssetID = f.readODL()
	td.AssetDataOffset = f.readCLS()
	td.AssetDataData = f.readCLS()
	td.ContainerPathOffset = f.readCLS()
	td.ContainerPathData = f.readCLS()
	if f.err != nil {
		return nil
	}
	return td
}
