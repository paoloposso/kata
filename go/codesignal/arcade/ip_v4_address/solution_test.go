package ipv4address

import "testing"

func TestShouldFailEmptyFirstTerm(t *testing.T) {
	res := solution(".254.255.0")

	if res == true {
		t.Fail()
	}
}
