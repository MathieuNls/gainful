package gainful

import "testing"

func TestAdd(t *testing.T) {

	bst := &bst{}

	bst.Add(10, newIndexable("a string"))

	if bst.root.key != 10 || bst.root.value.StringIndex() != "a string" {
		t.Error("Expected 10/a string got", bst.root)
	}

	bst.Add(15, newIndexable("another"))

	if bst.root.right.key != 15 || bst.root.right.value.StringIndex() != "another" {
		t.Error("Expected 15/another got", bst.root.right, bst.root.right.value)
	}

	bst.Add(8, newIndexable("another 8"))

	if bst.root.left.key != 8 || bst.root.left.value.StringIndex() != "another 8" {
		t.Error("Expected 8/another 8 got", bst.root.left, bst.root.left.value)
	}

	bst.Add(6, newIndexable("another 6"))

	if bst.root.left.left.key != 6 || bst.root.left.left.value.StringIndex() != "another 6" {
		t.Error("Expected 8/another 8 got", bst.root.left.left, bst.root.left.left.value)
	}

	bst.Add(16, newIndexable("another 16"))

	if bst.root.right.right.key != 16 || bst.root.right.right.value.StringIndex() != "another 16" {
		t.Error("Expected 16/another 16 got", bst.root.right.right, bst.root.right.right.value)
	}
}

func TestFetch(t *testing.T) {

	bst := &bst{}

	bst.Add(10, newIndexable("a string"))
	bst.Add(15, newIndexable("another"))
	bst.Add(8, newIndexable("another 8"))
	bst.Add(6, newIndexable("another 6"))
	bst.Add(16, newIndexable("another 16"))

	r, err := bst.Fetch(10)

	if r.key != 10 || err != nil {
		t.Error("expected 10 got", r, err)
	}

	r, err = bst.Fetch(6)

	if r.key != 6 || err != nil {
		t.Error("expected 6 got", r, err)
	}

	r, err = bst.Fetch(16)

	if r.key != 16 || err != nil {
		t.Error("expected 16 got", r, err)
	}

	r, err = bst.Fetch(99)

	if r != nil || err.Error() != "key not found" {
		t.Error("expected 'key not found'", r, err)
	}
}

func TestFloorKey(t *testing.T) {

	bst := &bst{}

	r, err := bst.FloorKey(7)

	if r != nil || err.Error() != "key not found" {
		t.Error("key not found", r, err)
	}

	bst.Add(10, newIndexable("a string"))
	bst.Add(15, newIndexable("another"))
	bst.Add(8, newIndexable("another 8"))
	bst.Add(6, newIndexable("another 6"))
	bst.Add(19, newIndexable("another 19"))

	bst.Print()

	r, err = bst.FloorKey(7)

	if r.key != 6 || err != nil {
		t.Error("expected 6 got", r, err)
	}

	r, err = bst.FloorKey(18)

	if r.key != 15 || err != nil {
		t.Error("expected 15 got", r, err)
	}

	r, err = bst.FloorKey(8)

	if r.key != 8 || err != nil {
		t.Error("expected 8 got", r, err)
	}
}

func TestAddKeys(t *testing.T) {

	ints := []int{10, 15, 8, 6, 19}
	values := []Indexable{
		newIndexable("a string"),
		newIndexable("another"),
		newIndexable("another 8"),
		newIndexable("another 6"),
		newIndexable("another 19"),
	}

	bst := FromKeys(ints, values, false)

	if bst.root.right.key != 15 || bst.root.right.value.StringIndex() != "another" {
		t.Error("Expected 15/another got", bst.root.right, bst.root.right.value)
	}

	if bst.root.left.key != 8 || bst.root.left.value.StringIndex() != "another 8" {
		t.Error("Expected 8/another 8 got", bst.root.left, bst.root.left.value)
	}

	if bst.root.left.left.key != 6 || bst.root.left.left.value.StringIndex() != "another 6" {
		t.Error("Expected 8/another 8 got", bst.root.left.left, bst.root.left.left.value)
	}

	if bst.root.right.right.key != 19 || bst.root.right.right.value.StringIndex() != "another 19" {
		t.Error("Expected 19/another 19 got", bst.root.right.right, bst.root.right.right.value)
	}

}

func TestSortedKeys(t *testing.T) {

	ints := []int{3, 5, 6, 8, 10, 15, 19}
	values := []Indexable{
		newIndexable("3"),
		newIndexable("5"),
		newIndexable("6"),
		newIndexable("8"),
		newIndexable("10"),
		newIndexable("15"),
		newIndexable("19"),
	}

	bst := FromKeys(ints, values, true)

	if bst.root.key != 8 || bst.root.value.StringIndex() != "8" {
		t.Error("Expected 8/8 got", bst.root, bst.root.value)
	}

	if bst.root.left.key != 5 || bst.root.left.value.StringIndex() != "5" {
		t.Error("Expected 5/5 got", bst.root.left, bst.root.left.value)
	}

	if bst.root.left.left.key != 3 || bst.root.left.left.value.StringIndex() != "3" {
		t.Error("Expected 3/3 got", bst.root.left.left, bst.root.left.left.value)
	}

	if bst.root.right.right.key != 19 || bst.root.right.right.value.StringIndex() != "19" {
		t.Error("Expected 19/19 got", bst.root.right.right, bst.root.right.right.value)
	}

}
