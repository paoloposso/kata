package haspathgivensum

import (
	"testing"
)

func Test1(t *testing.T) {
	if !solution(&Tree{
		Value: 2,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 4,
				},
				Right: &Tree{
					Value: 5,
				},
			},
			Right: &Tree{
				Value: 6,
				Left: &Tree{
					Value: 7,
				},
				Right: &Tree{
					Value: 8,
				},
			},
		},
	}, 12) {
		t.Error()
	}
}

func Test2(t *testing.T) {
	if solution(&Tree{
		Value: 8,
		Right: &Tree{
			Value: 2,
			Right: &Tree{
				Value: 3,
			},
		},
	}, 8) {
		t.Error()
	}
}
