// 5 october 2014
package main

import (
	"io"
)

type ContainerPath struct {
	Index	int
	Filename	string
	Path		string
	Vendor	string
}

// TODO adorn errors
func (td *TailData) ReadContainerPath(r io.ReadSeeker, index int) (c *ContainerPath, err error) {
	var i uint32

	c = new(ContainerPath)
	skip := 0
	if versionGreaterEqual(td.FileVersion, "1.1.0.0") {
		_, err = read7BitEncodedInt(r, &i)
		if err != nil {
			return nil, err
		}
		skip = int(i)
	}
	c.Index = index
	n, err := readString(r, &c.Filename)
	if err != nil {
		return nil, err
	}
	skip -= n
	n, err = readString(r, &c.Path)
	if err != nil {
		return nil, err
	}
	skip -= n
	if versionGreaterEqual(td.FileVersion, "1.0.0.20430") {
		n, err = readString(r, &c.Vendor)
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
	return c, nil
}
