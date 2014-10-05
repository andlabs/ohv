// 5 october 2014
package main

type HelpSource interface {
	Name() string
	Books() []Book
	Orphans() []Topic
//	SharedAssets() []Asset
}

type Book interface {
	Name() string
	Members() []Topic
}

type Topic interface {
	Name() string
	Prepare() (string, error)
	Children() []Topic
	Less(t Topic) bool
}

type TopicSorter []Topic

func (t TopicSorter) Len() int { return len(t) }
func (t TopicSorter) Less(i, j int) bool { return t[i].Less(t[j]) }
func (t TopicSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
