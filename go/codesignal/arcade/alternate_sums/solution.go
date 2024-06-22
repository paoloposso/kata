package alternatesums

func solution(a []int) []int {
	resultA, resultB := 0, 0

	for i, v := range a {
		if (i+1)%2 == 0 {
			resultA += v
		} else {
			resultB += v
		}
	}

	return []int{resultB, resultA}
}
