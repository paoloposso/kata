package main

import "fmt"

func main() {
	sorted := bubbleSort([]int{2, 3, 4, 1, 7, 9})
	fmt.Print(sorted)
}

func bubbleSort(array []int) []int {
	for {
		changed := false
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return array
}
