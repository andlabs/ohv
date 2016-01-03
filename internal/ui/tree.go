// 2 january 2016
package ui

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
