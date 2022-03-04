
// ========================================================================
// 11 数组矩形的最大面积
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		h := Min(height[left], height[right])
		width := right - left
		res = Max(res, h*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// ========================================================================
// 接雨水
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	// start end 为左右指针下标遍历
	// l,r 存储高度
	start, end, l, r := 0, len(height)-1, 0, 0
	res := 0
	for start < end {
		if height[start] < height[end] {
			if height[start] < l {
				res += l - height[start]
			} else {
				l = height[start]
			}
			start++
		} else {
			if height[end] < r {
				res += r - height[end]
			} else {
				r = height[end]
			}
			end--
		}
	}
	return res
}

// ========================================================================
// 移动零 将数组中 0移到最后面，非零保持原相对顺序
func moveZero(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 双指针将 nums 分为三部分
	// 0~i 为非零，j~n 为未访问
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}
		// 非零则交换
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
}

// ========================================================================
// 四数之和-注意非重复的去重,转化为两数之和
func forSum(nums []int, targetSum int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	res, i := [][]int{}, 0
	sort.Ints(nums)
	// 外层循环和内层循环边界
	for ; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				tempSum := nums[left] + nums[right] + nums[i] + nums[j]
				if targetSum == tempSum {
					res = append(res, []int{nums[left], nums[right], nums[i], nums[j]})
					left++
					right--
					// 避免重复
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if tempSum < targetSum {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

// ========================================================================
// 最小糖果奖励 每个学生至少一个，且高分比相邻低分奖励多
// 规律：相邻的学生a<b 则糖果B=A+1或B=1
func candy(ratings []int) int {
	left, right := make([]int, len(ratings)), make([]int, len(ratings))
	// 保证至少一个
	// 从左右两侧比较
	for i := 0; i < len(ratings); i++ {
		left[i], right[i] = 1, 1
	}
	// 从左往右
	for i := 1; i < len(ratings); i++ {
		if left[i] > left[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	// 从右往左
	for j := len(ratings) - 2; j >= 0; j-- {
		if right[j] > right[j+1] {
			right[j] = right[j+1] + 1
		}
	}
	sum := 0
	for i := 0; i < len(ratings); i++ {
		sum += Max(left[i], right[i])
	}
	return sum
}

// ========================================================================
// 对角线遍历矩阵避免使用大量指针出现越界操作
// 分解操作再合并 第一行斜向左下遍历，最后一列斜向左下遍历
// 偶数位翻转后将所有合并到 res
func FindDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return nil
	}
	m, n := len(mat), len(mat[0])
	lists := make([][]int, 0)
	res := make([]int, m*n)
	for i := 0; i < m; i++ {
		row, col := 0, i
		temp := []int{}
		for row < n && col >= 0 {
			temp = append(temp, mat[row][col])
			row++
			col--
		}
		lists = append(lists, temp)
	}
	// 最后一列 左下遍历
	for i := 1; i < n; i++ {
		row, col := i, m-1
		temp := []int{}
		for row < n && col >= 0 {
			temp = append(temp, mat[row][col])
			row++
			col--
		}
		lists = append(lists, temp)
	}
	// 偶数位翻转
	for index := 0; index < len(lists); index++ {
		if index%2 == 0 {
			for i, j := 0, len(lists[index])-1; i < j; i, j = i+1, j-1 {
				lists[index][i], lists[index][j] = lists[index][j], lists[index][i]
			}
		}
	}
	// 二维数组转化为一维
	index := 0
	for i := 0; i < len(lists); i++ {
		for j := 0; j < len(lists[i]); j++ {
			res[index] = lists[i][j]
			index++
		}
	}
	return res
}

// ========================================================================

// ========================================================================
// 二分查找板子
// 适用于一堆数中找出指定数字 O(logN)
// 有序数组采用普通，无须猜答案
func binarysearchDemo(nums []int, target int) {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	// (end,start)
	// 相等情况直接返回 mid
}

// ========================================================================
// H-410-分割数组的最大值
// 将整数数组 分割为 m 份，使得 m 份各个子数组之和中的最大值最小
// 二分猜答案
func splitArray(nums []int, m int) int {
	temp := make([]int, 0)
	temp = append(temp, nums...)
	sort.Ints(temp)
	// max 为可分割子数组各自和中的最小值,sum 为可分割子数组各自和中的最大值
	// max sum 为二分的上下界
	max, sum := temp[len(temp)-1], 0
	for _, num := range temp {
		sum += num
	}
	return binary(nums, m, sum, max)
}

// 二分猜答案
func binary(nums []int, m, high, low int) int {
	mid := 0
	for low <= high {
		mid = low + (high-low)>>1
		// 符合分割情况则往中间靠
		// 可将 nums 分割为 m 份，其中有和为 mid
		if valid(nums, m, mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// 判断猜的答案是否符合子数组各自和最大值最小
// 可分割 m 份且每份和不超过subArraySum
func valid(nums []int, m, subArraySum int) bool {
	// count代表分割的分数，一旦 curSum 超过限定的subArraySum，则 count++
	// curSum缩回越界的 num 开始累加
	curSum, count := 0, 1
	for _, num := range nums {
		curSum += num
		if curSum > subArraySum {
			// curSum 大于 subArraySum 则丢弃之前和，从当前加入越界后的 num 开始计数
			// 并记录目前分割的分数 count
			curSum = num
			count++
			if count > m {
				return false
			}
		}
	}
	return true
}

// ========================================================================

// ========================================================================
// 路径和，二叉树上找到从根节点触发的路径和等于 target
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// 遇到叶子结点查看是否符合 target
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	// 递归左右子树并改变 target
	left := hasPathSum(root.Left, sum-root.Val)
	right := hasPathSum(root.Right, sum-root.Val)
	return left || right
}

// ========================================================================
// 二叉树的最大路径和，可能不经过 root
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MinInt32
	dfsMaxPathSum(&res, root)
	return res
}
func dfsMaxPathSum(res *int, node *TreeNode) int {
	if node == nil {
		return 0
	}
	left, right := dfsMaxPathSum(res, node.Left), dfsMaxPathSum(res, node.Right)
	// 经过当前结点的最大路径和全局最大路径比较
	*res = Max(*res, node.Val+left+right)
	// 返回经过当前路径的最大,左右两侧
	return Max(0, node.Val+Max(left, right))
}

// ========================================================================
// 翻转二叉树,镜面翻转
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 交换左右子树
	right := root.Right
	root.Right = root.Left
	root.Left = right

	// 递归每个子树下的分支
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// ========================================================================
// 二叉树拆分为链表
// 递归 dfs 或者stack 前序遍历时node.left=nil node.right=stack.peek
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		node.Left = nil
		if len(stack) == 0 {
			node.Right = nil
		} else {
			node.Right = stack[len(stack)-1]
		}

	}
}

// ========================================================================
// M-200 岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n, res := len(grid), len(grid[0]), 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				bfsNumsIslands(&grid, i, j)
				// dfsNumIslands(grid, i, j)
			}
		}
	}
	return res
}

