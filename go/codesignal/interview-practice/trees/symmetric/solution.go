package symmetric

type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func solution(t *Tree) bool {
	if t == nil {
		return true
	}
	return recur(t.Left, t.Right)
}

func recur(one, other *Tree) bool {
	if one == nil && other == nil {
		return true
	}
	if (one != nil && other == nil) || (other != nil && one == nil) {
		return false
	}
	if one.Value != other.Value {
		return false
	}

	if !recur(one.Left, other.Right) {
		return false
	}
	if !recur(one.Right, other.Left) {
		return false
	}

	return true
}
