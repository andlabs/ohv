// 5 october 2014
package main

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/andlabs/ohv/mshi"
)

type MSHI struct {
	dir			string
	containers	[][]*mshi.ContainerPath
	assets		[][]*mshi.AssetData
	topics		map[string]*MSHITopic
	books		[]Book
	orphans		[]Topic
}

func (m *MSHI) Name() string {
	return "All Microsoft Help Viewer-format documentation"
}

func (m *MSHI) Books() []Book {
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
	err := filepath.Walk(m.dir, func(path string, info os.FileInfo, err error) error {
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
	if err != nil {
		return nil, err
	}

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

	return m, nil
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

func main() {
	m, err := OpenMSHI(os.Args[1])
	if err != nil { panic(err) }
	for _, t := range m.topics {
		println(t.asset.CompressionMethod)
	}
}
