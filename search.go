// 5 january 2016
package main

import (
	"fmt"

	"github.com/andlabs/ohv/internal/ui"
)

type Searcher struct {
	s		*ui.SearchIndex
	n		int64
	topics	map[string]Topic
}

func NewSearcher() (s *Searcher, err error) {
	s = new(Searcher)
	s.topics = make(map[string]Topic)
	s.s = ui.NewSearchIndex()
	return s, nil
}

func (s *Searcher) Close() {
	// TODO
}

func (s *Searcher) add(topic Topic) error {
	ns := fmt.Sprint(s.n)
	s.n++
	s.s.Add(ns, topic.Name())
	s.topics[ns] = topic
	return nil
}

func (s *Searcher) indexTopics(topics []Topic) error {
	for _, t := range topics {
		err := s.add(t)
		if err != nil {
			return err
		}
		err = s.indexTopics(t.Children())
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Searcher) Index() error {
	return s.indexTopics(Library.Topics())
}

type searchResult struct {
	t		Topic
}

func (s *searchResult) TreeNodeText() string {
	return s.t.TreeNodeText()
}

func (s *searchResult) TreeNodeParent() ui.TreeNode {
	return nil
}

func (s *searchResult) TreeNodeChildren() []ui.TreeNode {
	return nil
}

type SearchResults struct {
	roots	[]ui.TreeNode
	model	*ui.TreeModel
}

func (s *SearchResults) RootNodes() []ui.TreeNode {
	return s.roots
}

func (s *Searcher) Search(searchFor string) (*SearchResults, error) {
	sr := new(SearchResults)
	sr.model = ui.NewTreeModel(sr)
	r := s.s.Search(searchFor)
	defer r.Dismiss()
	for r.Next() {
		id := r.Result()
		node := &searchResult{s.topics[id]}
		sr.roots = append(sr.roots, node)
		sr.model.RowInserted(node, nil, len(sr.roots) - 1)
	}
	return sr, nil
}

func (sr *SearchResults) Dismiss() {
	sr.model.Destroy()
}

func (sr *SearchResults) Model() *ui.TreeModel {
	return sr.model
}

func (sr *SearchResults) Topic(node ui.TreeNode) Topic {
	return node.(*searchResult).t
}
