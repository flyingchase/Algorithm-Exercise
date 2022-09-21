package findfeeling

import (
	"container/heap"
	"container/list"
	"findfeeling/model"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type TreeNode = model.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	flag := false
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, (node.Left))
			}
			if node.Right != nil {
				queue = append(queue, (node.Right))
			}
			tmp = append(tmp, node.Val)
		}
		if flag {
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
		flag = !flag
		if len(tmp) != 0 {
			res = append(res, tmp)
		}
	}
	return res
}

func postOrderTraversalBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := make([]*TreeNode, 0)
	cur := root
	// stack = append(stack, cur)
	// postOrderTraversalBT root left right
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = (cur.Right)
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Left
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func postOrderTraversalBT2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	stack := list.New()
	stack.PushFront(cur)
	res := []int{}
	for stack.Len() != 0 {
		element := stack.Front()
		stack.Remove(element)
		node := element.Value.(*TreeNode)
		res = append([]int{node.Val}, res...)
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
	}
	return res
}

func Quicksort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	quicksortHelper(nums, 0, len(nums)-1)
	return nums

}
func quicksortHelper(nums []int, l, r int) {
	if l > r {
		return
	}
	k := rand.Intn(r-l+1) + l
	nums[k], nums[r] = nums[r], nums[k]
	p := paratition(nums, l, r)
	quicksortHelper(nums, l, p[0]-1)
	quicksortHelper(nums, p[1]+1, r)
}

// 分割
func paratition(nums []int, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if nums[l] < nums[r] {
			less++
			nums[l], nums[less] = nums[less], nums[l]
			l++
		} else if nums[l] > nums[r] {
			more--
			nums[more], nums[l] = nums[l], nums[more]
		} else {
			l++
		}
	}
	nums[r], nums[more] = nums[more], nums[r]
	return []int{less + 1, more}
}

func Heapsort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		heapsortInsert(nums, i)
	}
	size := len(nums)
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapsortIfy(nums, 0, size)
	}
	return nums
}

func heapsortIfy(nums []int, index, size int) {
	left, right := 2*index+1, 2*index+2
	for left < size {
		largest := left
		if right < size && nums[right] > nums[left] {
			largest = right
		}
		if nums[largest] < nums[index] {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left, right = 2*index+1, 2*index+2
	}
}

func heapsortInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[index] > nums[parent] {
		nums[index], nums[parent] = nums[parent], nums[index]
		index = parent
		parent = (index - 1) >> 1
	}
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mergeHelper(nums, 0, len(nums)-1)
	return nums
}

func mergeHelper(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + ((r - l) >> 1)
	mergeHelper(nums, l, mid)
	mergeHelper(nums, mid+1, r)
	merge(nums, l, mid, r)

	return
}

func merge(nums []int, l, mid, r int) {
	p1, p2, helper := l, mid+1, make([]int, r-l+1)
	var index int
	for p1 <= mid && p2 <= r {
		if nums[p1] <= nums[p2] {
			helper[index] = nums[p1]
			p1++
		} else {
			helper[index] = nums[p2]
			p2++
		}
		index++
	}
	copy(helper[index:], nums[p1:mid+1])
	copy(helper[index:], nums[p2:r+1])
	copy(nums[l:r+1], helper[:])
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
		return l
	}
	return r
}

// 最长上升子序列
func lis(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j])))
			}
		}
		dp[i] = dp[i] + 1
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}

// 下一个排列
func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	index := len(nums) - 2

	for ; index >= 0; index-- {
		if nums[index] < nums[index+1] {
			break
		}
	}
	if index >= 0 {
		flag := len(nums) - 1
		for ; flag > index; flag-- {
			if nums[flag] > nums[index] {
				break
			}
		}
		nums[flag], nums[index] = nums[index], nums[flag]
	}

	for i, j := index+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums

}

// 旋转数组找定值
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1

	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[l] {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			if nums[r] >= target && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else {
			if nums[mid] == nums[l] {
				l++
			} else if nums[r] == nums[mid] {
				r--
			}
		}
	}
	return -1
}

// 组合总数等于定值
// candidates中每个数字只使用一次，不得包含重复数组
func CombinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	// 非重
	sort.Ints(candidates)
	res, _ := [][]int{}, make([]bool, len(candidates))

	tmp := make(map[int]struct{}, len(candidates))
	dfsCombinationSum(candidates, []int{}, 0, 0, target, tmp, &res)
	return res

}

