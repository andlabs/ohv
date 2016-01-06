// 7 october 2014
package main

// TODO read other constant names since constants in a group are not listed individually

import (
	"path/filepath"
	"net/url"
	"strings"
	"sort"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/andlabs/ohv/internal/ui"
)

// TODO see if selecting the count and preallocating makes things faster
const appleAllInfo = `
SELECT znode.z_pk, znode.zkname, znodeurl.zpath, znodeurl.zanchor, zorderedsubnode.zparent, zorderedsubnode.zorder
	FROM znode
		LEFT JOIN znodeurl
			ON znode.z_pk = znodeurl.znode
		LEFT JOIN zorderedsubnode
			ON znode.z_pk = zorderedsubnode.znode
			AND znode.zprimaryparent = zorderedsubnode.zparent;
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
	orphans	[]Topic
}

func (a *Apple) Name() string {
	return "Apple documentation collection " + filepath.Base(a.dir)
}

func (a *Apple) Books() []Topic {
	return a.books
}

func (a *Apple) Orphans() []Topic {
	return a.orphans
}

func (a *Apple) Lookup(url *url.URL) Topic {
	// TODO
	return nil
}

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
	db.MapperFunc(strings.ToUpper)		// struct field names are uppercase
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

	a.topics = make(map[int64]*AppleTopic)
	for _, n := range a.nodes {
		if a.topics[n.Z_PK] != nil {
			panic("duplicate IDs")
		}
		topic := &AppleTopic{
			dir:		a.dircr,		// note: not a.dir
			name:	n.ZKNAME,
			pptr:		n.ZPARENT,
			source:	a,
		}
		if n.ZPATH != nil {
			// TODO
			topic.path = *n.ZPATH
		}
		if n.ZANCHOR != nil {
			topic.anchor = *n.ZANCHOR
		}
		if n.ZORDER != nil {
			// ZORDER is 1-based
			topic.order = *n.ZORDER - 1
		}
		a.topics[n.Z_PK] = topic
	}

	for _, v := range a.topics {
		if v.pptr == nil {		// is top-level
			a.books = append(a.books, v)
			continue
		}
		parent, ok := a.topics[*v.pptr]
		if ok {			// has parent
			parent.children = append(parent.children, v)
			v.parent = parent
			continue
		}
		a.orphans = append(a.orphans, v)			// parent elsewhere or missing
	}

	for _, t := range a.topics {
		sort.Sort(TopicSorter(t.children))
	}

	return a, nil
}

type AppleTopic struct {
	dir		string
	name	string
	path		string
	anchor	string
	pptr		*int64
	order	int64
	parent	Topic
	children	[]Topic
	source	HelpSource
}

func (a *AppleTopic) Name() string {
	return a.name
}

func (a *AppleTopic) Prepare() (*Prepared, error) {
	// TODO cache real paths
	cssdir := filepath.Join(a.dir, "Documents", "Resources", "1014", "CSS")
	return &Prepared{
		Path:			filepath.Join(a.dir, "Documents", a.path),
		Anchor:		a.anchor,
		CSSPath:		filepath.Join(cssdir, "xcode5.css"),
		CSSBaseDir:	cssdir,
	}, nil
}

func (a *AppleTopic) Parent() Topic {
	return a.parent
}

func (a *AppleTopic) Children() []Topic {
	return a.children
}

func (a *AppleTopic) Source() HelpSource {
	return a.source
}

func (a *AppleTopic) Less(t Topic) bool {
	tt := t.(*AppleTopic)
	return a.order < tt.order
}

func (a *AppleTopic) TreeNodeText() string {
	return a.Name()
}

func (a *AppleTopic) TreeNodeParent() ui.TreeNode {
	return a.Parent()
}

func (a *AppleTopic) TreeNodeChildren() []ui.TreeNode {
	if len(a.children) == 0 {
		return nil
	}
	c := make([]ui.TreeNode, len(a.children))
	for i, n := range a.children {
		c[i] = n
	}
	return c
}
