package gainful

import (
	"errors"
	"fmt"
)

//bst is a binary search tree
type bst struct {
	root *binaryNode
}

//Add adds a new key/value in the bst
func (bst *bst) Add(key int, value Indexable) {

	if bst.root == nil {
		bst.root = &binaryNode{
			value: value,
			key:   key,
		}
	} else {
		bst.add(key, value, bst.root)
	}
}

//FromKeys constructs a bst from keys array
//If sorted is true, the tree will be balanced
func FromKeys(keys []int, values []Indexable, sorted bool) *bst {

	bst := &bst{}
	if sorted {
		bst.root = bst.fromSortedKeys(keys, values, 0, len(keys)-1, nil)
	} else {
		for i := 0; i < len(keys); i++ {
			bst.Add(keys[i], values[i])
		}
	}
	return bst
}

//FloorKey returns the neareast lowest node with regards to key
func (bst *bst) FloorKey(key int) (*binaryNode, error) {

	return bst.floorKey(key, bst.root)
}

//Print prinst the bst
func (bst *bst) Print() {
	bst.print(bst.root, 0)
}

//Fetch fetches a key
func (bst *bst) Fetch(key int) (*binaryNode, error) {

	return bst.fetch(key, bst.root)

}

func (bst *bst) floorKey(key int, from *binaryNode) (*binaryNode, error) {

	if from != nil {
		//we found it
		if key == from.key {
			return from, nil
			//supposed to go right, nothing there
		} else if from.right == nil && key >= from.key {

			return from, nil
			//supposed to go left, nothing there
		} else if from.left == nil && key < from.key {
			return from.parent, nil
		} else if key < from.key {
			return bst.floorKey(key, from.left)
		} else {
			return bst.floorKey(key, from.right)
		}
	}

	return nil, errors.New("key not found")
}

func (bst *bst) fetch(key int, from *binaryNode) (*binaryNode, error) {

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

func (bst *bst) fromSortedKeys(keys []int, values []Indexable, start int, end int, parent *binaryNode) *binaryNode {

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

func (bst *bst) add(key int, value Indexable, node *binaryNode) {

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

func (bst *bst) print(node *binaryNode, indent int) {

	if node != nil {
		bst.print(node.right, indent+4)
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Println(node.key)
		bst.print(node.left, indent+4)
	}

}
