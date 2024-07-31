package arraysaresimilar

import "testing"

func Test1(t *testing.T) {
	if solution([]int{1, 2, 3}, []int{1, 2, 3}) == false {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if solution([]int{1, 2, 3}, []int{2, 1, 3}) == false {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if solution([]int{2, 2, 3}, []int{2, 1, 3}) == true {
		t.Fail()
	}
}
