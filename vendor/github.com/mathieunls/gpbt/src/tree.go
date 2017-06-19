package gpbt

import (
	"errors"
	"fmt"
)

//NavigableTree is a kind of tree that can return key closest
//to the demanded key
type NavigableTree interface {
	Add(key int, value interface{})
	FloorKey(key int) (*Node, error)
	Print()
	Fetch(key int) (*Node, error)
	Root() *Node
}

//Tree is a binary search tree
//with a recusive add method
type Tree struct {
	root *Node
}

//Root returns the root
func (tree *Tree) Root() *Node {
	return tree.root
}

//Add adds a new key/value in the tree
//Complexity is O(log(n)) where N is a the number
//of key in the tree
func (tree *Tree) Add(key int, value interface{}) {

	if tree.root == nil {
		tree.root = &Node{
			Value: value,
			Key:   key,
		}
	} else {
		tree.add(key, value, tree.root)
	}
}

//NewTree constructs a tree from keys array
//If sorted is true, the tree will be balanced
func NewTree(keys []int, values []interface{}, sorted bool) NavigableTree {

	tree := &Tree{}
	if sorted {
		tree.root = tree.fromSortedKeys(keys, values, 0, len(keys)-1, nil)
	} else {
		for i := 0; i < len(keys); i++ {
			tree.Add(keys[i], values[i])
		}
	}
	return tree
}

//FloorKey returns the neareast lowest node with regards to key
//Complexity is O(log(n)) where N is a the number
//of key in the tree
func (tree *Tree) FloorKey(key int) (*Node, error) {

	return tree.floorKey(key, tree.root)
}

//Print prinst the tree
func (tree *Tree) Print() {
	tree.print(tree.root, 0)
}

//Fetch fetches a key
//Complexity is O(log(n)) where N is a the number
//of key in the tree
func (tree *Tree) Fetch(key int) (*Node, error) {

	return tree.fetch(key, tree.root)
}

func (tree *Tree) floorKey(key int, from *Node) (*Node, error) {

	if from != nil {
		//we found it
		if key == from.Key {
			return from, nil
			//supposed to go right, nothing there
		} else if from.Right == nil && key >= from.Key {

			return from, nil
			//supposed to go left, nothing there
		} else if from.Left == nil && key < from.Key {
			return from.Parent, nil
		} else if key < from.Key {
			return tree.floorKey(key, from.Left)
		} else {
			return tree.floorKey(key, from.Right)
		}
	}

	return nil, errors.New("Key not found")
}

func (tree *Tree) fetch(key int, from *Node) (*Node, error) {

	if from != nil {
		if key == from.Key {
			return from, nil
		} else if key < from.Key {
			return tree.fetch(key, from.Left)
		} else {
			return tree.fetch(key, from.Right)
		}
	}

	return nil, errors.New("Key not found")
}

func (tree *Tree) fromSortedKeys(keys []int, values []interface{}, start int, end int, parent *Node) *Node {

	if end-start >= 0 {

		mid := (start + end) / 2
		node := &Node{
			Key:    keys[mid],
			Parent: parent,
			Value:  values[mid],
		}
		node.Left = tree.fromSortedKeys(keys, values, start, mid-1, node)
		node.Right = tree.fromSortedKeys(keys, values, mid+1, end, node)

		return node
	}
	return nil
}

func (tree *Tree) add(key int, value interface{}, node *Node) {

	if key < node.Key {
		if node.Left == nil {
			node.Left = &Node{
				Parent: node,
				Key:    key,
				Value:  value,
			}
		} else {
			tree.add(key, value, node.Left)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{
				Parent: node,
				Key:    key,
				Value:  value,
			}
		} else {
			tree.add(key, value, node.Right)
		}
	}
}

func (tree *Tree) print(node *Node, indent int) {

	if node != nil {
		tree.print(node.Right, indent+4)
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Println(node.Key)
		tree.print(node.Left, indent+4)
	}

}
