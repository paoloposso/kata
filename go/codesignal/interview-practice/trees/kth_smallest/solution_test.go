package kthsmallest

import (
	"testing"
)

func Test1(t *testing.T) {
	if solution(&Tree{
		Value: 3,
		Left: &Tree{
			Value: 1,
		},
		Right: &Tree{
			Value: 5,
			Left: &Tree{
				Value: 4,
			},
			Right: &Tree{
				Value: 6,
			},
		},
	}, 4) != 5 {
		t.Fail()
	}
}
