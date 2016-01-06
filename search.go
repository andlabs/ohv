// 5 january 2016
package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
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

func NewSearcher() {
	// TODO
}

func (s *Searcher) Close() {
	// TODO
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
