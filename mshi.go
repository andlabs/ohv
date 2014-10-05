// 5 october 2014
package main

import (
	"os"
	"path/filepath"
	"strings"
	"sort"
	"github.com/andlabs/ohv/mshi"
//	"github.com/davecheney/profile"
)

type MSHI struct {
	dir			string
	containers	[][]*mshi.ContainerPath
	assets		[][]*mshi.AssetData
	topics		map[string]*MSHITopic
	books		[]Topic
	orphans		[]Topic
}

func (m *MSHI) Name() string {
	return "All Microsoft Help Viewer-format documentation"
}

func (m *MSHI) Books() []Topic {
	return m.books
}

func (m *MSHI) Orphans() []Topic {
	return m.orphans
}

// TODO adorn errors
func OpenMSHI(dir string) (*MSHI, error) {
	m := new(MSHI)
	m.dir = dir

	// first read all indices
	err := m.readAllIndices()
	if err != nil {
		return nil, err
	}

	// then produce a map that goes from ID -> topic
	m.collectTopics()

	// now figure out the hierarchy, books, and orphans
	m.buildHierarchy()

	// now sort all children
	m.sortAllChildren()

	return m, nil
}

func (m *MSHI) readAllIndices() error {
	return filepath.Walk(m.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// TODO subdirectories
		if strings.ToLower(filepath.Ext(path)) != ".mshi" {
			return nil
		}
		of, err := os.Open(path)
		if err != nil {
			return err
		}
		f := mshi.Open(of)
		if f.Err() != nil {
			return f.Err()
		}
		c := f.ReadContainerPaths()
		if f.Err() != nil {
			return f.Err()
		}
		a := f.ReadAssetDatas()
		if f.Err() != nil {
			return f.Err()
		}
		of.Close()
		m.containers = append(m.containers, c)
		m.assets = append(m.assets, a)
		return nil
	})
}

func (m *MSHI) collectTopics() {
	// TODO deal with versions sensibly
	m.topics = make(map[string]*MSHITopic)
	for container, aa := range m.assets {
		for _, a := range aa {
			if m.topics[a.ID] != nil && m.topics[a.ID].asset.Version > a.Version {
				continue
			}
			m.topics[a.ID] = &MSHITopic{
				containers:	m.containers[container],
				asset:		a,
			}
		}
	}
}

func (m *MSHI) buildHierarchy() {
	// TODO are parents case-insensitive?
	for _, v := range m.topics {
		if v.asset.ParentID == "-1" {		// is top-level
			m.books = append(m.books, v)
			continue
		}
		parent, ok := m.topics[v.asset.ParentID]
		if ok {			// has parent
			parent.children = append(parent.children, v)
			continue
		}
		m.orphans = append(m.orphans, v)			// parent elsewhere or missing
	}
}

func (m *MSHI) sortAllChildren() {
	for _, t := range m.topics {
		sort.Sort(TopicSorter(t.children))
	}
}

type MSHITopic struct {
	containers	[]*mshi.ContainerPath
	asset			*mshi.AssetData
	children		[]Topic
}

func (m *MSHITopic) Name() string {
	return m.asset.Title
}

func (m *MSHITopic) Prepare() (string, error) {
	// TODO
	return "", nil
}

func (m *MSHITopic) Children() []Topic {
	return m.children
}

func (m *MSHITopic) Less(t Topic) bool {
	tt := t.(*MSHITopic)
	return m.asset.Order < tt.asset.Order
}

func walk(t Topic, level int) {
	println(strings.Repeat(" ", level) + t.Name())
	for _, c := range t.Children() {
		walk(c, level + 1)
	}
}

func xmain() {
//	defer profile.Start(profile.CPUProfile).Stop()
	m, err := OpenMSHI(os.Args[1])
	if err != nil { panic(err) }
	println("books:")
	for _, b := range m.Books() {
		walk(b, 0)
	}
	println("orphans:")
	for _, o := range m.Orphans() {
		walk(o, 1)
	}
}
