package preordersearch

import (
	"fmt"
	"testing"
	"trees/datastruct"
)

func TestPreOrderSearch(t *testing.T) {
	tr := datastruct.BinaryTree{}

	head := tr.Head

	head = &datastruct.BynaryNode{Value: 1}
	head.Left = &datastruct.BynaryNode{Value: 5}
	head.Right = &datastruct.BynaryNode{Value: 20}

	curr := head.Left
	curr.Left = &datastruct.BynaryNode{Value: 18}
	curr.Right = &datastruct.BynaryNode{Value: 30}

	curr = curr.Right
	curr.Left = &datastruct.BynaryNode{Value: 2}
	curr.Right = &datastruct.BynaryNode{Value: 4}

	curr = curr.Right
	curr.Left = &datastruct.BynaryNode{Value: 50}
	curr.Right = &datastruct.BynaryNode{Value: 44}

	curr = head.Right
	curr.Left = &datastruct.BynaryNode{Value: 18}
	curr.Right = &datastruct.BynaryNode{Value: 30}

	curr = curr.Left
	curr.Left = &datastruct.BynaryNode{Value: 1}
	curr.Right = &datastruct.BynaryNode{Value: 11}

	curr = curr.Right
	curr.Left = &datastruct.BynaryNode{Value: 1}
	curr.Right = &datastruct.BynaryNode{Value: 11}

	curr = curr.Left
	curr.Left = &datastruct.BynaryNode{Value: 18}
	curr.Right = &datastruct.BynaryNode{Value: 55}

	fmt.Printf("%v", PreOrderSearch(*head))
}
