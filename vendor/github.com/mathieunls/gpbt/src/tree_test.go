package gpbt

import (
	"testing"
)

func TestAdd(t *testing.T) {

	bst := &Tree{}

	bst.Add(10, "a string")

	if bst.Root().Key != 10 || bst.Root().Value != "a string" {
		t.Error("Expected 10/a string got", bst.Root())
	}

	bst.Add(15, "another")

	if bst.Root().Right.Key != 15 || bst.Root().Right.Value != "another" {
		t.Error("Expected 15/another got", bst.Root().Right, bst.Root().Right.Value)
	}

	bst.Add(8, "another 8")

	if bst.Root().Left.Key != 8 || bst.Root().Left.Value != "another 8" {
		t.Error("Expected 8/another 8 got", bst.Root().Left, bst.Root().Left.Value)
	}

	bst.Add(6, "another 6")

	if bst.Root().Left.Left.Key != 6 || bst.Root().Left.Left.Value != "another 6" {
		t.Error("Expected 8/another 8 got", bst.Root().Left.Left, bst.Root().Left.Left.Value)
	}

	bst.Add(16, "another 16")

	if bst.Root().Right.Right.Key != 16 || bst.Root().Right.Right.Value != "another 16" {
		t.Error("Expected 16/another 16 got", bst.Root().Right.Right, bst.Root().Right.Right.Value)
	}
}

func TestFetch(t *testing.T) {

	bst := &Tree{}

	bst.Add(10, "a string")
	bst.Add(15, "another")
	bst.Add(8, "another 8")
	bst.Add(6, "another 6")
	bst.Add(16, "another 16")

	r, err := bst.Fetch(10)

	if r.Key != 10 || err != nil {
		t.Error("expected 10 got", r, err)
	}

	r, err = bst.Fetch(6)

	if r.Key != 6 || err != nil {
		t.Error("expected 6 got", r, err)
	}

	r, err = bst.Fetch(16)

	if r.Key != 16 || err != nil {
		t.Error("expected 16 got", r, err)
	}

	r, err = bst.Fetch(99)

	if r != nil || err.Error() != "Key not found" {
		t.Error("expected 'Key not found'", r, err)
	}
}

func TestFloorKey(t *testing.T) {

	bst := &Tree{}

	r, err := bst.FloorKey(7)

	if r != nil || err.Error() != "Key not found" {
		t.Error("Key not found", r, err)
	}

	bst.Add(10, "a string")
	bst.Add(15, "another")
	bst.Add(8, "another 8")
	bst.Add(6, "another 6")
	bst.Add(19, "another 19")

	bst.Print()

	r, err = bst.FloorKey(7)

	if r.Key != 6 || err != nil {
		t.Error("expected 6 got", r, err)
	}

	r, err = bst.FloorKey(18)

	if r.Key != 15 || err != nil {
		t.Error("expected 15 got", r, err)
	}

	r, err = bst.FloorKey(8)

	if r.Key != 8 || err != nil {
		t.Error("expected 8 got", r, err)
	}
}

func TestAddKeys(t *testing.T) {

	ints := []int{10, 15, 8, 6, 19}
	Values := []interface{}{
		"a string",
		"another",
		"another 8",
		"another 6",
		"another 19",
	}

	bst := NewTree(ints, Values, false)

	if bst.Root().Right.Key != 15 || bst.Root().Right.Value != "another" {
		t.Error("Expected 15/another got", bst.Root().Right, bst.Root().Right.Value)
	}

	if bst.Root().Left.Key != 8 || bst.Root().Left.Value != "another 8" {
		t.Error("Expected 8/another 8 got", bst.Root().Left, bst.Root().Left.Value)
	}

	if bst.Root().Left.Left.Key != 6 || bst.Root().Left.Left.Value != "another 6" {
		t.Error("Expected 8/another 8 got", bst.Root().Left.Left, bst.Root().Left.Left.Value)
	}

	if bst.Root().Right.Right.Key != 19 || bst.Root().Right.Right.Value != "another 19" {
		t.Error("Expected 19/another 19 got", bst.Root().Right.Right, bst.Root().Right.Right.Value)
	}

}
func TestSortedKeys(t *testing.T) {

	ints := []int{3, 5, 6, 8, 10, 15, 19}
	Values := []interface{}{
		"3",
		"5",
		"6",
		"8",
		"10",
		"15",
		"19",
	}

	bst := NewTree(ints, Values, true)

	if bst.Root().Key != 8 || bst.Root().Value != "8" {
		t.Error("Expected 8/8 got", bst.Root(), bst.Root().Value)
	}

	if bst.Root().Left.Key != 5 || bst.Root().Left.Value != "5" {
		t.Error("Expected 5/5 got", bst.Root().Left, bst.Root().Left.Value)
	}

	if bst.Root().Left.Left.Key != 3 || bst.Root().Left.Left.Value != "3" {
		t.Error("Expected 3/3 got", bst.Root().Left.Left, bst.Root().Left.Left.Value)
	}

	if bst.Root().Right.Right.Key != 19 || bst.Root().Right.Right.Value != "19" {
		t.Error("Expected 19/19 got", bst.Root().Right.Right, bst.Root().Right.Right.Value)
	}

}
