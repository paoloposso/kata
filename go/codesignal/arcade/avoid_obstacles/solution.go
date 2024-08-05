package avoidobstacles

func solution(inputArray []int) int {
	stepLen := 1

	for {
		if !hasMultipleOf(inputArray, stepLen) {
			return stepLen
		}
		stepLen++
	}
}

func hasMultipleOf(inputArray []int, item int) bool {
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i]%item == 0 {
			return true
		}
	}
	return false
}
