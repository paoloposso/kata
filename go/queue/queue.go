package queue

import "fmt"

type Queue[T any] struct {
	tail *Item[T]
	head *Item[T]
}

type Item[T any] struct {
	Data T
	next *Item[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{tail: nil, head: nil}
}

func (queue *Queue[T]) Push(data T) {
	newItem := &Item[T]{
		Data: data,
		next: nil,
	}
	if queue.tail != nil {
		queue.tail.next = newItem
	}
	queue.tail = newItem
	if queue.head == nil {
		queue.head = newItem
	}
}

func (queue *Queue[T]) Pop() *Item[T] {
	result := queue.head
	if queue.head != nil {
		queue.head = queue.head.next
	}
	return result
}

func (queue Queue[T]) PrintQueue() {
	item := queue.head
	for item != nil {
		fmt.Printf("%v\n", item.Data)
		item = item.next
	}
}
