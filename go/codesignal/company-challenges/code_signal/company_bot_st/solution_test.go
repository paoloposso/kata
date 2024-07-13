package companybotst

import (
	"testing"
)

func Test1(t *testing.T) {
	data := [][]int{
		{4, 1},
		{4, -1},
		{0, 0},
		{6, 1},
	}

	if solution(data) != 5.0 {
		t.Fail()
	}
}
