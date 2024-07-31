package ipv4address

import (
	"strconv"
	"strings"
)

func solution(inputString string) bool {
	sp := strings.Split(inputString, ".")

	if len(sp) != 4 {
		return false
	}

	for _, s := range sp {
		if containsLeadingZeroes(s) {
			return false
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if i > 255 {
			return false
		}
	}

	return true
}

func containsLeadingZeroes(s string) bool {
	return len(s) > 1 && s[0] == '0'
}
