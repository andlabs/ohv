// 2 january 2016
package ui

import (
	"net/url"
//	"time"
	"testing"
)

type node struct {
	text		string
	children	[]TreeNode
	model	*TreeModel
}

func (n *node) TreeNodeText() string {
	return n.text
}

func (n *node) TreeNodeChildren() []TreeNode {
	return n.children
}

func (n *node) AddChild(text string) *node {
	n2 := new(node)
	n2.text = text
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

func newRoot(text string) *node {
	n := new(node)
	n.text = text
	n.model = roots.model
	roots.nodes = append(roots.nodes, n)
	n.model.RowInserted(n, nil, len(roots.nodes) - 1)
	return n
}

func populate() {
	QueueMain(func() {
		n := newRoot("Root 1")
			n2 := n.AddChild("Child 1")
			n2 = n.AddChild("Child 2")
				n2.AddChild("Subchild")
			n2 = n.AddChild("Child 3")
		n = newRoot("Root 2")
		n = newRoot("Root 3")
			n2 = n.AddChild("Only One Child")
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
		split := NewSplitter(t, wv)
		split.SetPosition(20)
		w.SetChild(split)
		w.Show()
		go populate()
		QueueMain(func() {
			u, _ := url.Parse("http://google.com/")
			wv.Navigate(u)
		})
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
}
