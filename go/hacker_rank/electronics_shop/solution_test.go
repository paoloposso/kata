package electronicsshop_test

import (
	electronicsshop "hacker_rank/electronics_shop"
	"testing"
)

func TestInput1(t *testing.T) {
	spent := electronicsshop.Solve([]int32{3, 1}, []int32{5, 2, 8}, 10)
	if spent != 9 {
		t.Fail()
	}
}

func TestInput2(t *testing.T) {
	spent := electronicsshop.Solve([]int32{40, 50, 60}, []int32{5, 8, 12}, 60)
	if spent != 58 {
		t.Fail()
	}
}
