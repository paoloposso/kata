package reverseparentheses

import "testing"

func TestShoudInvert1(t *testing.T) {
	if solution("(bar)") != "rab" {
		t.Fail()
	}
}

func TestShoudInvert2(t *testing.T) {
	if solution("foo(bar(baz))blim") != "foobazrabblim" {
		t.Fail()
	}
}
