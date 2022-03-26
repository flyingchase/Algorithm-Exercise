package codetop

import (
	"math"
)

// 升序排列
func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	heapsort(nums)
	return nums
}
func heapsort(nums []int) {
	size := len(nums)
	for i := 0; i < size; i++ {
		heapinsert(nums, i)
	}
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapify(nums, 0, size)
	}
}
func heapinsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[parent] <= nums[index] {
		nums[index], nums[parent] = nums[parent], nums[index]
		index = parent
		parent = (index - 1) >> 1
	}
}
func heapify(nums []int, index, size int) {
	left := 2*index + 1
	for left < size {
		largest := left
		if left+1 < size && nums[left] < nums[left+1] {
			largest = left + 1
		}
		if nums[largest] < nums[index] {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left = 2*index + 1
	}
}

func LengthOfLongestSubstringMeituan(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	dict := map[byte]struct{}{}
	l, r, res := 0, 0, 0
	for r < len(s) {

		if _, ok := dict[s[r]]; !ok {
			dict[s[r]] = struct{}{}
			r++

		} else {
			delete(dict, s[l])
			l++
		}
		if res < r-l {
			res = r - l
		}
	}
	return res
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	dict := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '{', '(':
			stack = append(stack, s[i])
		case ']', '}', ')':
			if len(stack) == 0 {
				return false
			}
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node != dict[s[i]] {
				return false
			}
		}
	}
	return len(stack) == 0
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummtyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummtyHead, head
	for cur != nil {
		next := cur.Next
		flag := false
		for next != nil && next.Val == cur.Val {
			next = next.Next
			flag = true
		}
		if flag {
			prev.Next = next
			cur = prev.Next
			continue
		}
		prev, cur = prev.Next, cur.Next
	}
	return dummtyHead.Next
}
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

type LRUCache struct {
	data       map[int]node
	cap        int
	head, tail *node
}
type node struct {
	prev, nex *node
	val       int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		data: map[int]node{},
		cap:  capacity,
		head: nil,
		tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}
func (this *LRUCache) remove(key int) {

}
func (this *LRUCache) get(key int) {

}
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	helper(root, &max)
	return max
}
func helper(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	l, r := maxMax(0, helper(node.Left, max)), maxMax(0, helper(node.Right, max))
	*max = maxMax(*max, l+r+node.Val)
	return maxMax(l, r) + node.Val
}
func maxMax(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// 两数组的最长公共子数组的长度
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	res := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] {
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
				if res < dp[i][j] {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}

// 最长回文子串
func longestPalindromeWith5(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = axial(s, i, i, res)
		res = axial(s, i, i+1, res)
	}
	return res
}
func axial(s string, l, r int, res string) string {
	sub := ""
	for l >= 0 && r < len(s) && s[l] == s[r] {
		sub = s[l : r+1]
		l--
		r++
	}
	if len(sub) < len(res) {
		return res
	}
	return sub
}
func longestPalindromeWith502(s string) string {
	if len(s) <= 1 {
		return s
	}
	l, r, pr, pl := 0, -1, 0, 0
	for l < len(s) {
		// 相同的边界
		for r+1 < len(s) && s[l] == s[r+1] {
			r++
		}
		for l > 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > pr-pl {
			pr, pl = r, l
		}
		l = l + (r-l)>>1 + 1
		r = l
	}
	return s[pl : pr+1]
}

// 连续最大和的子数组
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

// lca
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}

// lca another dfs
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// parent 记录各个子节点和对应的父节点映射
	parent, visited := map[int]*TreeNode{}, map[int]bool{}
	lcaDfs(root, &parent)

	// 从 p q 开始向上遍历，不断移动
	for p != nil {
		// 记录 p 所访问过得路径
		// 借助 map 向上查找
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}
func lcaDfs(node *TreeNode, parent *map[int]*TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		(*parent)[node.Left.Val] = node
		lcaDfs(node.Left, parent)
	}
	if node.Right != nil {
		(*parent)[node.Right.Val] = node
		lcaDfs(node.Right, parent)
	}
}

// k 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	for i := 1; i*k <= count; i++ {
		head = reverseBetween(head, (i-1)*k+1, i*k)
	}
	return head
}
func reverseBetween(head *ListNode, l, r int) *ListNode {
	if head == nil {
		return head
	}
	dummtyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummtyHead, head
	for i := 0; i < l-1; i++ {
		prev, cur = prev.Next, cur.Next
	}
	next := cur.Next
	for i := l; i < r; i++ {
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = cur.Next
	}
	return dummtyHead.Next
}
