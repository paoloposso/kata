package jumpingclouds

import "testing"

func TestInput1(t *testing.T) {
	energy := jumpingOnClouds([]int32{0, 0, 1, 0}, 2)
	if energy != 96 {
		t.Fail()
	}
}

func TestInput2(t *testing.T) {
	energy := jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2)
	if energy != 92 {
		t.Fail()
	}
}
