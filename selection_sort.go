package main

import (
	"fmt"
)

// SelectionSort function sorts an array using Selection Sort algorithm
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Assume the first element of the unsorted part is the smallest
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)

	SelectionSort(arr)

	fmt.Println("Sorted array:", arr)
}
