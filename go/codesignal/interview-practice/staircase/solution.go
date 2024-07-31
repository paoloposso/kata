package staircase

func solution(n, k int) [][]int {
	var result [][]int
	var backtrack func(currentSum int, path []int)
	backtrack = func(currentSum int, path []int) {
		if currentSum == n {
			result = append(result, append([]int(nil), path...))
			return
		}
		if currentSum > n {
			return
		}
		for step := 1; step <= k; step++ {
			path = append(path, step)
			backtrack(currentSum+step, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, []int{})
	return result
}
