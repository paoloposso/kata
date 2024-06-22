package reverseparentesis

func solution(inputString string) string {
	q := []string{}
	result := ""

	ix := 0

	for ix < len(inputString) {
		cur := inputString[ix]

		if cur == '(' {
			q = append(q, string(cur))
		} else if len(q) == 0 {
			result += string(cur)
		} else if cur == ')' {
			rev := reverseString(q[len(q)-1])
			q = q[:len(q)-1]
			if len(q) == 0 {
				result += string(rev)
			} else {
				q[len(q)-1] += rev
			}
		} else {
			q[len(q)-1] += string(cur)
		}
		ix++
	}

	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	runes = runes[1 : len(runes)-1]

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
