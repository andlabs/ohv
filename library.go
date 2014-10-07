// 5 october 2014
package main

import (
	"net/url"
	"os"
	"path/filepath"
)

type HelpSource interface {
	Name() string
	Books() []Topic
	Orphans() []Topic
//	SharedAssets() []Asset
	Lookup(url *url.URL) Topic
}

type Topic interface {
	Name() string
	Prepare() (string, error)
	Parent() Topic
	Children() []Topic
	Source() HelpSource
	Less(t Topic) bool
}

type TopicSorter []Topic

func (t TopicSorter) Len() int { return len(t) }
func (t TopicSorter) Less(i, j int) bool { return t[i].Less(t[j]) }
func (t TopicSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

var Library []Topic

func LoadMSHILibrary(dir string) {
	m, err := OpenMSHI(filepath.Join(dir, "mshi"))
	if err != nil {
		// TODO
		panic(err)
	}
	for _, b := range m.Books() {
		Library = append(Library, b)
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
		for _, b := range a.Books() {
			Library = append(Library, b)
		}
		return filepath.SkipDir
	})
	if err != nil {
		// TODO
		panic(err)
	}
}

func LoadLibraries() {
	LoadMSHILibrary(os.Args[1])
	LoadAppleLibraries(os.Args[1])
}
