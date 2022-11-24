package main

type LinkedList struct {
	head *Node
}

type Node struct {
	value    string
	next     *Node
	previous *Node
}

func (l *LinkedList) append(value string) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = newNode
	tail.next.previous = tail
}

func (l *LinkedList) prepend(value string) {
	newNode := &Node{value: value}
	if l.head != nil {
		newNode.next = l.head
		l.head.previous = newNode
	}
	l.head = newNode
}

func main() {
	ll := LinkedList{}
	ll.append("aaaa")
	ll.append("bbb")
	ll.append("bbbssssdsd")
	ll.prepend("ffff")
}
