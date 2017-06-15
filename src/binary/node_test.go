package binary

import (
	"strings"
	"testing"

	"github.com/mathieunls/gainful/src/indexable"
)

func TestNode(t *testing.T) {

	var n *Node
	n = &Node{}
	n.Parent = nil
	n.Key = 10
	n.Value = indexable.New("plop")
	n.Left = nil
	n.Right = nil

	if strings.Count(n.String(), "nil") != 3 || strings.Count(n.String(), "10") != 1 {
		t.Error("Expected 3 nil & 10 got", n.String())
	}
}
