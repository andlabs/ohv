// 5 october 2014
package main

import (
	"net/url"
	"os"
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

func LoadMSHILibrary() {
	m, err := OpenMSHI(os.Args[1])
	if err != nil {
		// TODO
		panic(err)
	}
	for _, b := range m.Books() {
		Library = append(Library, b)
	}
}

func LoadLibraries() {
	LoadMSHILibrary()
}
