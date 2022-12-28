package main

import "fmt"

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	// sort it
	for i := 0; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			// swap
			arr[idx] = arr[idx] + arr[i]
			arr[i] = arr[idx] - arr[i]
			arr[idx] = arr[idx] - arr[i]
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx //pivot index
}

func main() {
	arr := []int{4, 5, 8, 10, 16, 1, 7}

	quickSort(arr)

	fmt.Println(arr)
}
