package removek

import (
	"testing"
)

func TestShouldRemoveAll3s(t *testing.T) {
	list := &ListNode{Value: 3}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	head := solution(list, 3)

	for head != nil {
		if head.Value == 3 {
			t.Fail()
		}
		head = head.Next
	}
}

func TestShouldRemoveAll3sEmpty(t *testing.T) {
	var list ListNode

	head := solution(&list, 3)

	for head != nil {
		if head.Value == 3 {
			t.Fail()
		}
		head = head.Next
	}
}
