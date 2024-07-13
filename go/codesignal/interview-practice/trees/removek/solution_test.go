package removek

import (
	"testing"
)

func TestTraverse(t *testing.T) {
	tree := &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Right: &Tree{
				Value: 3,
			},
		},
		Right: &Tree{
			Value: 4,
			Left: &Tree{
				Value: 5,
			},
		},
	}

	solution(tree)
}

func TestTraverse2(t *testing.T) {
	tree := &Tree{
		Value: 1000,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 10,
				Left:  nil,
				Right: &Tree{
					Value: -1000,
					Left:  nil,
					Right: &Tree{
						Value: 0,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &Tree{
				Value: 4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Tree{
			Value: 3,
			Left:  nil,
			Right: &Tree{
				Value: 99,
				Left: &Tree{
					Value: 6,
					Left: &Tree{
						Value: 1,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	solution(tree)
}
