// 4 january 2016
package main

import (
	"os"
	"encoding/xml"
	"path/filepath"
	"net/url"
	"strings"
	"errors"

	"github.com/andlabs/ohv/internal/ui"
)

type devhelpSub struct {
	Name	string			`xml:"name,attr"`
	Link		string			`xml:"link,attr"`
	Subs		[]*devhelpSub		`xml:"sub"`
}

type devhelpKeyword struct {
	Name	string		`xml:"name,attr"`
	Link		string		`xml:"link,attr"`
}

type devhelpBook struct {
	Title			string			`xml:"title,attr"`
	Link			string			`xml:"link,attr"`
	Chapters		[]*devhelpSub		`xml:"chapters>sub"`
	Functions		[]*devhelpKeyword	`xml:"functions>keyword"`
}

type Devhelp struct {
	name	string
	dir		string
	root		Topic
}

func (d *Devhelp) Name() string {
	return d.name
}

func (d *Devhelp) Books() []Topic {
	return []Topic{d.root}
}

func (d *Devhelp) Orphans() []Topic {
	return nil
}

func (d *Devhelp) Lookup(url *url.URL) Topic {
	// TODO
	println(url.String())
	return nil
}

func OpenDevhelp(dir string) (*Devhelp, error) {
	d := new(Devhelp)
	d.dir = dir

	files, err := filepath.Glob(filepath.Join(d.dir, "*.devhelp2"))
	if err != nil {
		return nil, err
	}
	if len(files) != 1 {
		return nil, errors.New("zero or multiple possible devhelp files in " + d.dir)
	}

	file, err := os.Open(files[0])
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dd := new(devhelpBook)
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(dd)
	if err != nil {
		return nil, err
	}

	d.root = newDevhelpTopic(dd.Title, dd.Link, dd.Chapters, d.dir, nil, d)
	return d, nil
}

type DevhelpTopic struct {
	name	string
	dir		string
	path		string
	anchor	string
	parent	Topic
	children	[]Topic
	source	HelpSource
}

func newDevhelpTopic(name string, link string, children []*devhelpSub, dir string, parent Topic, source HelpSource) *DevhelpTopic {
	d := new(DevhelpTopic)
	d.name = name
	// devhelp splits filename from anchor with strchr(): https://git.gnome.org/browse/devhelp/tree/src/dh-assistant-view.c?id=b65a383b5f46908ac16ae4f55c7dcce26c891713#n228
	// TODO does the code REQUIRE an anchor? some <sub>s don't have them
	parts := strings.SplitN(link, "#", 2)
	d.path = parts[0]
	if len(parts) == 2 {
		d.anchor = parts[1]
	}
	d.dir = dir
	d.parent = parent
	d.source = source
	for _, child := range children {
		ct := newDevhelpTopic(child.Name, child.Link, child.Subs, d.dir, d, d.source)
		d.children = append(d.children, ct)
	}
	return d
}

func (d *DevhelpTopic) Prepare() (*Prepared, error) {
	return &Prepared{
		Path:			filepath.Join(d.dir, d.path),
		Anchor:		d.anchor,
	}, nil
}

func (d *DevhelpTopic) Name() string {
	return d.name
}

func (d *DevhelpTopic) Parent() Topic {
	return d.parent
}

func (d *DevhelpTopic) Children() []Topic {
	return d.children
}

func (d *DevhelpTopic) Source() HelpSource {
	return d.source
}

func (d *DevhelpTopic) Less(t Topic) bool {
	panic("unreachable; DevhelpTopics are already sorted")
}

func (d *DevhelpTopic) TreeNodeText() string {
	return d.Name()
}

func (d *DevhelpTopic) TreeNodeParent() ui.TreeNode {
	return d.Parent()
}

func (d *DevhelpTopic) TreeNodeChildren() []ui.TreeNode {
	if len(d.children) == 0 {
		return nil
	}
	c := make([]ui.TreeNode, len(d.children))
	for i, n := range d.children {
		c[i] = n
	}
	return c
}
