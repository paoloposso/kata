package linkedlist

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
