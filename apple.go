// 7 october 2014
package main

import (
	"path/filepath"
	"net/url"
"strings"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// TODO see if selecting the count and preallocating makes things faster
const appleAllInfo = `
SELECT znode.z_pk, znode.zkname, znodeurl.zpath, znodeurl.zanchor, zorderedsubnode.zparent, zorderedsubnode.zorder
	FROM znode
		LEFT JOIN znodeurl ON znode.z_pk = znodeurl.znode
		LEFT JOIN zorderedsubnode ON znode.z_pk = zorderedsubnode.znode;
`

type appleNode struct {
	Z_PK			int64
	ZKNAME		string
	// TODO this next field shouldn't be nil...?????
	ZPATH		*string		// three cheers for SQL!!!!111one
	ZANCHOR		*string
	ZPARENT		*int64
	ZORDER		*int64		// 1-based!
}

type Apple struct {
	dir		string
	dircr		string
	nodes	[]appleNode
	topics	map[int64]*AppleTopic
	books	[]Topic
	// TODO orphans
}

func (a *Apple) Name() string {
	return "Apple documentation collection " + filepath.Base(a.dir)
}

func (a *Apple) Books() []Topic {
	return a.books
}

func (a *Apple) Orphans() []Topic {
	// TODO
	return nil
}

func (a *Apple) Lookup(url *url.URL) Topic {
	// TODO
	return nil
}

// TODO adorn errors
func OpenApple(dir string) (*Apple, error) {
	a := new(Apple)
	a.dir = dir
	a.dircr = filepath.Join(a.dir, "Contents", "Resources")

	dbname := filepath.Join(a.dircr, "docSet.dsidx")
	db, err := sqlx.Connect("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	defer db.Close()
db.MapperFunc(strings.ToUpper)//TODO this is needed for some reason
	p, err := db.Preparex(appleAllInfo)
	if err != nil {
		return nil, err
	}
	err = p.Select(&a.nodes)
	if err != nil {
		return nil, err
	}
	err = p.Close()
	if err != nil {
		return nil, err
	}

	return a, nil
}

type AppleTopic struct {
}
