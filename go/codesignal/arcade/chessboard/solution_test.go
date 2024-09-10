package chessboard

import "testing"

func Test1(t *testing.T) {
	if !solution("A1", "C3") {
		t.Fail()
	}
	if solution("A1", "C4") {
		t.Fail()
	}
}
