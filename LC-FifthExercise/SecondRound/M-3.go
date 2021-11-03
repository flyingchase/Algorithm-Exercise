package secondround

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func M3lengthOfUniqueueSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	left, right, res := 0, -1, 0
	for right+1 < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		if res < right-left+1 {
			res = right - left + 1
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func quicksortDemo(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	quicksortDemoHelper(nums, 0, len(nums)-1)
	return nums
}

func quicksortDemoHelper(nums []int, l, r int) {
	if l < r {
		p := paratitions(nums, l, r)
		quicksortDemoHelper(nums, l, p[0]-1)
		quicksortDemoHelper(nums, p[1]+1, r)
	}
}
func paratitions(nums []int, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if nums[l] < nums[r] {
			less++
			nums[less], nums[l] = nums[l], nums[less]
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
func zigzaTraversalBT(root *TreeNode) {
	fmt.Println("zhou")
}
