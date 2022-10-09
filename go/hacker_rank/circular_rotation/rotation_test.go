package circularrotation_test

import (
	"circularrotation"
	"testing"
)

func TestQueryCirculatRotation(t *testing.T) {
	arr := []int32{1, 3, 9, 8, 7, 4}
	result := circularrotation.QueryArrayCircularRotation(arr, 2, []int32{0, 4, 2})
	if result[0] != 7 {
		t.Fail()
	}
	if result[1] != 9 {
		t.Fail()
	}
}
