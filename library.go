// 5 october 2014
package main

import (
	"os"
)

type HelpSource interface {
	Name() string
	Books() []Topic
	Orphans() []Topic
//	SharedAssets() []Asset
}

type Topic interface {
	Name() string
	Prepare() (string, error)
	Parent() Topic
	Children() []Topic
	Less(t Topic) bool
}

type TopicSorter []Topic

func (t TopicSorter) Len() int { return len(t) }
func (t TopicSorter) Less(i, j int) bool { return t[i].Less(t[j]) }
func (t TopicSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

var Library []Topic

func LoadLibraries() {
	m, err := OpenMSHI(os.Args[1])
	if err != nil {
		// TODO
		panic(err)
	}
	Library = append(Library, m.Books()...)
}
