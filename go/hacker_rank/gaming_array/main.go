package main

func gamingArray(arr []int32) string {

	turn := 0

	for len(arr) > 0 {
		turn++
		indexToRemove := 0
		valueToRemove := arr[0]

		for i, v := range arr {
			if v > valueToRemove {
				indexToRemove = i
				valueToRemove = v
			}
		}

		arr = arr[:indexToRemove]
	}

	result := ""

	if turn%2 == 0 {
		result = "BOB"
	} else {
		result = "ANDY"
	}

	return result
}

func main() {
	gamingArray([]int32{2, 5, 6, 3, 4})
}

func isInArray(target int32, arr []int32) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
