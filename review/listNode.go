package review

import (
	"sort"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummyHead *ListNode
	dummyHead.Next = head
	prev, cur := dummyHead, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || len(nums) < k {
		return 0
	}
	sort.Ints(nums)
	return nums[len(nums)-k-1]
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		for i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			tmp := []int{nums[i], nums[l], nums[r]}
			if sum == 0 {
				res = append(res, tmp)
				for r < len(nums)-1 && nums[r] == nums[r+1] {
					r--
				}
				for l > i+1 && nums[l] == nums[l-1] {
					l++
				}
				r--
				l++
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}
	var res [][]int
	sort.Ints(nums)
	dfsPermuteDemo(nums, &res, []int{})
	return res
}
func dfsPermuteDemo(nums []int, res *[][]int, cur []int) {
	if len(nums) == len(cur) {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		dfsPermuteDemo(nums, res, cur)
		cur = cur[:len(cur)-1]
	}
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	cur, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > cur {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		if cur < 0 {
			cur = 0
		}
	}
	return sum
}
