package binary_search_tree

import (
	"errors"
	"fmt"
)

type binNode struct {
	Data  int
	Left  *binNode
	Right *binNode
}

type binTree struct {
	Root  *binNode
	count int
}

func newBinNode() *binNode {
	return &binNode{}
}

func newBinTree() *binTree {
	tree := &binTree{nil, 0}
	return tree
}

func (bt *binTree) Search(item int) (bool, error) {
	if bt.Root == nil {
		return false, errors.New("tree is empty")
	}
	node := bt.Root
	for {
		switch {
		case node == nil:
			return false, nil
		case item == node.Data:
			return true, nil
		case item > node.Data:
			node = node.Right
		default:
			node = node.Left
		}
	}
}

func (bt *binTree) Insert(item int) (bool, error) {
	var node *binNode
	var newNode **binNode // Адрес указателя на корень дерева
	newNode = &bt.Root
	node = bt.Root
	for {
		switch {
		case node == nil:
			*newNode = newBinNode()
			node = *newNode
			node.Data = item
			node.Left, node.Right = nil, nil
			bt.count++
			return true, nil
		case item == node.Data:
			return false, errors.New(fmt.Sprintf("item %v already exists in tree", item))
		case item > node.Data:
			newNode = &node.Right
			node = node.Right
		default:
			newNode = &node.Left
			node = node.Left
		}
	}

}

func (bt *binTree) Delete(item int) bool {
	var node *binNode
	var delNode **binNode // Адрес указателя на корень дерева
	delNode = &bt.Root
	node = bt.Root
Loop:
	for {
		switch {
		case node == nil:
			return false
		case item == node.Data:
			break Loop
		case item > node.Data:
			delNode = &node.Right
			node = node.Right
		default:
			delNode = &node.Left
			node = node.Left
		}
	}

	if node.Right == nil {
		*delNode = node.Left
	} else {
		var y *binNode
		y = node.Right
		if y.Left == nil {
			y.Left = node.Left
			*delNode = y
		} else {
			var x *binNode
			x = y.Left
			for x.Left != nil {
				y = x
				x = y.Left
			}
			y.Left = x.Right
			x.Left = node.Left
			x.Right = node.Right
			*delNode = x
		}
	}

	bt.count--
	return true
}

func (bt *binTree) recursiveWalk(node *binNode) {
	if node == nil {
		return
	}
	bt.recursiveWalk(node.Left)
	fmt.Printf("%v ", node.Data)
	bt.recursiveWalk(node.Right)
}

func (bt *binTree) RecursiveBinWalk() {
	bt.recursiveWalk(bt.Root)
}

func (bt *binTree) iterWalk(node *binNode) {
	stack := make([]*binNode, 10, 10)
	var count int
	for {
		for node != nil {
			stack[count] = node
			count++
			node = node.Left
		}
		if count == 0 {
			return
		}
		count--
		node = stack[count]
		fmt.Printf("%v ", (*node).Data)
		node = node.Right
	}
}

func (bt *binTree) IterBinWalk() {
	bt.iterWalk(bt.Root)
}

func (bt *binTree) NodeCount() int {
	return bt.count
}
