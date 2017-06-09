package gainful

import (
	"errors"
	"fmt"
)

type Bst struct {
	root *binaryNode
}

func (bst *Bst) Add(key int, value interface{}) {

	if bst.root == nil {
		bst.root = &binaryNode{
			value: value,
			key:   key,
		}
	} else {
		bst.add(key, value, bst.root)
	}
}

func (bst *Bst) FromKeys(keys []int, values []interface{}, sorted bool) {

	if sorted {
		bst.root = bst.fromSortedKeys(keys, values, 0, len(keys)-1, nil)
	} else {
		for i := 0; i < len(keys); i++ {
			bst.Add(keys[i], values[i])
		}
	}
}

func (bst *Bst) FloorKey(key int) (*binaryNode, error) {

	return bst.floorKey(key, bst.root)
}

func (bst *Bst) Print() {
	bst.print(bst.root, 0)
}

func (bst *Bst) Fetch(key int) (*binaryNode, error) {

	return bst.fetch(key, bst.root)

}

func (bst *Bst) floorKey(key int, from *binaryNode) (*binaryNode, error) {

	if from != nil {
		//we found it
		if key == from.key {

			return from, nil
		} else if
		//supposed to go left or right, nothing there
		(from.left == nil && key < from.key) ||
			(from.right == nil && key >= from.key) {
			return from.parent, nil
		} else if key < from.key {
			return bst.floorKey(key, from.left)
		} else {
			return bst.floorKey(key, from.right)
		}
	}

	return nil, errors.New("key not found")
}

func (bst *Bst) fetch(key int, from *binaryNode) (*binaryNode, error) {

	if from != nil {
		if key == from.key {
			return from, nil
		} else if key < from.key {
			return bst.fetch(key, from.left)
		} else {
			return bst.fetch(key, from.right)
		}
	}

	return nil, errors.New("key not found")
}

func (bst *Bst) fromSortedKeys(keys []int, values []interface{}, start int, end int, parent *binaryNode) *binaryNode {

	if end-start >= 0 {

		mid := (start + end) / 2
		node := &binaryNode{
			key:    keys[mid],
			parent: parent,
			value:  values[mid],
		}
		node.left = bst.fromSortedKeys(keys, values, start, mid-1, node)
		node.right = bst.fromSortedKeys(keys, values, mid+1, end, node)

		return node
	}
	return nil
}

func (bst *Bst) add(key int, value interface{}, node *binaryNode) {

	if key < node.key {
		if node.left == nil {
			node.left = &binaryNode{
				parent: node,
				key:    key,
				value:  value,
			}
		} else {
			bst.add(key, value, node.left)
		}
	} else {
		if node.right == nil {
			node.right = &binaryNode{
				parent: node,
				key:    key,
				value:  value,
			}
		} else {
			bst.add(key, value, node.right)
		}
	}
}

func (bst *Bst) print(node *binaryNode, indent int) {

	if node != nil {
		bst.print(node.right, indent+4)
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Println(node.key)
		bst.print(node.left, indent+4)
	}

}
