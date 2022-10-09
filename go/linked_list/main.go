package main

import (
	"fmt"
	"linkedlist/list"
)

func main() {
	ll := list.NewLinkedList[string]()

	ll.Prepend("1")
	ll.Prepend("2")
	ll.Append("3")
	ll.InsertAt("200", -1)

	ll.Print()
	fmt.Println()

	val, err := ll.Get(2)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s", *val)
		fmt.Println()
	}
}
