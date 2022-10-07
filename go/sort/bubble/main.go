package main

import "fmt"

func main() {
	fmt.Print(bubbleSort([]int{3, 2, 4, 8, 1, 7}))
}

func bubbleSort(array []int) []int {
	f := len(array)
	for ix := 0; ix <= f; ix++ {
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i] += array[i+1]
				array[i+1] = array[i] - array[i+1]
				array[i] = array[i] - array[i+1]
			}
		}
		f--
	}
	return array
}
