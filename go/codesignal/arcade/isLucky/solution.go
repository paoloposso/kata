package islucky

import (
	"strconv"
)

func solution(n int) bool {

	str := strconv.Itoa(n)

	firstHalf := 0

	for i := 0; i < len(str)/2; i++ {
		firstHalf += int(str[i] - '0')
	}

	secHalf := 0

	for i := len(str) / 2; i < len(str); i++ {
		secHalf += int(str[i] - '0')
	}

	return firstHalf == secHalf
}
