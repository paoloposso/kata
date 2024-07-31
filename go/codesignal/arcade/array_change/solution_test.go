package arraychange

import "testing"

func Test1(t *testing.T) {
	if solution([]int{1, 1, 1}) == 3 {
		return
	}
	t.Fail()
}
