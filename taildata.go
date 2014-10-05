// 4 october 2014
package main

import (
	"io"
	"encoding/binary"
	"strings"
	"strconv"
)

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
	ContainerPathOffset	CLS
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

// TODO eliminate
var le = binary.LittleEndian

// TODO adorn error messages?
func ReadTailData(r io.ReadSeeker) (td *TailData, err error) {
	var tdoff, i uint32

	td = new(TailData)
	_, err = r.Seek(-8, 2)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &tdoff)
	if err != nil {
		return nil, err
	}
	_, err = r.Seek(int64(tdoff), 0)
	if err != nil {
		return nil, err
	}
	_, err = readString(r, &td.Version)
	if err != nil {
		return nil, err
	}
	_, err = readString(r, &td.FileVersion)
	if err != nil {
		return nil, err
	}
	skip := 0
	if versionGreaterEqual(td.FileVersion, "1.1.0.0") {
		_, err = read7BitEncodedInt(r, &i)
		if err != nil {
			return nil, err
		}
		skip = int(i)
	}
	n, err := read7BitEncodedInt(r, &td.OffsetInterval)
	if err != nil {
		return nil, err
	}
	skip -= int(n)
	n, err = readString(r, &td.Product)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &td.ProductVersion)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &td.ProductLocale)
	if err != nil {
		return nil, err
	}
	skip -= n
	if skip > 0 {
		_, err := r.Seek(int64(skip), 1)
		if err != nil {
			return nil, err
		}
	}
	// skip length of stats array
	err = binary.Read(r, binary.LittleEndian, &i)
	if err != nil {
		return nil, err
	}
	if versionGreaterEqual(td.FileVersion, "2.0.0.0") {
		err = binary.Read(r, le, &td.AssetEntryOffset)
		if err != nil {
			return nil, err
		}
		err = binary.Read(r, le, &td.AssetEntryData)
		if err != nil {
			return nil, err
		}
	}
	if versionGreaterEqual(td.FileVersion, "1.0.21123.0") {
		err = binary.Read(r, le, &td.SearchTerm)
		if err != nil {
			return nil, err
		}
	}
	err = binary.Read(r, le, &td.TOCRootData)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.ContentType)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.ContentFilter)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.Category)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.Parent)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.Content)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.Keyword)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.F1)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.AssetID)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.AssetDataOffset)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.AssetDataData)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.ContainerPathOffset)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, le, &td.ContainerPathData)
	if err != nil {
		return nil, err
	}
	return td, nil
}

func readString(r io.Reader, out *string) (int, error) {
	var n [1]byte

	q, err := r.Read(n[:])
	if err != nil {
		return 0, err
	} else if q != 1 {			// TODO premature
		return 0, io.EOF
	}
	b := make([]byte, n[0])
	q, err = r.Read(b)
	if err != nil {
		return 0, err
	} else if q != int(n[0]) {		// TODO premature
		return 0, io.EOF
	}
	// don't worry, the string is supposed to be UTF-8
	*out = string(b)
	return int(n[0]) + len(b), nil
}

func read7BitEncodedInt(r io.Reader, out *uint32) (int, error) {
	var n [1]byte

	nBytes := 0
	*out = 0
	shift := uint32(0)
	for {
		q, err := r.Read(n[:])
		if err != nil {
			return 0, err
		} else if q != 1 {			// TODO premature
			return 0, io.EOF
		}
		nBytes++
		*out |= uint32(n[0] & 0x7F) << shift
		shift += 7
		if n[0] & 0x80 == 0 {
			break
		}
	}
	return nBytes, nil
}

// TODO return error?
func versionGreaterEqual(ver string, against string) bool {
	var err error

	vparts := strings.Split(ver, ".")
	aparts := strings.Split(against, ".")
	if len(vparts) != len(aparts) {
		panic("vparts and aparts have different lengths in versionGreaterEqual()")
	}
	vints := make([]int, len(vparts))
	for i, v := range vparts {
		vints[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	aints := make([]int, len(aparts))
	for i, a := range vparts {
		aints[i], err = strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < len(vints); i++ {
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
	// versions match
	return true
}
