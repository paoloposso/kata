package arraychange

func solution(inputArray []int) int {
	moves := 0

	i := 0

	for i < len(inputArray)-1 {
		for inputArray[i] >= inputArray[i+1] {
			inputArray[i+1]++
			moves++
		}
		i++
	}

	return moves
}
