package addborder

import "testing"

func TestShouldCompareStringCorrectly(t *testing.T) {
	res := solution([]string{"abc", "def"})

	if len(res) != 4 {
		t.Fail()
	}
	if len(res[0]) != 5 {
		t.Fail()
	}
}
