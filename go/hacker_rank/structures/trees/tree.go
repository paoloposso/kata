package trees

type Node struct {
	Right *Node
	Left  *Node
	Value int
}

func Insert(head *Node, value int) *Node {
	if head == nil {
		return &Node{Value: value}
	}

	if value >= head.Value {
		head.Right = Insert(head.Right, value)
	} else {
		head.Left = Insert(head.Left, value)
	}

	return head
}
