package staircase

import "testing"

func TestShouldTestNumberCorrectly(t *testing.T) {
	r := solution(4, 2)

	if len(r) < 3 {
		t.Fail()
	}
}
