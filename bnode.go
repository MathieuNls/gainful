package gainful

import (
	"strconv"
)

type binaryNode struct {
	parent, left, right *binaryNode
	key                 int
	value               Indexable
}

func (node *binaryNode) String() string {

	if node != nil {
		str := ""

		if node.parent != nil {
			str += "   " + strconv.Itoa(node.parent.key) + "   \n"
		} else {
			str += "   nil   \n"
		}

		str += "   " + strconv.Itoa(node.key) + "   \n"

		if node.left != nil {
			str += " " + strconv.Itoa(node.left.key)
		} else {
			str += " nil"
		}
		str += " "

		if node.right != nil {
			str += " " + strconv.Itoa(node.right.key)
		} else {
			str += " nil"
		}

		return str
	}

	return "nil"
}
