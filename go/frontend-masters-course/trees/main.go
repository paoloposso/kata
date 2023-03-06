package main

import "fmt"

type Tree struct {
	Head *BynaryNode
}

type BynaryNode struct {
	Value int32
	Left  *BynaryNode
	Right *BynaryNode
}

func walk(curr *BynaryNode, path []int32) []int32 {
	//base case
	if curr == nil {
		return path
	}

	//pre
	path = append(path, curr.Value)

	//recurse
	path = walk(curr.Left, path)
	path = walk(curr.Right, path)

	//post
	return path
}

func preOrderSearch(head BynaryNode) []int32 {
	return walk(&head, []int32{})
}

func main() {
	t := Tree{}

	head := t.Head

	head = &BynaryNode{Value: 1}
	head.Left = &BynaryNode{Value: 5}
	head.Right = &BynaryNode{Value: 20}

	curr := head.Left
	curr.Left = &BynaryNode{Value: 18}
	curr.Right = &BynaryNode{Value: 30}

	curr = curr.Right
	curr.Left = &BynaryNode{Value: 2}
	curr.Right = &BynaryNode{Value: 4}

	curr = curr.Right
	curr.Left = &BynaryNode{Value: 50}
	curr.Right = &BynaryNode{Value: 44}

	curr = head.Right
	curr.Left = &BynaryNode{Value: 18}
	curr.Right = &BynaryNode{Value: 30}

	curr = curr.Left
	curr.Left = &BynaryNode{Value: 1}
	curr.Right = &BynaryNode{Value: 11}

	curr = curr.Right
	curr.Left = &BynaryNode{Value: 1}
	curr.Right = &BynaryNode{Value: 11}

	curr = curr.Left
	curr.Left = &BynaryNode{Value: 18}
	curr.Right = &BynaryNode{Value: 55}

	fmt.Printf("%v", preOrderSearch(*head))
}
