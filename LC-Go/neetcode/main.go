package main

import (
	"fmt"
	"neetcode/tricks"
)

func main() {
	nums := []int{0, -10, 1, -1, 5, 9, 2, 4, 8, 6, 3, 7, -3}
	// tricks.Quicksort(nums, 0, len(nums)-1)
	// tricks.Mergesort(nums, 0, len(nums)-1)
	tricks.Heapsort(nums)
	fmt.Println(nums)

}
