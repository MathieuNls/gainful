package gpbt

import (
	"strconv"
	"testing"
)

func TestNewPTree(t *testing.T) {

	ints := make([]int, 30)
	values := make([]interface{}, 30)

	for i := 0; i < 30; i++ {
		ints[i] = i * 3
		values[i] = strconv.Itoa(i * 3)
	}

	tree := NewParralelTree(ints, values, 12)

	node, err := tree.Fetch(81)

	if node.Key != 81 || node.Parent.Key != 78 || node.Right != nil || node.Left != nil || err != nil {
		t.Error("Expected 81 got", node, err)
	}

	node, err = tree.Fetch(3)

	if node.Key != 3 || node.Parent != nil || node.Right.Key != 6 || node.Left.Key != 0 || err != nil {
		t.Error("Expected 3 got", node, err)
	}

	node, err = tree.Fetch(73)

	if err == nil || node != nil {
		t.Error("Expected error got", node, err)
	}

	node, err = tree.FloorKey(62)

	if node.Key != 60 || node.Parent != nil || node.Right.Key != 63 || node.Left != nil || err != nil {
		t.Error("Expected 60 got", node, err)
	}

	node, err = tree.FloorKey(13)

	if node.Key != 12 || node.Parent != nil || node.Right.Key != 15 || node.Left.Key != 9 || err != nil {
		t.Error("Expected 12 got", node, err)
	}

	tree.Add(99, "99")

	node, err = tree.Fetch(99)

	if node.Key != 99 || err != nil {
		t.Error("Expected 99 got", node, err)
	}

	tree = NewParralelTree([]int{}, []interface{}{}, -1)

	r, err := tree.FloorKey(7)

	if r != nil || err.Error() != "Key not found" {
		t.Error("Key not found", r, err)
	}

	r, err = tree.Fetch(7)

	if r != nil || err.Error() != "Key not found" {
		t.Error("Key not found", r, err)
	}
}
