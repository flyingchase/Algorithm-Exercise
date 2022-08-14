package main

import (
	"fmt"
	"reflect"
)

func main() {
	// nums := []int{1, 5, 8, 7, 3, 0, 9, 2, 6, 4, 10, 12, 11}
	// fmt.Printf("findfeeling.Search(nums, 5): %v\n", findfeeling.Search(nums, 5))
	// fmt.Printf("sort.SearchInts(nums, 5): %v\n", sort.SearchInts(nums, 5))
	// fmt.Printf("findfeeling.Quicksort(nums): %v\n", findfeeling.Quicksort(nums))
	// fmt.Printf("findfeeling.Heapsort(nums): %v\n", findfeeling.Heapsort(nums))
	// fmt.Printf("findfeeling.MergeSort(nums): %v\n", findfeeling.MergeSort(nums))

	// candidates := []int{10, 1, 2, 7, 6, 8, 2, 1, 5, 10}
	// target:=8
	// fmt.Println(findfeeling.CombinationSum2(candidates, target))

	// fmt.Println(findfeeling.Lis2(candidates))

	// num1 := "123"
	// num2 := "456"
	// fmt.Println(findfeeling.Multiply(num1, num2))

	runeDemo := rune('我')
	fmt.Println(reflect.TypeOf(runeDemo))

}
