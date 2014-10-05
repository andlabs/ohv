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
func (f *File) readContainerPath(index int) (c *ContainerPath) {
	var i uint32

	c = new(ContainerPath)
	f.readSkip()
	c.Index = index
	c.Filename = f.readString()
	c.Path = f.readString()
	if f.versionGreaterEqual("1.0.0.20430") {
		c.Vendor = f.readString()
	}
	f.doSkip()
	if f.err != nil {
		return nil
	}
	return c
}

// TODO isolate core?
// TODO adorn errors?
func (f *File) ReadContainerPaths() (c []*ContainerPath) {
	list := f.readOffsetArray(td.ContainerPathOffset, td.ContainerPathData)
	if f.err != nil {
		return nil
	}
	realOffset := f.realOffset(list, td.ContainerPathData)
	f.seek(realOffset, 0)
	c = make([]*ContainerPath, len(list))
	for i := 0; i < len(list); i++ {
		c[i] := f.readContainerPath(i)
	}
	if f.err != nil {
		return nil
	}
	return c
}
