package main

import (
	"fmt"
	"linkedlist/list"
)

func main() {
	ll := list.NewLinkedList[string]()

	ll.InsertOnHead("1")
	ll.InsertOnHead("2")
	ll.InsertOnTail("3")

	ll.Print()
	fmt.Print("\n")

	ll.InsertOnTail("5")
	ll.InsertOnHead("7")

	ll.Print()
	fmt.Print("\n")

	val, err := ll.GetValueByIndex(4)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(*val)
	}
}
