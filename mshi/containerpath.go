// 5 october 2014
package mshi

type ContainerPath struct {
	Index	int
	Filename	string
	Path		string
	Vendor	string
}

func (f *File) readContainerPath(index int) (c *ContainerPath) {
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
func (f *File) ReadContainerPaths() (c []*ContainerPath) {
	list := f.readOffsetArray(f.TailData.ContainerPathOffset, f.TailData.ContainerPathData)
	if f.err != nil {
		return nil
	}
	realOffset := f.realOffset(list, f.TailData.ContainerPathData)
	f.seek(realOffset)
	c = make([]*ContainerPath, len(list))
	for i := 0; i < len(list); i++ {
		c[i] = f.readContainerPath(i)
	}
	if f.err != nil {
		return nil
	}
	return c
}
