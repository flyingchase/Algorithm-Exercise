package codetop

import (
	intership "22intership"
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

type ListNode = intership.ListNode
type TreeNode = intership.TreeNode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	sort.Ints(nums)
	return nums[len(nums)-1-k]
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for i > 0 && nums[i] == nums[i-1] {
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
	res := [][]int{}
	sort.Ints(nums)
	dfsPermute(nums, &res, []int{})
	return res
}
func dfsPermute(nums []int, res *[][]int, cur []int) {
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		dfsPermute(nums, res, cur)
		cur = cur[:len(cur)-1]
	}
}

func TestPermute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); reflect.DeepEqual(got, tt.want) {
				t.Errorf("got:%v,but want:%v", got, tt.want)
			}
		})
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
		cur = max(nums[i], cur+nums[i])
		if cur > sum {
			sum = cur
		}
		if cur < 0 {
			cur = 0
		}
	}
	return sum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	dict := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '(', '{':
			stack = append(stack, s[i])
		case '}', ']', ')':
			if len(s) == 0 {
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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, res, cur := make([]*TreeNode, 0), [][]int{}, root
	queue = append(queue, cur)
	for len(queue) != 0 {
		size := len(queue)
		tep := []int{}
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			tep = append(tep, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tep)
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, cur, res := make([]*TreeNode, 0), root, []int{}
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		cur = node.Right
	}
	return res
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			fast = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return fast
		}
	}
	return nil
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, cur := make([]*TreeNode, 0), root
	queue = append(queue, cur)
	deepth := 0
	for len(queue) != 0 {
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		deepth++
	}
	return deepth
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var freq [256]int
	res, l, r := 0, 0, 0
	for r < len(s) {
		if r < len(s) && freq[s[r]-'a'] == 0 {
			freq[s[r]-'a']++
			r++
		} else {
			freq[s[l]-'a']--
			l++
		}
		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	nums := append(nums1, nums2...)

	min, max := &minHeap{}, &maxHeap{}
	heap.Init(min)
	heap.Init(max)
	for i := 0; i < len(nums); i++ {
		if min.Len() == max.Len() {
			heap.Push(max, nums[i])
			node := heap.Pop(max).(int)
			heap.Push(min, node)
		} else {
			heap.Push(min, nums[i])
			node := heap.Pop(min).(int)
			heap.Push(max, node)
		}
	}
	if min.Len() == max.Len() {
		return (float64(heap.Pop(min).(int)) + float64(heap.Pop(max).(int))) * 0.5
	}
	return float64(heap.Pop(min).(int))
}

type (
	minHeap []int
	maxHeap []int
)

func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *minHeap) Pop() interface{} {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}
func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *maxHeap) Pop() interface{} {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}

// 中心轴对称法则
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		res = checkDoubleSide(s, i, i+1, res)
		res = checkDoubleSide(s, i, i, res)
	}
	return res
}
func checkDoubleSide(s string, l, r int, res string) string {
	sub := ""
	for r < len(s) && l >= 0 && s[l] == s[r] {
		sub = s[l : r+1]
		l--
		r++
	}
	if len(sub) < len(res) {
		return res
	}
	return sub
}

func longestPalindromeWindows(s string) string {
	if len(s) == 0 {
		return ""
	}
	l, r, pl, pr := 0, 0, 0, 0
	for l < len(s) {
		// 重复边界
		for r < len(s) && s[l] == s[r] {
			r++
		}
		for l > 0 && r < len(s)-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if pr-pl < r-l {
			pl, pr = l, r
		}
		l = (r-l)>>1 + 1
		r = l
	}
	return s[pl : pr+1]
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, cur, queue := []int{}, root, make([]*TreeNode, 0)
	queue = append(queue, cur)
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
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
		if len(tmp) != 0 {
			res = append(res, tmp[len(tmp)-1])
		}
	}
	return res
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev := dummyHead
	p1, p2 := head, head
	for p1 != nil {
		for n > 0 {
			n--
			p1 = p1.Next
		}
		if p1 != nil {
			p1, p2 = p1.Next, p2.Next
			prev = prev.Next
		}
	}
	prev.Next = p2.Next
	return dummyHead.Next
}

func LengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	// dict存储 char-index
	dict := make(map[byte]int, 0)
	res, l, r := 0, 0, 0
	for r < len(s) {
		// 右指针所指的 char 出现过且下标>左指针
		if _, ok := dict[s[r]]; ok && dict[s[r]] >= l {
			l = dict[s[r]] + 1
		}
		if res < r-l+1 {
			res = r - l + 1
		}
		dict[s[r]] = r
	}
	return res
}

func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cur, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if  {
			
		}
	}
}
