package main

import (
	"errors"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value    string
	next     *Node
	previous *Node
}

func (l *LinkedList) append(value string) {
	l.length++
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	tail := l.tail
	tail.next = newNode
	tail.next.previous = tail
	l.tail = newNode
}

func (l *LinkedList) prepend(value string) {
	l.length++
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head.previous = newNode
	l.head = newNode
}

func (l *LinkedList) insertAt(value string, index int) error {
	if l.length < index {
		return errors.New("index must be within list")
	}
	l.length++
	if index == 0 {
		l.prepend(value)
		return nil
	}
	if index == l.length-1 {
		l.append(value)
		return nil
	}
	newNode := &Node{value: value}
	nex := l.head
	for i := 0; i < index; i++ {
		nex = nex.next
	}
	newNode.next = nex
	nex.previous.next = newNode
	newNode.previous = nex.previous
	nex.previous = newNode
	return nil
}

func (l *LinkedList) removeAt(index int) error {

	if l.length < index {
		return errors.New("index must be within list")
	}

	l.length--

	cur := l.head

	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next.previous = cur.previous
	cur.previous.next = cur.next
	cur.next = nil
	cur.previous = nil

	return nil
}

func (l *LinkedList) printList() {
	cur := l.head
	for cur != nil {
		print(cur.value + "<->")
		cur = cur.next
	}
}

func main() {
	ll := LinkedList{}
	ll.prepend("11111")
	ll.append("aaaa")
	ll.append("bbb")
	ll.append("bbbssssdsd")
	ll.prepend("ffff")
	ll.prepend("ffxxx")
	ll.append("ffxxx")
	ll.insertAt("insertedAt3", 3)
	ll.insertAt("insertedAt8", 100)
	ll.printList()
	println("")
	ll.removeAt(1)
	ll.printList()
	println("")
	println(ll.length)
}
