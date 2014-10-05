// 5 october 2014
package mshi

type AssetData struct {
	Index				int
	ContainerPath			uint32
	EntryName			string
	CompressedDataOffset	uint32
	CompressedSize		uint32
	UncompressedSize		uint32
	CompressionMethod	uint16
	Locale				string
	Title					string
	ID					string
	Version				uint32
	ParentID				string
	Order				uint32
	ContentFilter			string
	ContentType			string
	Description			string
	Category				string
	DisplayVersion			string
	ParentVersion			uint32
	ParentLocale			string
	F1KeywordCount		uint32
}

// TODO adorn errors
func (f *File) readAssetData(index int) (a *AssetData) {
	a = new(AssetData)
	f.readSkip()
	a.Index = index
	a.ContainerPath = f.read7BitEncodedInt()
	a.EntryName = f.readString()
	a.CompressedDataOffset = f.read7BitEncodedInt()
	a.CompressedSize = f.read7BitEncodedInt()
	a.UncompressedSize = f.read7BitEncodedInt()
	a.CompressionMethod = f.readu16()
	a.Locale = f.readString()
	a.Title = f.readString()
	a.ID = f.readString()
	a.Version = f.read7BitEncodedInt()
	a.ParentID = f.readString()
	a.Order = f.read7BitEncodedInt()
	a.ContentFilter = f.readString()
	a.ContentType = f.readString()
	a.Description = f.readString()
	if f.versionGreaterEqual("2.0.0.0") {
		a.Category = f.readString()
	}
	if f.versionGreaterEqual("2.0.0.3") {
		a.DisplayVersion = f.readString()
	}
	if f.versionGreaterEqual("2.0.0.5") {
		a.ParentVersion = f.read7BitEncodedInt()
		a.ParentLocale = f.readString()
	}
	if f.versionGreaterEqual("2.0.0.6") {
		a.F1KeywordCount = f.read7BitEncodedInt()
	}
	f.doSkip()
	if f.err != nil {
		return nil
	}
	return a
}

// TODO isolate core?
// TODO adorn errors?
func (f *File) ReadAssetDatas() (a []*AssetData) {
	list := f.readOffsetArray(f.TailData.AssetDataOffset, f.TailData.AssetDataData)
	if f.err != nil {
		return nil
	}
	realOffset := f.realOffset(list, f.TailData.AssetDataData)
	f.seek(realOffset)
	a = make([]*AssetData, len(list))
	for i := 0; i < len(list); i++ {
		a[i] = f.readAssetData(i)
	}
	if f.err != nil {
		return nil
	}
	return a
}
