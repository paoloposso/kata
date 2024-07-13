package haspathgivensum

type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func solution(t *Tree, s int) bool {
	currentSum := 0

	if t == nil {
		return false
	}

	currentSum = t.Value.(int)

	return recur(t, currentSum, s)
}

func recur(node *Tree, currentSum, s int) bool {
	if node.Left == nil && node.Right == nil { //is leaf
		return currentSum == s
	}

	if node.Left != nil {
		if recur(node.Left, currentSum+node.Left.Value.(int), s) {
			return true
		}
	}
	if node.Right != nil {
		if recur(node.Right, currentSum+node.Right.Value.(int), s) {
			return true
		}
	}

	return false
}
