package arcade

import "testing"

func TestShouldCompareStringCorrectly(t *testing.T) {
	if solution("aabcc", "adcaa") != 3 {
		t.Error("Result should be equals 3")
	}
}
