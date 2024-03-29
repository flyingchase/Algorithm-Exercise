package main

import (
	"fmt"
)

func main() {
	// var a int
	// fmt.Scan(&a)
	nums := []int{1, 3, 5, 2, 4, 6, 8, 0}
	res := longestIncreasingSubArray(nums)

	fmt.Println(res)
}
func longestIncreasingSubArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}
	res := []int{}
	for i := 0; i < len(nums)-1; i++ {
		temp := []int{nums[i]}
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] > nums[i] {
				temp = append(temp, nums[j])
			} else {
				break
			}
		}
		i = j
		if len(res) < len(temp) {
			res = append([]int{}, temp...)
		}
	}
	return res
}
