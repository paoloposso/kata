package kthsmallest

import (
	"sort"
)

type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func solution(t *Tree, k int) int {
	if t == nil {
		return 0
	}

	arr := []int{}

	arr = recur(t, arr, k)

	sort.Ints(arr)

	return arr[k-1]
}

func recur(t *Tree, arr []int, k int) []int {
	if t == nil {
		return arr
	}
	if t.Left == nil {
		arr = append(arr, t.Value.(int))
		// arr = recur(t.Left, arr, k)
		arr = recur(t.Right, arr, k)
	} else {
		arr = recur(t.Left, arr, k)
		arr = append(arr, t.Value.(int))
		arr = recur(t.Right, arr, k)
	}

	return arr
}
