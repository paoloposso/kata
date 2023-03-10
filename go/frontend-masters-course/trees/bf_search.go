package trees

func Bfs(head BinaryNode, needle int32) bool {
	q := []BinaryNode{head}

	for len(q) > 0 {
		// pop
		curr := q[0]
		q = q[1:]

		if curr.Value == needle {
			return true
		}

		//push children
		if curr.Left != nil {
			q = append(q, *curr.Left)
		}
		if curr.Right != nil {
			q = append(q, *curr.Right)
		}
	}

	return false
}

func insert(head *BinaryNode, value int32) {
	newNode := &BinaryNode{Value: value}
	if head == nil {
		head = newNode
		return
	}

	insertRecur(head, newNode)
}

func insertRecur(curr, newNode *BinaryNode) {
	if newNode.Value <= curr.Value {
		if curr.Left == nil {
			curr.Left = newNode
			return
		}
		insertRecur(curr.Left, newNode)
	} else if newNode.Value >= curr.Value {
		if curr.Right == nil {
			curr.Right = newNode
			return
		}
		insertRecur(curr.Right, newNode)
	}
}
