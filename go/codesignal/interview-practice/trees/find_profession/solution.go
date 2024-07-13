package findprofession

import (
	"strconv"
)

func solution(level int, pos int) string {
	b := convertIntToBinary(pos-1, level+1)
	p := "E"

	for i, v := range b {
		if i == 0 {
			continue
		}
		if v == '0' {
			if p == "E" {
				p = "E"
			} else {
				p = "D"
			}
		} else {
			if p == "E" {
				p = "D"
			} else {
				p = "E"
			}
		}
	}

	if p == "E" {
		return "Engineer"
	}
	return "Doctor"
}

func convertIntToBinary(num, levels int) string {
	binaryStr := strconv.FormatInt(int64(num), 2)

	for len(binaryStr) < levels {
		binaryStr = "0" + binaryStr
	}

	return binaryStr
}
