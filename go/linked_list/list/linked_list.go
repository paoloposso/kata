package list

import (
	"errors"
	"fmt"
)

type LinkedList[T any] struct {
	head *Node[T]
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (linkedList *LinkedList[T]) InsertOnHead(value T) {
	newNode := &Node[T]{value: value, next: nil}
	newNode.next = linkedList.head
	linkedList.head = newNode
}

func (linkedList *LinkedList[T]) InsertOnTail(value T) {
	newNode := &Node[T]{value: value, next: nil}
	if linkedList.head == nil {
		linkedList.head = newNode
		return
	}
	tail := linkedList.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = newNode
}

func (linkedList *LinkedList[T]) Print() {
	cur := linkedList.head
	for cur != nil {
		fmt.Printf("->%v", cur.value)
		cur = cur.next
	}
}

func (linkedList LinkedList[T]) GetValueByIndex(index int) (*T, error) {
	result := linkedList.head
	for i := 0; i < index; i++ {
		result = result.next
		if result == nil {
			return nil, errors.New("not found")
		}
	}
	return &result.value, nil
}
