package alphabeticshift

import "testing"

func Test1(t *testing.T) {
	if solution("crazy") != "dsbaz" {
		t.Fail()
	}
}
