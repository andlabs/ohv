// 5 january 2016
package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/andlabs/ohv/internal/ui"
)

// TODO settings? explict _ separator?
const searchCreateTable = `
CREATE VIRTUAL TABLE searchTopics USING fts4 (
	title,

	tokenize = unicode61
);
`

const searchInsert = `
INSERT INTO searchTopics(title) VALUES (?);
`

// TODO ORDER BY rank?
const searchQuery = `
SELECT rowid
	FROM searchTopics
	WHERE searchTopics MATCH ?;
`

type Searcher struct {
	db		*sql.DB
	insert	*sql.Stmt
	query	*sql.Stmt
	topics	map[int64]Topic
}

func NewSearcher() (s *Searcher, err error) {
	s = new(Searcher)
	s.topics = make(map[int64]Topic)

	s.db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	create, err := s.db.Prepare(searchCreateTable)
	if err != nil {
		s.db.Close()
		return nil, err
	}
	_, err = create.Exec()
	if err != nil {
		create.Close()
		s.db.Close()
		return nil, err
	}
	create.Close()

	s.insert, err = s.db.Prepare(searchInsert)
	if err != nil {
		s.db.Close()
		return nil, err
	}

	s.query, err = s.db.Prepare(searchQuery)
	if err != nil {
		s.insert.Close()
		s.db.Close()
		return nil, err
	}

	return s, nil
}

func (s *Searcher) Close() {
	s.query.Close()
	s.insert.Close()
	s.db.Close()
}

func (s *Searcher) add(topic Topic) error {
	result, err := s.insert.Exec(topic.Name())
	if err != nil {
		return err
	}
	n, err := result.LastInsertId()
	if err != nil {
		return err
	}
	s.topics[n] = topic
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
	r, err := s.query.Query(searchFor)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	sr := new(SearchResults)
	sr.model = ui.NewTreeModel(sr)

	for r.Next() {
		var id int64

		err = r.Scan(&id)
		if err != nil {
			sr.model.Destroy()
			return nil, err
		}
		node := &searchResult{s.topics[id]}
		sr.roots = append(sr.roots, node)
		sr.model.RowInserted(node, nil, len(sr.roots) - 1)
	}
	err = r.Err()
	if err != nil {
		sr.model.Destroy()
		return nil, err
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
