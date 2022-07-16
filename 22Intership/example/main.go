package main

import (
	"fmt"
)

func main() {
	fmt.Println("debug")
	// 1,2,3,4,5,6
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//}
	//
	//fmt.Println(intership.PreOrderTrvaersalBT(root))
	//
	//matrix := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 0, 0, 0}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	//fmt.Println(len(matrix[0]))

	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	nums := strings.Split(input.Text(), " ")
	// 	res := 0
	// 	for i := 1; i < len(nums); i++ {
	// 		num, _ := strconv.Atoi(nums[i])
	// 		res += num
	// 	}
	// 	fmt.Println(res)
	// }

	// fmt.Println(interview.MinArray([]int{3, 4, 5, 1, 2}))
	// fmt.Println(codetop.LengthOfLongestSubstring("abcabcbb"))
	// fmt.Println(interview.TranslateNum(624))
	fmt.Println(interview.LengthOfLongestSubstring("abcabcbb"))
}
func quicksortDemo(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	quickersortHelper(nums, 0, len(nums)-1)
	return nums
}
func quickersortHelper(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := paratition(nums, l, r)
	quickersortHelper(nums, l, p[0]-1)
	quickersortHelper(nums, p[1]+1, r)
}
func paratition(nums []int, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if nums[l] < nums[r] {
			less++
			nums[less], nums[l] = nums[l], nums[less]
			l++
		} else if nums[l] > nums[r] {
			more--
			nums[more], nums[l] = nums[l], nums[more]
		} else {
			l++
		}
	}
	nums[more], nums[r] = nums[r], nums[more]
	return []int{less + 1, more}
}

func mergesort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	mergesortHelp(nums, 0, len(nums)-1)
	return nums

}
func mergesortHelp(nums []int, l, r int) {
	if l > r {
		return
	}

	if len(nums) == 0 {
		return
	}

}

func mergesorthelper2(nums []int) {
	if len(nums) == 0 {
		return
	}
}
