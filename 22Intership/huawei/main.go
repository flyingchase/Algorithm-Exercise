package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	nums := []int{}
	fmt.Scan(&n)
	for n > 0 {
		n--
		tmp := 0
		fmt.Scan(&tmp)
		nums = append(nums, tmp)
	}
	fmt.Println(randDelte(nums))
}
func qishuiping(n int, nums []int) []int {
	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, helper(nums[i]))
	}
	return res
}
func helper(num int) int {

	return num
}
func randDelte(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return nums
}
