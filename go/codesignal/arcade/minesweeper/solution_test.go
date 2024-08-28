package minesweeper

import "testing"

func TestShouldWorkForSquare(t *testing.T) {
	res := solution(
		[][]bool{
			{true, false, false},
			{false, true, false},
			{false, false, false}})

	if len(res) != 1 {
		t.Fail()
	}
}
