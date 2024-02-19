package islucky

func solution(a []int) []int {
	sortRecur(a, 0)

	return a
}

func sortRecur(a []int, start int) {

	if start >= len(a) {
		return
	}

	if a[start] == -1 {
		sortRecur(a, start+1)
		return
	}

	for i := start + 1; i < len(a); i++ {
		if a[start] > a[i] && a[i] > -1 {
			a[i], a[start] = a[start], a[i]
		}
	}

	sortRecur(a, start+1)
}
