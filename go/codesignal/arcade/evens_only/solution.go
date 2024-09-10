package evensonly

import "fmt"

func solution(n int) bool {
	for _, v := range fmt.Sprint(n) {
		if int(v)%2 != 0 {
			return false
		}
	}

	return true
}
