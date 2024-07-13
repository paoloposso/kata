package symmetric

import (
	"testing"
)

func Test1(t *testing.T) {
	tree := &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 3,
			},
			Right: &Tree{
				Value: 4,
			},
		},
		Right: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 4,
			},
			Right: &Tree{
				Value: 3,
			},
		},
	}

	if !solution(tree) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	tree := &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 3,
				},
				Right: &Tree{
					Value: 2,
				},
			},
			Right: &Tree{
				Value: 4,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 3,
				},
			},
		},
		Right: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 4,
				Left: &Tree{
					Value: 3,
				},
				Right: &Tree{
					Value: 2,
				},
			},
			Right: &Tree{
				Value: 3,
				Left: &Tree{
					Value: 2,
				},
				Right: &Tree{
					Value: 3,
				},
			},
		},
	}

	if !solution(tree) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	tree := &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Right: &Tree{
				Value: 3,
			},
		},
		Right: &Tree{
			Value: 2,
			Right: &Tree{
				Value: 3,
			},
		},
	}

	if solution(tree) {
		t.Fail()
	}
}
