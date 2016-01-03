// 2 january 2016
package main

// #include "ui.h"
import "C"

type TreeModelRootNodes interface {
	RootNodes() []TreeNode
}

type TreeNode interface {
	TreeNodeText() string
	// TODO readd TreeNodeParent()?
	TreeNodeChildren() []TreeNode
}

type TreeModel struct {
	root			TreeModelRootNodes
	id			C.id
	trees			map[*Tree]struct{}
	nodeObjects	map[TreeNode]C.id
	objectNodes	map[C.id]TreeNode
}

var treeModels = make(map[C.id]*TreeModel)

func NewTreeModel(root TreeModelRootNodes) *TreeModel {
	m := new(TreeModel)
	m.root = root
	m.id = C.newTreeModel()
	m.trees = make(map[*Tree]struct{})
	m.nodeObjects = make(map[TreeNode]C.id)
	m.objectNodes = make(map[C.id]TreeNode)
	treeModels[m.id] = m
	return m
}

func (m *TreeModel) Destroy() {
	if len(m.trees) != 0 {
		panic("attempt to destroy TreeModel while still attached to Trees")
	}
	for id, _ := range m.objectNodes {
		C.treeNodeDestroy(id)
	}
	C.treeModelDestroy(m.id)
}

func (m *TreeModel) RowInserted(node TreeNode, parent TreeNode, index int) {
	id := C.newTreeNode()
	m.nodeObjects[node] = id
	m.objectNodes[id] = node

	parentid := C.id(nil)
	if parent != nil {
		parentid = m.nodeObjects[parent]
	}
	for t, _ := range m.trees {
		C.treeInsertRow(t.id, parentid, C.intmax_t(index))
	}
}

func (m *TreeModel) NodeChanged(node TreeNode) {
	id := m.nodeObjects[node]
	for t, _ := range m.trees {
		C.treeUpdateNode(t.id, id)
	}
}

func (m *TreeModel) RowDeleted(node TreeNode, parent TreeNode, index int) {
	id := m.nodeObjects[node]
	delete(m.nodeObjects, node)
	delete(m.objectNodes, id)
	C.treeNodeDestroy(id)

	parentid := C.id(nil)
	if parent != nil {
		parentid = m.nodeObjects[parent]
	}
	for t, _ := range m.trees {
		C.treeDeleteRow(t.id, parentid, C.intmax_t(index))
	}
}

//export treeModelChild
func treeModelChild(mm C.id, index C.intmax_t, nodeobj C.id) C.id {
	var child TreeNode

	m := treeModels[mm]
	if nodeobj == nil {
		child = m.root.RootNodes()[index]
	} else {
		node := m.objectNodes[nodeobj]
		child = node.TreeNodeChildren()[index]
	}
	return m.nodeObjects[child]
}

//export treeModelChildCount
func treeModelChildCount(mm C.id, nodeobj C.id) C.intmax_t {
	m := treeModels[mm]
	if nodeobj == nil {
		return len(m.root.RootNodes())
	}
	node := m.objectNodes[nodeobj]
	return len(node.TreeNodeChildren())
}

//export treeModelNodeText
func treeModelNodeText(mm C.id, nodeobj C.id) *C.char {
	m := treeModels[mm]
	node := m.objectNodes[nodeobj]
	return C.CString(node.TreeNodeText())		// freed on the C side
}
