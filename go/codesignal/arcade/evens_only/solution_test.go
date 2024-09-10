package evensonly

import "testing"

func Test1(t *testing.T) {
	if solution(12345) {
		t.Fail()
	}
	if !solution(24648) {
		t.Fail()
	}
}
