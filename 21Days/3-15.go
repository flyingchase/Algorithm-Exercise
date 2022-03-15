package days

import "sort"

// k 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return head
	}
	var count int
	cur := head
	for cur != nil {
		cur = cur.Next
		count++
	}
	// 注意边界处理
	// i 从 1 开始，i*k<=length 同时每次增长 k
	for i := 1; i*k <= count; i++ {
		// 起始 l 左侧为(i-1)*k+1，为上一次右侧的下标加一
		head = reverseBetween(head, (i-1)*k+1, i*k)
	}
	return head
}

// 在 l 和 r 区间内翻转
func reverseBetween(head *ListNode, l, r int) *ListNode {
	if head == nil || l > r {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev := dummyHead
	// 找到要插入区间在整条链表中的起始位置的 prev
	for i := 0; i < l-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	next := cur.Next
	// 头插后逐个右移
	for i := l; i < r; i++ {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return dummyHead.Next
}

// 三数之和
func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return nil
	}
	// 排序之后方便去重
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// 避免重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[r] + nums[l] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[r], nums[l], nums[i]})
				// 优化，避免重复
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
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

// 912排序数组
// 已在 3-14 中 215 题写出三种 nlogn 排序实现
func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
