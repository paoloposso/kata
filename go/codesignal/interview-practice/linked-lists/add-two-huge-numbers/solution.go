package addtwohugenumbers

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// Add is a method that adds a node to the end of the list
func (head *ListNode) Add(value interface{}) *ListNode {
	cur := head

	if head == nil {
		return &ListNode{Value: value}
	}

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = &ListNode{Value: value}

	return cur.Next
}

func solution(a *ListNode, b *ListNode) *ListNode {
	cur := a
	aNumber := ""
	for cur != nil {
		strWithLeadingZeroes := fmt.Sprintf("%04d", cur.Value)
		aNumber += strWithLeadingZeroes
		cur = cur.Next
	}
	cur = b
	bnumber := ""
	for cur != nil {
		strWithLeadingZeroes := fmt.Sprintf("%04d", cur.Value)
		bnumber += strWithLeadingZeroes
		cur = cur.Next
	}

	resultA, _ := strconv.Atoi(aNumber)
	resultB, _ := strconv.Atoi(bnumber)

	result := &ListNode{}
	head := result

	total := fmt.Sprint(resultA + resultB)

	l := len(total) % 4
	if l == 0 {
		l = 4
	}

	fmt.Println(total[0:l])
	result.Value = total[0:l]

	i := l
	for i < len(total)-3 {
		result.Next = &ListNode{Value: total[i : i+4]}
		fmt.Println(total[i : i+4])
		i += 4
		result = result.Next
	}

	return head
}
