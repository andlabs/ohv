// 4 october 2014
package main

import (
	"io"
	"encoding/binary"
	"errors"
	"strings"
	"strconv"
)

// TODO adorn io.ErrUnexpectedEOF?

type File struct {
	TailData	*TailData

	r		io.ReadSeeker
	version	string
	skip		int
	err		error
}

func Open(r io.ReadSeeker) *File {
	f := new(File)
	f.r = r
	f.readTailData()
	return f
}

func Error() error {
	return f.err
}

func (f *File) seek(offset int64, whence int) {
	if f.err != nil {
		return
	}
	_, f.err = f.r.Seek(offset, whence)
}

func (f *File) readu16() uint16 {
	var p [2]byte

	if f.err != nil {
		return 0
	}
	n, err := f.r.Read(p)
	if err != nil {
		f.err = err
		return 0
	} else if n != 2 {
		f.err = io.ErrUnexpectedEOF
		return 0
	}
	f.skip -= 2
	return binary.LittleEndian.Uint16(p)
}

func (f *File) readu32() uint32 {
	var p [4]byte

	if f.err != nil {
		return 0
	}
	n, err := f.r.Read(p)
	if err != nil {
		f.err = err
		return 0
	} else if n != 4 {
		f.err = io.ErrUnexpectedEOF
		return 0
	}
	f.skip -= 4
	return binary.LittleEndian.Uint32(p)
}

func (f *File) readSkip() {
	if f.err != nil {
		return
	}
	f.skip = 0
	if f.versionGreaterEqual("1.1.0.0") {
		f.skip = int(f.read7BitEncodedInt())
	}
}

func (f *File) doSkip() {
	if f.err != nil {
		return
	}
	if f.skip > 0 {
		_, f.err = r.Seek(int64(f.skip), 1)
	}
	f.skip = 0
}

func (f *File) readCLS() (cls CLS) {
	cls.Count = f.readu32()
	cls.Length = f.readu32()
	cls.StartPos = f.readu32()
	return cls
}

func (f *File) readODL() (odl ODL) {
	odl.Offset = f.readCLS()
	odl.Data = f.readCLS()
	odl.Location = f.readCLS()
	return odl
}

func (f *File) readString() string {
	if f.err != nil {
		return ""
	}
	length := f.read7BitEncodedInt()
	if f.err != nil {
		return ""
	}
	p := make([]byte, length)
	n, err := r.Read(p)
	if err != nil {
		f.err = err
		return ""
	} else if n != int(length) {
		f.err = io.ErrUnexpectedEOF
		return ""
	}
	f.skip -= n
	// don't worry, the string is supposed to be UTF-8
	return string(p)
}

func (f *File) read7BitEncodedInt() uint32 {
	var n [1]byte

	if f.err != nil {
		return 0
	}
	out := uint32(0)
	shift := uint32(0)
	for {
		q, err := r.Read(n[:])
		if err != nil {
			f.err = err
			return 0
		} else if q != 1 {
			f.err = io.ErrUnexpectedEOF
			return 0
		}
		f.skip--
		out |= uint32(n[0] & 0x7F) << shift
		shift += 7
		if n[0] & 0x80 == 0 {
			break
		}
	}
	return out
}

func versionGreaterEqual(against string) bool {
	var err error

	if f.err != nil {
		return false
	}
	vparts := strings.Split(f.FileVersion, ".")
	aparts := strings.Split(against, ".")
	if len(vparts) != len(aparts) {
		f.err = errors.New("vparts and aparts have different lengths in File.versionGreaterEqual()")
		return false
	}
	vints := make([]int, len(vparts))
	for i, v := range vparts {
		vints[i], err = strconv.Atoi(v)
		if err != nil {
			// TODO adorn error message
			f.err = err
			return false
		}
	}
	aints := make([]int, len(aparts))
	for i, a := range vparts {
		aints[i], err = strconv.Atoi(a)
		if err != nil {
			// TODO adorn error message
			f.err = err
			return false
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

// TODO adorn error messages?
func (f *File) readOffsetArray(offset CLS, data CLS) []uint32 {
	var err error

	if f.err != nil {
		return nil
	}
	if data.Count == 0 {
		return nil
	}
	n := data.Count / td.OffsetInterval
	if data.Count % td.OffsetInterval != 0 {
		n++
	}
	list := make([]uint32, n)
	f.seek(offset.StartPos, 0)
	postv1 := f.versionGreaterEqual("1.1.0.0")
	for i := uint32(0); i < n; i++ {
		if postv1 {
			f.readSkip()
			list[i] = f.read7BitEncodedInt()
			f.doSkip()
		} else {
			list[i] = f.readu32()
		}
	}
	if f.err != nil {
		return nil
	}
	return list
}

func (f *File) realOffset(offsets []uint32, cls CLS) uint32 {
	if f.err != nil {
		return 0
	}
	if !f.versionGreaterEqual("1.1.0.0") {
		return offsets[0]
	}
	return offsets[0] + cls.StartPos
}
