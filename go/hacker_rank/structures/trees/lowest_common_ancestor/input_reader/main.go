package main

import (
	"strconv"
	"strings"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// _, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	os.Exit(1)
	// }

	// input2, err := reader.ReadString('\n')
	// if err != nil {
	// 	os.Exit(1)
	// }

	// input3, err := reader.ReadString('\n')
	// if err != nil {
	// 	os.Exit(1)
	// }

	input2 := "4 2 3 1 7 6"
	input3 := "1 7"

	arrInput2 := strings.Split(input2, " ")
	arrInput3 := strings.Split(input3, " ")

	i, _ := strconv.Atoi(arrInput2[0])
	v1, _ := strconv.Atoi(arrInput3[0])
	v2, _ := strconv.Atoi(arrInput3[1])

	head := Insert(nil, i)

	for _, v := range arrInput2[1:] {
		i, _ = strconv.Atoi(v)
		Insert(head, i)
	}

	Solve(head, v1, v2)
}

func Solve(root *Node, v1, v2 int) {

}

type Node struct {
	Right *Node
	Left  *Node
	Value int
}

func Insert(head *Node, value int) *Node {
	if head == nil {
		head = &Node{Value: value}
		return head
	}

	if value >= head.Value {
		head.Right = Insert(head.Right, value)
	} else {
		head.Left = Insert(head.Left, value)
	}

	return head
}
