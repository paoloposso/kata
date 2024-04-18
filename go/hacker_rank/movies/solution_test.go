package movies

import "testing"

func TestInput1(t *testing.T) {
	b := beautifulDays(20, 23, 6)
	if b != 2 {
		t.Fail()
	}
}
