package blur

import "testing"

func TestShouldWorkForSquare(t *testing.T) {
	res := solution(
		[][]int{
			{1, 1, 1},
			{1, 7, 1},
			{1, 1, 1}})

	if len(res) != 1 {
		t.Fail()
	}
	if res[0][0] != 1 {
		t.Fail()
	}
}

func TestShouldWorkForFourSquares(t *testing.T) {
	res := solution(
		[][]int{
			{7, 4, 0, 1},
			{5, 6, 2, 2},
			{6, 10, 7, 8},
			{1, 4, 2, 0},
		})

	if len(res) != 2 {
		t.Fail()
	}
	if res[0][0] != 5 {
		t.Fail()
	}
	if res[1][1] != 4 {
		t.Fail()
	}
}
