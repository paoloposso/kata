package islucky

import "testing"

func TestShouldTestNumberCorrectly(t *testing.T) {
	if solution(434433) {
		t.Fail()
	}
}
