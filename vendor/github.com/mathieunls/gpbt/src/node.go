package gpbt

import (
	"strconv"
)

//Node is a node for binary trees with a int key and
//an interface value
type Node struct {
	Parent, Left, Right *Node
	Key                 int
	Value               interface{}
}

//String displays the node
func (node *Node) String() string {

	if node != nil {
		str := ""

		if node.Parent != nil {
			str += "   " + strconv.Itoa(node.Parent.Key) + "   \n"
		} else {
			str += "   nil   \n"
		}

		str += "   " + strconv.Itoa(node.Key) + "   \n"

		if node.Left != nil {
			str += " " + strconv.Itoa(node.Left.Key)
		} else {
			str += " nil"
		}
		str += " "

		if node.Right != nil {
			str += " " + strconv.Itoa(node.Right.Key)
		} else {
			str += " nil"
		}

		return str
	}

	return "nil"
}
