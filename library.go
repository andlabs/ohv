// 5 october 2014
package main

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/andlabs/ohv/internal/ui"
)

type HelpSource interface {
	Name() string
	Books() []Topic
	Orphans() []Topic
//	SharedAssets() []Asset
	Lookup(url *url.URL) Topic
}

var HelpSources []HelpSource

type Topic interface {
	ui.TreeNode
	Name() string
	Prepare() (*Prepared, error)
	Parent() Topic
	Children() []Topic
	Source() HelpSource
	Less(t Topic) bool
}

type Prepared struct {
	Path			string
	Anchor		string
	CSSPath		string
	CSSBaseDir	string
}

type TopicSorter []Topic

func (t TopicSorter) Len() int { return len(t) }
func (t TopicSorter) Less(i, j int) bool { return t[i].Less(t[j]) }
func (t TopicSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type LibraryModel struct {
	library	[]Topic
	model	*ui.TreeModel
}

var Library LibraryModel

func (l *LibraryModel) Append(topic Topic) {
	l.library = append(l.library, topic)
	if l.model != nil {
		l.model.RowInserted(topic, nil, len(l.library) - 1)
	}
}

func (l *LibraryModel) RootNodes() []ui.TreeNode {
	if len(l.library) == 0 {
		return nil
	}
	c := make([]ui.TreeNode, len(l.library))
	for i, n := range l.library {
		c[i] = n
	}
	return c
}

func (l *LibraryModel) MakeModel() {
	l.model = ui.NewTreeModel(l)
	for i, topic := range l.library {
		l.model.RowInserted(topic, nil, i)
	}
}

func (l *LibraryModel) Model() *ui.TreeModel {
	return l.model
}

func LoadMSHILibrary(dir string) {
	m, err := OpenMSHI(filepath.Join(dir, "mshi"))
	if err != nil {
		// TODO
		panic(err)
	}
	HelpSources = append(HelpSources, m)
	for _, b := range m.Books() {
		Library.Append(b)
	}
}

func LoadAppleLibraries(dir string) {
	dir = filepath.Join(dir, "apple")
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// TODO better check
		if path == dir {
			return nil
		}
		a, err := OpenApple(path)
		if err != nil {
			return err
		}
		HelpSources = append(HelpSources, a)
		for _, b := range a.Books() {
			Library.Append(b)
		}
		return filepath.SkipDir
	})
	if err != nil {
		// TODO
		panic(err)
	}
}

func LoadDevhelpLibraries(dir string) {
	dir = filepath.Join(dir, "devhelp")
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// TODO better check
		if path == dir {
			return nil
		}
		if !info.IsDir() {			// skip .tar.gz files
			return nil
		}
		d, err := OpenDevhelp(path)
		if err != nil {
			return err
		}
		HelpSources = append(HelpSources, d)
		for _, b := range d.Books() {
			Library.Append(b)
		}
		return filepath.SkipDir
	})
	if err != nil {
		// TODO
		panic(err)
	}
}

func LoadLibraries() {
	// all directories /must/ be absolute
	dir, err := filepath.Abs(os.Args[1])
	if err != nil {
		// TODO
		panic(err)
	}
	LoadMSHILibrary(dir)
	// TODO load from system installation path
	LoadAppleLibraries(dir)
	LoadDevhelpLibraries(dir)
}
