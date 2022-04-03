package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// l1 := &ListNode{1, &ListNode{2, nil}}
	// // l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// // days.SortList(head)
	// // days.ReorderList(head)
	// // for head != nil {
	// // 	fmt.Println(head.Val)
	// // 	head = head.Next
	// // }
	// fmt.Println(isPalindrome(l1))
	//
	// nums := []int{1, 2, 3}
	// fmt.Println(subarraySum(nums, 3))
	fmt.Println(searchRange([]int{8}, 8))
}

func searchRange(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			i, j := mid, mid
			for i >= -1 && j <= len(nums) && (i >= 0 && nums[i] == target || j < len(nums) && nums[j] == target) {
				if i >= 0 && nums[i] == target {
					i--
				}
				if j < len(nums) && nums[j] == target {
					j++
				}
				if i >= 0 && j < len(nums) && nums[i] != target && nums[j] != target {
					break
				}
			}
			return []int{i + 1, j - 1}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return []int{-1, -1}
}

// 逆转链表，使用栈
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	mid := count / 2
	cur = head
	stack := []*ListNode{}
	for mid > 0 {
		mid--
		stack = append(stack, cur)
		cur = cur.Next
	}
	if count%2 == 1 {
		cur = cur.Next
	}
	for cur != nil {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val != cur.Val {
			return false
		}
		cur = cur.Next
	}
	return true
}

// 单向遍历不断逆转
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev, slow = slow, next
	}
	// odd
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow, prev = slow.Next, prev.Next
	}
	return true
}

// 前缀和
func subarraySum(nums []int, k int) int {
	count, hash := 0, map[int]int{0: 1}
	// 前 i 项的总和 Sum从 i——j 之和=preSum[j]-preSum[i-1]
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		hash[preSum]++
	}
	return count
}
