package alphabeticshift

func solution(inputString string) string {
	result := ""

	for _, c := range inputString {
		s := c + 1
		if s > 122 {
			s = s - 26
		}
		result += string(s)
	}

	return result
}
