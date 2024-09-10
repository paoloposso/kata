package arrayreplace

func solution(inputArray []int, elemToReplace int, substitutionElem int) []int {
	result := []int{}

	for _, v := range inputArray {
		if v == elemToReplace {
			result = append(result, substitutionElem)
		} else {
			result = append(result, v)
		}
	}

	return result
}
