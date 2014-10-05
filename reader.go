// 4 october 2014
package main

import (
	"io"
	"encoding/binary"
)

// TODO adorn error messages?
func (td *TailData) ReadOffsetArray(r io.ReadSeeker, offset CLS, data CLS) ([]uint32, error) {
	var err error

	if data.Count == 0 {
		return nil, nil
	}
	n := data.Count / td.OffsetInterval
	if data.Count % td.OffsetInterval != 0 {
		n++
	}
	list := make([]uint32, n)
	_, err = r.Seek(int64(offset.StartPos), 0)
	if err != nil {
		return nil, err
	}
	postv1 := versionGreaterEqual(td.FileVersion, "1.1.0.0")
	for i := uint32(0); i < n; i++ {
		if postv1 {
			var dskip uint32

			_, err := read7BitEncodedInt(r, &dskip)
			if err != nil {
				return nil, err
			}
			skip := int(dskip)
			n, err := read7BitEncodedInt(r, &list[i])
			if err != nil {
				return nil, err
			}
			skip -= n
			_, err = r.Seek(int64(skip), 1)
			if err != nil {
				return nil, err
			}
		} else {
			err = binary.Read(r, binary.LittleEndian, &list[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return list, nil
}

func (td *TailData) RealOffset(offsets []uint32, cls CLS) int64 {
	if !versionGreaterEqual(td.FileVersion, "1.1.0.0") {
		return int64(offsets[0])
	}
	return int64(offsets[0] + cls.StartPos)
}

func readString(r io.Reader, out *string) (int, error) {
	var length uint32

	n, err := read7BitEncodedInt(r, &length)
	if err != nil {
		return 0, err
	}
	b := make([]byte, length)
	q, err := r.Read(b)
	if err != nil {
		return 0, err
	} else if q != len(b) {		// TODO premature
		return 0, io.EOF
	}
	// don't worry, the string is supposed to be UTF-8
	*out = string(b)
	return n + len(b), nil
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