func dfsCombinationSum(candidates []int, cur []int, index, sum, target int, visited map[int]struct{}, res *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		*res = append(*res, append([]int{}, cur...))
	}

	for i := index; i < len(candidates); i++ {
		if _, ok := visited[candidates[i]]; ok {
			continue
		}
		if i > index && candidates[i] == candidates[i-1] {
			if _, ok := visited[candidates[i]]; !ok {
				continue
			}
		}
		// if (visited)[i] || (i > index && candidates[i] == candidates[i-1] && !(visited)[i-1]) {
		// 	continue
		// }
		// if i>index && candidates[i]==candidates[i-1] && (_,fine:=visited[candidates[i-1]];!fine) {
		// 	continue
		// }
		sum += candidates[i]
		cur = append(cur, candidates[i])
		(visited)[candidates[i]] = struct{}{}
		dfsCombinationSum(candidates, cur, i+1, sum, target, visited, res)
		sum -= candidates[i]
		cur = cur[:len(cur)-1]
		delete(visited, candidates[i])
	}
}

// 最长不含重复字符的子字符串
func lengthofLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var freq [256]int
	res, l, r := 0, 0, -1
	for l < len(s) {
		if r+1 < len(s) && freq[s[r+1]-'0'] == 0 {
			freq[s[r+1]-'a']++
			r++
		} else {
			freq[s[l]-'a']--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}

func lengthofLongestSubstringlis(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp, res := make([]int, len(nums)), 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
			if res < dp[i] {
				res = dp[i]
			}
		}

	}
	return res
}

// lis use binarySearch
// 单调栈加二分
func Lis2(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	stack := make([]int, 0)
	stack = append(stack, nums[0])
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			index := binarySearch(nums, 0, len(stack)-1, stack[len(stack)-1])
			stack[index] = nums[i]
		}
		if res < len(stack) {
			res = len(stack)

		}
	}
	return res
}

func binarySearch(nums []int, l, r int, target int) int {
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

// 滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return nil
	}

	windows := make([]int, 0)
	var res []int
	for i := 0; i < len(nums); i++ {
		// specially
		if len(windows) > 0 && windows[0] <= i-k {
			windows = windows[1:]
		}

		// shrink
		for len(windows) > 0 && nums[windows[len(windows)-1]] < nums[i] {
			windows = windows[:len(windows)-1]
		}

		// extend
		windows = append(windows, nums[i])

		if i >= k-1 {
			res = append(res, windows[0])
		}
	}
	return res
}

// 滑动窗口的最大值
// 双端队列使用 lis 实现
type element struct {
	index int
	value int
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) < k {
		return nil
	}

	dequeue := list.New()
	var res []int
	for i := 0; i < len(nums); i++ {
		if i >= k-1 {
			res = append(res, dequeue.Back().Value.(element).value)
		}

	}

	return res

}

// 字符串相乘
func Multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m*n == 0 {
		return ""
	}
	res := make([]byte, m+n+2)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			tmp := (num1[i]-'0')*(num2[j]-'0') + res[i+j+1]
			res[i+j] += tmp / 10
			res[i+j+1] = tmp % 10

		}
	}
	for i := 0; i < len(res); i++ {
		res[i] += '0'
	}

	var index int
	for ; index < len(res); index++ {
		if res[index] != '0' {
			return string(res[index:])
		}
	}
	return "0"
}

