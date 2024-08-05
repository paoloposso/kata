package avoidobstacles

import "testing"

func Test1(t *testing.T) {
	if solution([]int{5, 3, 6, 7, 9}) != 4 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if solution([]int{2, 3}) != 4 {
		t.Fail()
	}
}
