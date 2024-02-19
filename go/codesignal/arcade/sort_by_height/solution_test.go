package islucky

import "testing"

func TestShouldTestNumberCorrectly(t *testing.T) {
	if solution([]int{-1, 150, 190, 170, -1, -1, 160, 180})[0] != -1 {
		t.Fail()
	}
}
