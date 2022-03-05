package intership

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
)

//  修改字符出现最多的次数的 num
func maxFrequency1838(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	countMap, res := make(map[int]int, 0), 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		count := countMap[cur]
		if i > 0 {
			operationCount := k
			for j := i - 1; j >= 0; j-- {
				tmp := nums[j]
				diff := cur - tmp
				if operationCount >= diff {
					add := operationCount / diff
					min := math.Min(float64(countMap[tmp]), float64(add))
					operationCount -= int(min) * (diff)
					count += int(min)
				} else {
					break
				}
			}
		}
		if res <= count {
			res = count
		}
	}
	return res
}

func UpdateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return nil
	}
	m, n := len(mat), len(mat[0])
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	dirX, dirY := []int{0, 0, 1, -1}, []int{1, -1, 0, 0}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		x, y := node[0], node[1]
		for i := 0; i < 4; i++ {
			newX, newY := x+dirX[i], y+dirY[i]
			if mat[newX][newY] == -1 && newX >= 0 && newX < m && newY >= 0 && newY < n {
				mat[newX][newY] = mat[x][y] + 1
				queue = append(queue, []int{newX, newY})
			}
		}
	}

	return mat
}

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
				bfsSearchNumsIsLands(grid, i, j)
			}
		}
	}
	return res
}
func bfsSearchNumsIsLands(grid [][]byte, row int, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}
	(grid)[row][col] = 0
	bfsSearchNumsIsLands(grid, row+1, col)
	bfsSearchNumsIsLands(grid, row-1, col)
	bfsSearchNumsIsLands(grid, row, col+1)
	bfsSearchNumsIsLands(grid, row, col-1)
}

func lca(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l, r := lca(root.Left, p, q), lca(root.Right, p, q)

	if l != nil {
		if r != nil {
			return root
		}
		return r
	}
	return l
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeKListsHelper(lists, 0, len(lists)-1)
}
func mergeKListsHelper(lists []*ListNode, l, r int) *ListNode {
	if l > r {
		return nil
	}
	if l == r {
		return lists[l]
	}
	mid := l + (r-l)>>1
	p := mergeKListsHelper(lists, l, mid)
	q := mergeKListsHelper(lists, mid+1, r)
	return mergeHelper(p, q)
}
func mergeHelper(p, q *ListNode) *ListNode {
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	dummyHead := new(ListNode)
	cur := dummyHead
	for p != nil && q != nil {
		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}
	return dummyHead.Next
}

type listHeap []*ListNode

func (h listHeap) Len() int {
	return len(h)
}
func (h listHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h listHeap) Less(i, j int) bool {
	return h[i].Val <= h[j].Val
}
func (h *listHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *listHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
func mergeKlistsWithHeap(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	h := &listHeap{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(h, lists[i])
	}
	dummyHead := &ListNode{Val: -1}
	cur := dummyHead
	for h.Len() != 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxLISassist(dp[j]+1, dp[i])
			}
			res = maxLISassist(dp[i], res)
		}
	}
	return res
}
func maxLISassist(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func fractionTODecimal(numerator int, denominator int) string {
	res := ""
	if numerator*denominator < 0 {
		res += "-"
		numerator = int(math.Abs(float64(numerator)))
		denominator = int(math.Abs(float64(denominator)))
	}
	res += strconv.Itoa(numerator / denominator)
	if numerator%denominator == 0 {
		return res
	}
	res += "."
	numerator %= denominator
	// 除数-在 res 中的位置
	data := make(map[int]int)
	data[numerator] = len(res)
	for numerator != 0 {
		numerator *= 10
		res += strconv.Itoa(numerator / denominator)
		numerator %= denominator
		if index, ok := data[numerator]; ok {
			res = res[:index] + "(" + res[index:] + ")"
			break
		} else {
			data[numerator] = len(res)
		}
	}
	return res
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	cur := root
	flag := false
	queue = append(queue, cur)
	res := [][]int{}
	for len(queue) != 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if flag {
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
		flag = !flag
		if tmp != nil {
			res = append(res, tmp)
		}
	}
	return res
}

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	res := make([]byte, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			num := ((num1[i] - '0') * (num2[j] - '0')) + res[i+j+1]
			res[i+j], res[i+j+1] = res[i+j]+num/10, num%10
		}
	}
	index := 0
	for ; index < len(res) && res[index] == '0'; index++ {
	}
	for i := 0; i < len(res); i++ {
		res[i] += '0'
	}
	res = res[index:]
	return string(res)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		} else if nums[mid] > nums[l] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] == nums[mid] {
				l++
			}
			if nums[r] == nums[mid] {
				r--
			}
		}
	}
	return -1
}

func canPartition(nums []int) bool {

	return false
}
