package main

import (
	"fmt"
	ankiwrite "review/ankiWrite"
)

func main() {
	nums := []int{1, 0, 9, 2, 6, 3, 8, 4, 5, 11, 7}
	fmt.Println(ankiwrite.QuickSort(nums))
}
