package ispalindrome

import (
	"linkedlist"
)

func solution(l *linkedlist.ListNode) bool {
	if l == nil {
		return true
	}

	f := l
	s := l
	finished := false
	halfLength := 0

	for !finished {
		if f.Next != nil && f.Next.Next != nil {
			f = f.Next.Next
		} else if f.Next != nil {
			f = f.Next
		} else {
			finished = true
		}

		if !finished {
			s = s.Next
			halfLength++
		}
	}

	revList := revert(s)

	for revList != nil && l != nil {
		if revList.Value != l.Value {
			return false
		}
		revList = revList.Next
		l = l.Next
	}

	return true
}

func revert(l *linkedlist.ListNode) *linkedlist.ListNode {
	cur := l
	var next, prev *linkedlist.ListNode = nil, nil

	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}
