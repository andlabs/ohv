// 2 january 2016
package ui

// #include "ui.h"
import "C"

type Tree struct {
	id		C.id
	model	*TreeModel
}

var trees = make(map[C.id]*Tree)

func NewTree() *Tree {
	t := new(Tree)
	t.id = C.newTree()
	trees[t.id] = t
	return t
}

func (t *Tree) Destroy() {
	t.SetModel(nil)
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
