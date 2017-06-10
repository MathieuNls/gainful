package gainful

import "testing"

import "strings"

func TestNode(t *testing.T) {

	var n *binaryNode
	n = &binaryNode{}
	n.parent = nil
	n.key = 10
	n.value = newIndexable("plop")
	n.left = nil
	n.right = nil

	if strings.Count(n.String(), "nil") != 3 || strings.Count(n.String(), "10") != 1 {
		t.Error("Expected 3 nil & 10 got", n.String())
	}
}

type indexableImpl struct {
	value string
}

func (indexable indexableImpl) StringIndex() string {
	return indexable.value
}

func newIndexable(str string) *indexableImpl {
	return &indexableImpl{
		value: str,
	}
}