// bfs 宽度优先遍历
func bfsNumsIslands(grid *[][]byte, row, col int) {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := make([][]int, 0)
	queue = append(queue, []int{row, col})
	for len(queue) != 0 {
		size := len(queue)
		for size > 0 {
			size--
			cur := queue[0]
			queue = queue[1:]
			x, y := cur[0], cur[1]
			if x < 0 || y < 0 || x >= len(*grid) || y >= len((*grid)[0]) || (*grid)[x][y] == '0' {
				continue
			}
			(*grid)[x][y] = '0'
			for _, dir := range dirs {
				queue = append(queue, []int{x + dir[0], y + dir[1]})
			}
		}
	}
}

// dfs 深度优先遍历
func dfsNumIslands(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	dfsNumIslands(grid, row-1, col)
	dfsNumIslands(grid, row+1, col)
	dfsNumIslands(grid, row, col+1)
	dfsNumIslands(grid, row, col-1)
}

// ========================================================================
// M-90-子集
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	// 非重复前提为有序
	sort.Ints(nums)
	batcktrackSubsSet(nums, &res, []int{}, 0)
	return res
}
func batcktrackSubsSet(nums []int, res *[][]int, temp []int, index int) {
	*res = append(*res, append([]int{}, temp...))
	// 子集需要从 index 开始，backtrack 需要从i+1
	for i := index; i < len(nums); i++ {
		if i > index && i < len(nums) && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		// recutsion 从 i+1 而非 index+1 开始
		batcktrackSubsSet(nums, res, temp, i+1)
		// 剪枝
		temp = temp[:len(temp)-1]
	}
}

