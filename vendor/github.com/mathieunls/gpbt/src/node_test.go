package gpbt

import (
	"strings"
	"testing"
)

func TestNode(t *testing.T) {

	var n *Node
	n = &Node{}
	n.Parent = nil
	n.Key = 10
	n.Value = "plop"
	n.Left = nil
	n.Right = nil

	if strings.Count(n.String(), "nil") != 3 || strings.Count(n.String(), "10") != 1 {
		t.Error("Expected 3 nil & 10 got", n.String())
	}

	n.Left = &Node{Value: 10, Parent: n}
	n.Right = &Node{Value: 10, Parent: n}
	n.Parent = &Node{Value: 8}

	if len(n.String()) == 0 {
		t.Error("Expected 2 10 got\n", n.String())
	}

	n = nil
	n.String()
}
