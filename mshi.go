// 5 october 2014
package main

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sort"
	"io"
	"archive/zip"
	"compress/flate"
	"bytes"
	"io/ioutil"
	"encoding/base64"
	"github.com/andlabs/ohv/mshi"
	"code.google.com/p/go.net/html"
)

type MSHI struct {
	dir			string
	containers	[][]*mshi.ContainerPath
	assets		[][]*mshi.AssetData
	topics		map[string]*MSHITopic		// maps asset IDs; these are case-insensitive, so are stored here lowercase
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

func (m *MSHI) Lookup(url *url.URL) Topic {
	id := url.Query().Get("Id")
	id = strings.ToLower(id)
	return m.topics[id]
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
			m.topics[strings.ToLower(a.ID)] = &MSHITopic{
				dir:			m.dir,
				containers:	m.containers[container],
				asset:		a,
				source:		m,
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
		parent, ok := m.topics[strings.ToLower(v.asset.ParentID)]
		if ok {			// has parent
			parent.children = append(parent.children, v)
			v.parent = parent
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
	dir			string
	containers	[]*mshi.ContainerPath
	asset			*mshi.AssetData
	parent		Topic
	children		[]Topic
	source		HelpSource
}

func (m *MSHITopic) Name() string {
	return m.asset.Title
}

// TODO adorn errors
func (m *MSHITopic) Prepare() (string, error) {
	var r io.Reader

	mshc := filepath.Join(m.dir, m.containers[m.asset.ContainerPath].Filename)

	// first get the HTML
	src, err := os.Open(mshc)
	if err != nil {
		return "", err
	}
	_, err = src.Seek(int64(m.asset.CompressedDataOffset), 0)
	if err != nil {
		return "", err
	}
	r = io.LimitReader(src, int64(m.asset.CompressedSize))
	switch m.asset.CompressionMethod {
	case zip.Store:
		// do nothing
	case zip.Deflate:
		r = flate.NewReader(r)
	default:
		return "", zip.ErrAlgorithm
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r)
	if err != nil {
		return "", err
	}

	// unfortunately WebKit2 doesn't provide a way to intercept resource loading
	// WebKit1, however, doesn't provide a method for custom URI schemas
	// so we have to muck with the HTML :|
	src.Close()
	zipfile, err := zip.OpenReader(mshc)
	if err != nil {
		return "", err
	}
	defer zipfile.Close()
	files := make(map[string]*zip.File)
	for _, f := range zipfile.File {
		files[strings.ToLower(f.Name)] = f
	}
	already := make(map[string]string)
	t := html.NewTokenizer(buf)
	out := new(bytes.Buffer)
	for {
		tt := t.Next()
		if tt == html.ErrorToken {
			err := t.Err()
			if err == io.EOF {
				break
			}
			return "", err
		}
		tok := t.Token()
		switch tok.Type {
		case html.StartTagToken, html.SelfClosingTagToken:
			if tok.Data != "img" {
				break
			}
			for i, a := range tok.Attr {
				if a.Key != "src" {
					continue
				}
				filename := strings.ToLower(a.Val)
				if already[filename] == "" {
					f := files[filename]
					r, err := f.Open()
					if err != nil {
						return "", err
					}
					b, err := ioutil.ReadAll(r)
					if err != nil {
						return "", err
					}
					// TODO figure out what the MIME type really is... (WebKit2 requires one to load the image properly...)
					already[filename] = "data:image/png;base64," + base64.StdEncoding.EncodeToString(b)
					r.Close()
				}
				tok.Attr[i].Val = already[filename]
			}
		}
		_, err = out.WriteString(tok.String())
		if err != nil {
			return "", err
		}
	}

	return out.String(), nil
}

func (m *MSHITopic) Parent() Topic {
	return m.parent
}

func (m *MSHITopic) Children() []Topic {
	return m.children
}

func (m *MSHITopic) Source() HelpSource {
	return m.source
}

func (m *MSHITopic) Less(t Topic) bool {
	tt := t.(*MSHITopic)
	return m.asset.Order < tt.asset.Order
}
