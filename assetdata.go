// 5 october 2014
package main

import (
	"io"
	"encoding/binary"
)

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
func (td *TailData) ReadAssetData(r io.ReadSeeker, index int) (a *AssetData, err error) {
	var i uint32

	a = new(AssetData)
	skip := 0
	if versionGreaterEqual(td.FileVersion, "1.1.0.0") {
		_, err = read7BitEncodedInt(r, &i)
		if err != nil {
			return nil, err
		}
		skip = int(i)
	}
	a.Index = index
	n, err := read7BitEncodedInt(r, &a.ContainerPath)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &a.EntryName)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = read7BitEncodedInt(r, &a.CompressedDataOffset)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = read7BitEncodedInt(r, &a.CompressedSize)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = read7BitEncodedInt(r, &a.UncompressedSize)
	if err != nil {
		return nil, err
	}
	skip -= n
	err = binary.Read(r, binary.LittleEndian, &a.CompressionMethod)
	if err != nil {
		return nil, err
	}
	skip -= 2
	n, err = readString(r, &a.Locale)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &a.Title)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &a.ID)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = read7BitEncodedInt(r, &a.Version)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &a.ParentID)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = read7BitEncodedInt(r, &a.Order)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &a.ContentFilter)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &a.ContentType)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &a.Description)
	if err != nil {
		return nil, err
	}
	skip -= n
	if versionGreaterEqual(td.FileVersion, "2.0.0.0") {
		n, err = readString(r, &a.Category)
		if err != nil {
			return nil, err
		}
		skip -= n
	}
	if versionGreaterEqual(td.FileVersion, "2.0.0.3") {
		n, err = readString(r, &a.DisplayVersion)
		if err != nil {
			return nil, err
		}
		skip -= n
	}
	if versionGreaterEqual(td.FileVersion, "2.0.0.5") {
		n, err = read7BitEncodedInt(r, &a.ParentVersion)
		if err != nil {
			return nil, err
		}
		skip -= n
		n, err = readString(r, &a.ParentLocale)
		if err != nil {
			return nil, err
		}
		skip -= n
	}
	if versionGreaterEqual(td.FileVersion, "2.0.0.6") {
		n, err = read7BitEncodedInt(r, &a.F1KeywordCount)
		if err != nil {
			return nil, err
		}
		skip -= n
	}
	if skip > 0 {
		_, err := r.Seek(int64(skip), 1)
		if err != nil {
			return nil, err
		}
	}
	return a, nil
}
