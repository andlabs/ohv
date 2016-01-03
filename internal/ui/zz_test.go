// 2 january 2016
package ui

import (
	"net/url"
//	"time"
	"testing"
)

type node struct {
	text		string
	url		string
	children	[]TreeNode
	model	*TreeModel
}

func (n *node) TreeNodeText() string {
	return n.text
}

func (n *node) TreeNodeChildren() []TreeNode {
	return n.children
}

func (n *node) AddChild(text string, url string) *node {
	n2 := new(node)
	n2.text = text
	n2.url = url
	n2.model = n.model
	n.children = append(n.children, n2)
	n.model.RowInserted(n2, n, len(n.children) - 1)
	return n2
}

type root struct {
	nodes	[]TreeNode
	model	*TreeModel
}

func (r *root) RootNodes() []TreeNode {
	return r.nodes
}

var roots = new(root)

func newRoot(text string, url string) *node {
	n := new(node)
	n.text = text
	n.url = url
	n.model = roots.model
	roots.nodes = append(roots.nodes, n)
	n.model.RowInserted(n, nil, len(roots.nodes) - 1)
	return n
}

func populate() {
	QueueMain(func() {
		n := newRoot("Root 1", "http://google.com")
			n2 := n.AddChild("Child 1", "http://apple.com/")
			n2 = n.AddChild("Child 2", "http://golang.org/")
				n2.AddChild("Subchild", "http://mozilla.org/")
			n2 = n.AddChild("Child 3", "http://swtch.com/")
		n = newRoot("Root 2", "http://aaljgrhljljgrsahr.rgoajgrs/")
		n = newRoot("Root 3", "http://plan9.bell-labs.com/")
			n2 = n.AddChild("Only One Child", "http://kernel.org/")

		n = new(node)
		n.text = "A"
		n.url = "http://i.imgur.com/3ZtlLKR.png"
		n.model = n2.model
		n2 = new(node)
		n2.text = "B"
		n2.url = "http://tilde.town/~sanqui/blog/00.html"
		n2.model = n.model
		n.children = append(n.children, n2)
		roots.nodes = append(roots.nodes, n)
		n.model.RowInserted(n, nil, len(roots.nodes) - 1)
	})
}

func TestIt(t *testing.T) {
	err := Main(func() {
		w := NewWindow("Test", 640, 480)
		w.OnClosing(func() bool {
			Quit()
			return true
		})
		w.Move(100, 100)
		t := NewTree()
		tm := NewTreeModel(roots)
		roots.model = tm
		t.SetModel(tm)
		wv := NewWebView()
		wv.OnLoadFailed(func(sysError uintptr) {
			w.MsgBoxSysError(sysError)
		})
		t.OnSelected(func() {
			n, ok := t.Selected().(*node)
			if !ok { return }
			u, _ := url.Parse(n.url)
			wv.Navigate(u)
		})
		se := NewSearchEntry()
		split := NewSplitter(NewBox(NewMargin(se), t), wv)
		split.SetPosition(20)
		w.SetChild(split)
		se.OnChanged(func() {
			println(se.Text())
		})
		w.Show()
		go populate()
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
