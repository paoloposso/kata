package removek

type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func solution(t *Tree) []int {
	if t == nil {
		return []int{}
	}

	resultItems := []interface{}{}

	q := []Tree{*t}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		resultItems = append(resultItems, cur.Value)
		if cur.Left != nil {
			q = append(q, *cur.Left)
		}
		if cur.Right != nil {
			q = append(q, *cur.Right)
		}
	}

	result := []int{}

	for _, v := range resultItems {
		if val, ok := v.(int); ok {
			result = append(result, val)
		}
	}

	return result
}