// ========================================================================
// 非重复的全排列
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res, visited := make([][]int, 0), make([]bool, len(nums))
	// 非重复前提为有序
	sort.Ints(nums)
	backtrackPermuteUnique(nums, &res, &visited, []int{}, 0)
	return res
}
func backtrackPermuteUnique(nums []int, res *[][]int, visited *[]bool, cur []int, index int) {
	// return
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	for i := 0; i < len(nums); i++ {
		// 全排列去重
		// 与前一个相等并且前一个未visited||当前指向已visited
		if (*visited)[i] || (i > 0 && nums[i] == nums[i-1] && !(*visited)[i-1]) {
			continue
		}
		(*visited)[i] = true
		cur = append(cur, nums[i])
		// backtrack 从 index+1 开始
		backtrackPermuteUnique(nums, res, visited, cur, index+1)
		// 剪枝 注意(*visited)[i]==false
		cur = cur[:len(cur)-1]
		(*visited)[i] = false
	}
}

// 下一个稍大的排列
// M-31-下一个排列
func nextpermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 找到第一个降序 number
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}
	// 不存在更大的排列
	if i < 0 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		return
	}
	// 找到第一个比nums[i]稍小的数
	j := i + 1
	for ; j < len(nums) && nums[j] > nums[i]; j++ {
	}
	// 多加了一次
	j--
	// 将第一个降序的 number 和第一个比降序 number 稍大的数交换
	nums[i], nums[j] = nums[j], nums[i]
	i++
	// 将 i 之后部分逆序
	for j = len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// ========================================================================

// ========================================================================
// 最长回文子串的长度
func countSubstringsPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, dp := 0, make([][]bool, len(s))
	// dp[i][j]表示从i~j 位置内为回文串
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			// 状态转移方程表示
			// 当 dp[i+1][j-1]内部为回文串且两侧边界外的s[i]==s[j]相同则代表此处i,j 亦为回文
			// 考虑 s[i]==s[j]但长度 j-i<3 三个以内肯定为回文
			dp[i][j] = (s[i] == s[j]) && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] {
				if res < (j - i + 1) {
					res = j - i + 1
				}
			}
		}
	}
	return res
}

// ========================================================================
// 最大子数组之和
func maxSubArraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i]为以 i 为结束位置的最大子数组之和
	length, dp := len(nums), make([]int, len(nums))
	for i := 0; i < length; i++ {
		dp[i] = math.MinInt32
	}
	// 边界 dp[0]为第一个 num
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < length; i++ {
		if nums[i] <= dp[i-1]+nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if res <= dp[i] {
			res = dp[i]
		}
	}
	return res
}

// ========================================================================

// ========================================================================
// 合并 k 个有序链表
type heapLists []*ListNode

func (h heapLists) Len() int {
	return len(h)
}
func (h heapLists) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapLists) Less(i, j int) bool {
	return h[i].Val <= h[j].Val
}
func (h *heapLists) Pop() interface{} {
	old := *h
	x := old[h.Len()-1]
	old = old[:h.Len()-1]
	*h = old
	return x
}
func (h *heapLists) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// 合并 k 个有序链表
func mergeKSortedLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	h := &heapLists{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
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
		cur = node
	}
	return dummyHead.Next
}

//  最少的不同元素
// nums 删除 k 个元素后剩下不同的元素最少

// 重排字符串，使相同字符不相邻

// 有序矩阵中的第 k 小元素

// 最小范围，包含每个数组的至少一个数

