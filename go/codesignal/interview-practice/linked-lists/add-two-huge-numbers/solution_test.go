package addtwohugenumbers

import (
	"testing"
)

func TestShouldTestNumberCorrectly(t *testing.T) {
	// list := &ListNode{Value: 123}
	// list.Add(4)
	// list.Add(5)

	// list2 := &ListNode{Value: 100}
	// list2.Add(100)
	// list2.Add(100)

	list := &ListNode{Value: 9876}
	list.Add(5432)
	list.Add(0)

	list2 := &ListNode{Value: 1}
	list2.Add(8001)

	solution(list, list2)
}
