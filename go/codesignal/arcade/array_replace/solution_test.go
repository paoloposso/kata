package arrayreplace

import "testing"

func TestShouldCompareStringCorrectly(t *testing.T) {
	res := solution([]int{1, 2, 1}, 1, 3)

	if len(res) != 3 {
		t.Fail()
	}

	if res[0] != 3 {
		t.Fail()
	}

	if res[1] != 2 {
		t.Fail()
	}
}
