// 2 january 2016
package ui

// TODO
// - tree should horizontally scroll if cells aren't wide enough

// #include "ui.h"
import "C"

type Tree struct {
	id			C.id
	model		*TreeModel
	onSelected	func()
}

var trees = make(map[C.id]*Tree)

func NewTree() *Tree {
	t := new(Tree)
	t.id = C.newTree()
	trees[t.id] = t
	// this is needed for the selection event
	trees[C.treeOutlineView(t.id)] = t
	return t
}

func (t *Tree) Destroy() {
	t.SetModel(nil)
	delete(trees, C.treeOutlineView(t.id))
	delete(trees, t.id)
	C.treeDestroy(t.id)
}

func (t *Tree) Handle() uintptr {
	return touintptr(t.id)
}

func (t *Tree) SetModel(model *TreeModel) {
	if t.model != nil {
		delete(t.model.trees, t)
	}
	t.model = model
	if t.model != nil {
		t.model.trees[t] = struct{}{}
		C.treeSetModel(t.id, t.model.id)
	} else {
		C.treeSetModel(t.id, nil)
	}
}

func (t *Tree) Selected() TreeNode {
	if t.model == nil {
		return nil
	}
	nodeobj := C.treeSelected(t.id)
	return t.model.objectNodes[nodeobj]
}

// in order to be able to access a node, its parent and all the parents leading up to the node must be open
// there doesn't seem to be a way to do this in Cocoa so we have to do it ourselves
// note that we do not expand node itself
func (t *Tree) ensureCanAccessNode(node TreeNode) {
	parent := node.TreeNodeParent()
	if parent == nil {		// the root is always open
		return
	}
	// first make sure the parent of the parent is expanded
	t.ensureCanAccessNode(parent)
	// and expand the parent
	parentobj := t.model.nodeObjects[parent]
	C.treeExpandItem(t.id, parentobj)
}

// TODO handle nil
func (t *Tree) SetSelected(node TreeNode) {
	// first we have to make sure the node can be accessed
	t.ensureCanAccessNode(node)
	// then we can select it
	nodeobj := t.model.nodeObjects[node]
	C.treeSetSelected(t.id, nodeobj)
	// TODO avoid raising an OnSelected
}

func (t *Tree) OnSelected(f func()) {
	t.onSelected = f
}

//export doTreeSelected
func doTreeSelected(tt C.id) {
	t := trees[tt]
	if t.onSelected != nil {
		t.onSelected()
	}
}
