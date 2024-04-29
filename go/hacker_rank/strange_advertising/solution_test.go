package strangeadvertising

import "testing"

func TestInput1(t *testing.T) {
	b := viralAdvertising(5)
	if b != 24 {
		t.Fail()
	}
}
