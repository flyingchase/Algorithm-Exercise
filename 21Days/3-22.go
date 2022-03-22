package days

import "sort"

// 删除链表倒数第 k 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, fast := dummyHead, head
	slow := head
	for fast != nil {
		for n > 0 {
			n--
			fast = fast.Next
		}
		if fast != nil {
			prev = slow
			fast, slow = fast.Next, slow.Next
		}
	}
	return dummyHead.Next
}

// 区间合并，将重叠区间合并返回非重叠区间
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	type node struct {
		start int
		end   int
	}
	nodes := []node{}
	for _, nums := range intervals {
		nodes = append(nodes, node{
			start: nums[0],
			end:   nums[1],
		})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].start <= nodes[j].start {
			return true
		}
		return nodes[i].end <= nodes[j].end
	})
	queue := []node{}
	queue = append(queue, nodes[0])
	curIndex := 0
	for i := 1; i < len(nodes); i++ {
		// 逐个与已排序的数组 start 和 end 比较，起点小于待比较的 nodes 的终点则追
		// 起点大于待比较的终点则取两者终点最大值
		if nodes[i].start > queue[curIndex].end {
			curIndex++
			queue = append(queue, nodes[i])
		} else {
			if nodes[i].end > queue[curIndex].end {
				queue[curIndex].end = nodes[i].end
			}
		}
	}
	res := [][]int{}
	for i := 0; i < len(queue); i++ {
		tmp := []int{}
		tmp = append(tmp, queue[i].start, queue[i].end)
		res = append(res, tmp)
	}
	return res
}
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	low, high := 0, x
	var mid, sqr int
	for {
		mid = low + (high-low)>>1
		if low == mid {
			return mid
		}
		sqr = mid * mid
		if sqr == x {
			return mid
		}
		if sqr > x {
			high = mid
		}
		if sqr < x {
			low = mid
		}
	}
}
