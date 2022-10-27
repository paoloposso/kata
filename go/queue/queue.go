package queue

type Queue[T any] struct {
	tail *Item[T]
	head *Item[T]
}

type Item[T any] struct {
	data T
	next *Item[T]
}

func (queue *Queue[T]) Push(data T) {
	newItem := &Item[T]{
		data: data,
		next: nil,
	}
	if queue.tail != nil {
		queue.tail.next = newItem
	}
	queue.tail = newItem
}

func (queue *Queue[T]) Pop() *Item[T] {
	result := queue.head
	if queue.head != nil {
		queue.head = queue.head.next
	}
	return result
}