// 查找和最小的 k 对数字

// ========================================================================

// ========================================================================
// 判断链表是否有环
// 快慢指针，起步快指针指向 head，慢指针指向 head.Next
// 一旦 fast 到达末尾则无环，否则最后一定有环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head.Next
	for slow != fast {
		if fast.Next == nil || fast == nil {
			return false
		}
		fast, slow = fast.Next.Next, slow.Next
	}
	return true
}

// ========================================================================
// 判断链表是否有环
// 使用 dict 字典，其中空结构体为 value 占位
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	m := make(map[*ListNode]struct{}, 0)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}
func detechCycle(head *ListNode) (*ListNode, bool) {
	if head == nil || head.Next == nil {
		return head, false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			fast = head
			for fast != slow {
				fast, slow = fast.Next, slow.Next
			}
			return fast, true
		}
	}
	return head, false
}

// 链表循环右移 k 位
// 链接为循环链表，找倒数 k 位断开即可，注意 prev 和 head 的更新
// ========================================================================

// ========================================================================
func Quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	return quicksortHelper(nums, 0, len(nums)-1)
}
func quicksortHelper(nums []int, l, r int) []int {
	if l >= r {
		return nums
	}
	p := paratition(nums, l, r)
	quicksortHelper(nums, l, p[0]-1)
	quicksortHelper(nums, p[1]+1, r)
	return nums
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
func Mergesort(nums []int) {

	if len(nums) <= 1 {
		return
	}
	mergeHelp(nums, 0, len(nums)-1)
}
func mergeHelp(nums []int, l, r int) {
	if l >= r {
		return
	}
	if l < r {
		mid := l + (r-l)>>1
		mergeHelp(nums, l, mid)
		mergeHelp(nums, mid+1, r)
		merge(nums, l, mid, r)
	}

}
func merge(nums []int, l, mid, r int) {
	pl, pr, helper := l, mid+1, make([]int, r-l+1)
	i := 0
	for pl <= mid && pr <= r {
		if nums[pl] <= nums[pr] {
			helper[i] = nums[pl]
			pl++
		} else {
			helper[i] = nums[pr]
			pr++
		}
		i++
	}
	copy(helper[i:], nums[pl:mid+1])
	copy(helper[i:], nums[pr:r+1])
	// 注意选择 l:r+1范围内
	copy(nums[l:r+1], helper)
}

func Heapsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		heapInsert(nums, i)
	}
	size := len(nums)
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapIfy(nums, 0, size)
	}
}
func heapInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[parent] <= nums[index] {
		nums[parent], nums[index] = nums[index], nums[parent]
		index = parent
		parent = (index - 1) >> 1
	}

}
func heapIfy(nums []int, index, size int) {
	left := index*2 + 1
	for left < size {
		largest := left
		if left+1 < size && nums[left+1] > nums[left] {
			largest = left + 1
		}
		if nums[index] > nums[largest] {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left = (index * 2) + 1
	}
}

// ========================================================================

// ========================================================================
// 最小子数组之和>target的长度
func minSubArrayLen(nums []int, target int) int {
	res := math.MaxInt32
	l, r, sum := 0, 0, 0
	// 右指针前进，不断累积右指针
	for r < len(nums) {
		sum += nums[r]
		// 符合条件则收缩
		for sum >= target {
			// 更新 res
			if r-l+1 < res {
				res = r - l + 1
			}
			// shrink
			sum -= nums[l]
			l++
		}
		r++
	}
	// 特例判断
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

// ========================================================================
// 至多有 k 个不同字符的最长子串
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	dict := make(map[byte]int, 0)
	l, r, res := 0, 0, 0
	for r < len(s) {
		// 右指针前进，不断计数出现的字符的次数
		dict[byte(s[r])]++
		// 出现字符个数符合条件，收缩
		for len(dict) > k {
			// shrink
			dict[byte(s[l])]--
			if dict[byte(s[l])] == 0 {
				delete(dict, byte(s[l]))
			}
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}
	return res
}

// ========================================================================
// 判断给定字符串 s 是否含有 p 的排列
// 定长滑动窗口，两个 map 分别记录 p 中字符次数和窗口中字符频次
func checkInclusion(s1, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}
	map1, map2 := make([]int, 26), make([]int, 26)
	// 首个窗口入 map2，同时记录 s1中的字符频次
	for i := 0; i < len1; i++ {
		map1[s1[i]-'a']++
		map2[s2[i]-'a']++
	}
	// 在字符串s2 上从 len1到 len2 滑动窗口
	// 窗口大小固定为 len1
	for i := len1; i < len2; i++ {
		if reflect.DeepEqual(map2, map1) {
			return true
		}
		// 加入新的字符
		map2[s2[i]-'a']++
		// 移出 i-len1 窗口最左侧的字符频次
		map2[s2[i-len1]-'a']--
	}
	return reflect.DeepEqual(map1, map2)
}

// ========================================================================
// 最小子串覆盖
// s 中包含 t 中每个字符的子串的最小子串
func MinWindow_76(s, t string) string {
	dict, start, l, r := make(map[byte]int, 0), 0, 0, 0
	length := math.MaxInt32
	match := 0
	// dict 存储 target 中字符频次
	for i := 0; i < len(t); i++ {
		dict[t[i]]++
	}
	// 在 s 上滑动窗口
	for r < len(s) {
		// 遇见s[r]符合条件
		// 在 dict 内，则将字符频次--
		if _, boolV := dict[s[r]]; boolV {
			dict[s[r]]--
			// 改字符在目前 l-r 内已经被匹配
			if dict[s[r]] == 0 {
				match++
			}
		}
		// 所有字符均在 s[l:r+1]内被匹配
		for match == len(dict) {
			// 更新有效子串的长度，同时记录起始位置
			if length > r-l+1 {
				length = r - l + 1
				start = l
			}
			// shrink left window boarder
			if count, boolValue := dict[s[l]]; boolValue {
				// 更新已经匹配的字符
				if count == 0 {
					match--
				}
				dict[s[l]]++
			}
			l++
		}
		r++
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// ========================================================================

// ========================================================================
func PreOrderTrvaersalBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, cur := make([]*TreeNode, 0), root
	res := []int{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Right
	}
	return res
}

func InorderTraversalBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, cur, res := make([]*TreeNode, 0), root, []int{}
	for cur != nil || len(stack) != 0 {
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

func PostorderTraversalBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, cur, stack := []int{}, root, make([]*TreeNode, 0)
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
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

// ========================================================================
func ZigzalTraversalaBT(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, res, cur := make([]*TreeNode, 0), [][]int{}, root
	queue = append(queue, cur)

	flag := false
	for len(queue) != 0 {
		size := len(queue)
		temp := []int{}
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if flag {
			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
		flag = !flag
		res = append(res, temp)
	}
	return res
}

// ========================================================================
// 前中后遍历构建二叉树
func ConstructBTFromPreInorderTraversal(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	num := preorder[0]
	index := findIndexInTraversalRes(inorder, num)
	return &TreeNode{
		Val:   num,
		Left:  ConstructBTFromPreInorderTraversal(preorder[1:index+1], inorder[:index]),
		Right: ConstructBTFromPreInorderTraversal(preorder[index+1:], inorder[index+1:]),
	}
}
func findIndexInTraversalRes(nums []int, num int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			res = num
			break
		}
	}
	return res
}

func ConstructBTFromPostInorderTraversal(postorder []int, inorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	num := postorder[len(postorder)-1]
	index := findIndexInTraversalRes(inorder, num)
	return &TreeNode{
		Val:   num,
		Left:  ConstructBTFromPostInorderTraversal(postorder[:index], inorder[:index]),
		Right: ConstructBTFromPostInorderTraversal(postorder[index:len(postorder)-1], inorder[index+1:]),
	}
}

// ========================================================================
// 递归前序遍历二叉树
func PreOrdreTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, root.Val)
	left := PreOrdreTraversalRecursion(root.Left)
	right := PreOrdreTraversalRecursion(root.Right)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

