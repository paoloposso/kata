package removek

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

func solution(l *ListNode, k int) *ListNode {
	for l != nil && l.Value == k {
		l = l.Next
	}
	head := l
	rem(l, k)
	return head
}

func rem(l *ListNode, k int) {
	for l != nil && l.Value == k {
		l = l.Next
	}
	for l != nil {
		for l.Next != nil && l.Next.Value == k {
			l.Next = l.Next.Next
		}
		l = l.Next
	}
}
