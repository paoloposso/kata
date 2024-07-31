package palindromerearraging

func solution(inputString string) bool {
	m := make(map[rune]int)

	for _, c := range inputString {
		m[c]++
	}

	countOdds := 0

	for _, v := range m {
		if v%2 != 0 {
			countOdds++
		}
		if countOdds > 1 {
			return false
		}
	}

	return true
}