func countSubStrings(s string) int {

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	var res int
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && ((i-j) < 2 || dp[i-1][j+1]) {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}

// LRUCache
// 双链表+map

// LFU 更新和插入在链表中任意位置，删除在链尾
// 访问频率相同按照新旧顺序
// 只关心最小访问频次，维护minFrequency即可无须排序算法

type LRUCache struct {
	cap, len int

	keys       map[int]*LruNode
	head, tail *LruNode
}
type LruNode struct {
	key, val   int
	prev, next *LruNode
}

// Get 将节点移动到链表头部
func (this *LRUCache) Get(key int) int {
	if node, ok := this.keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Add(node *LruNode) {

	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node

	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}
func (this *LRUCache) Remove(node *LruNode) {

	if this.head == node {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.prev = nil
		node.next = nil
		return
	}

	if this.tail == node {
		this.tail = node.prev
		if node.prev != nil {
			node.prev.next = nil
		}
		node.prev = nil
		node.next = nil
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
}

// Put 写入缓存
// 存在则更新并移至链首，不存在则插入到首位置
// 超容删除链尾
func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.keys[key]; ok {
		node.val = value
		this.Remove(node)
		this.Add(node)
		return
	}
	node := &LruNode{
		key: key,
		val: value,
	}
	this.len++
	this.keys[key] = node
	this.Add(node)

	if this.len > this.cap {
		delete(this.keys, this.tail.key)
		this.Remove(this.tail)
		this.len--
	}

}

// 最近最少访问频次置换
// 维护一个最小堆存储最小访问频次和最近访问次数

type LFUCache struct {
	keys     map[int]*Node
	pq       pq
	cap, len int
}

type Node struct {
	key, val  int
	frequency int
	count     int64
	index     int
}

type pq []*Node

func (p pq) Len() int {
	return len(p)
}
func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}
func (p pq) Less(i, j int) bool {
	if p[i].frequency == p[j].frequency {
		return p[i].count <= p[i].count
	}
	return p[i].frequency < p[j].frequency
}
func (p *pq) Push(x interface{}) {
	*p = append(*p, x.(*Node))
}
func (p *pq) Pop() interface{} {
	old := *p
	top := old[len(old)-1]
	*p = old[:len(old)-1]
	return top
}

func (p *pq) update(node *Node, value, frequency int, count int64) {
	node.count = count
	node.frequency = frequency
	node.val = value
	heap.Fix(p, node.index)
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		keys: map[int]*Node{},
		pq:   pq{},
		cap:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {

	if this.cap == 0 {
		return -1
	}

	if node, ok := this.keys[key]; ok {
		count := time.Now().Unix()
		node.count = count
		node.frequency++
		this.pq.update(node, node.val, node.frequency, node.count)
		return node.val
	}

	return -1

}

func (this *LFUCache) Put(key int, value int) {

	if this.cap == 0 {
		return
	}

	if node, ok := this.keys[key]; ok {
		this.pq.update(node, value, node.frequency+1, time.Now().Unix())
		return
	}

	node := &Node{
		val:   value,
		key:   key,
		count: time.Now().Unix(),
	}
	heap.Push(&this.pq, node)
	this.keys[key] = node
	this.len++

	if this.len > this.cap {
		node := heap.Pop(&this.pq).(*Node)
		delete(this.keys, node.key)
		this.len--
	}
}

// 螺旋遍历矩阵，返回遍历结果
func spiralorder(matrix [][]int) []int {
	res := []int{}
	m, n := len(matrix), len(matrix[0])
	if m*n == 0 {
		return res
	}
	var count int
	count, sum := 0, m*n
	left, right, top, bottom := 0, n-1, 0, m-1
	for count < sum {
		i, j := top, left
		for count < sum && j <= right {
			count++
			res = append(res, matrix[i][j])
			j++
		}
		i, j = top+1, right
		for count < sum && i <= bottom {
			count++
			res = append(res, matrix[i][j])
			i++
		}
		i, j = bottom, right
		for count < sum && j >= left {
			count++
			res = append(res, matrix[i][j])
			j--
		}
		i, j = bottom, left
		for count < sum && i > top {
			count++
			res = append(res, matrix[i][j])
			i--
		}
		top, left, right, bottom = top+1, left+1, right-1, bottom-1
	}
	return res
}

func lisWithBinarySearch(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	n := len(nums)
	stack := make([]int, 0)
	stack = append(stack, nums[0])

	for i := 1; i < n; i++ {
		if stack[len(stack)-1] < nums[i] {
			stack = append(stack, nums[i])
			continue
		}
		index := binarySearchLis(stack, 0, len(stack)-1, nums[i])
		stack[index] = nums[i]
	}
	return len(stack)
}

func binarySearchLis(nums []int, l int, r int, target int) int {
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
	return l
}

func numEnclaves(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				if matrix[i][j] == 0 {
					dfsNum(matrix, i, j)
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfsNum(matrix [][]int, i, j int) {
	matrix[i][j] = 0
	offsetX, offsetY := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
	for k := 0; k < 4; k++ {
		i += offsetX[k]
		j += offsetY[k]
		if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) && matrix[i][j] == 1 {
			dfsNum(matrix, i, j)
		}
	}
}

func mergeTwoLists2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummyHead := new(ListNode)
	p1, p2 := list1, list2
	cur := dummyHead
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
	}
	if p2 != nil {
		cur.Next = p2
	}
	return dummyHead.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil

	}
	if length == 1 {
		return lists[0]

	}

	mid := length / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeSelf(left, right)
}

func mergeSelf(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeSelf(left.Next, right)
		return left
	}
	right.Next = mergeSelf(right.Next, left)
	return right
}

func demo() {
	var str string = "abc  def ghi  "
	fmt.Println(str)
}
