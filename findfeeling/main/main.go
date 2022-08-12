package main

import (
	"findfeeling"
	"fmt"
)

func main() {
	nums := []int{1, 5, 8, 7, 3, 0, 9, 2, 6, 4, 10, 12, 11}
	// fmt.Printf("findfeeling.Search(nums, 5): %v\n", findfeeling.Search(nums, 5))
	// fmt.Printf("sort.SearchInts(nums, 5): %v\n", sort.SearchInts(nums, 5))
	fmt.Printf("findfeeling.Quicksort(nums): %v\n", findfeeling.Quicksort(nums))
	// fmt.Printf("findfeeling.Heapsort(nums): %v\n", findfeeling.Heapsort(nums))
	// fmt.Printf("findfeeling.MergeSort(nums): %v\n", findfeeling.MergeSort(nums))
}
