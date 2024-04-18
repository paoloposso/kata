package equalizearrays

import "testing"

func TestInput1(t *testing.T) {
	removed := equalizeArray([]int32{3, 2, 1, 2})
	if removed != 2 {
		t.Fail()
	}
}
