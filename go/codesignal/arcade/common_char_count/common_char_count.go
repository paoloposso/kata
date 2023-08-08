package arcade

func solution(s1 string, s2 string) int {
	map1 := populateMap(s1)
	map2 := populateMap(s2)

	result := checkQtdInCommon(map1, map2)
	result += checkQtdInCommon(map2, map1)

	return result
}

func checkQtdInCommon(map1, map2 map[string]int) int {
	result := 0

	for cur, qtd1 := range map1 {
		qtd2, exists := map2[cur]

		if exists {
			result += min(qtd1, qtd2)
		}

		delete(map1, cur)
	}

	return result
}

func populateMap(str string) map[string]int {
	result := make(map[string]int)

	for _, charR := range str {
		char := string(charR)
		result[char]++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
