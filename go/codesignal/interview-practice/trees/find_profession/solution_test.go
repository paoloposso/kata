package findprofession

import (
	"testing"
)

func Test1(t *testing.T) {
	if solution(3, 3) != "Doctor" {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if solution(3, 2) != "Doctor" {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if solution(3, 1) != "Engineer" {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if solution(1, 1) != "Engineer" {
		t.Fail()
	}
}
