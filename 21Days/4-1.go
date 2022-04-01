package days

// 多数元素
// map 辅助空间记录出现的次数
// sort排序后 return nums[len(nums)>>1]
// 摩尔投票法
/*
candidate 初始化为 nums[0]，票数 count=1
遍历 nums 发现相同 candidate则 count+1，不同则-1
count==0 则更换当前候选人，重置 count=1
原理：多数元素与其他元素出现次数逐个抵消
*/
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	dict := map[int]int{}
	for index := 0; index < len(nums); index++ {
		dict[nums[index]]++
	}
	for k, v := range dict {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 删除链表中重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	for cur != nil {
		next := cur.Next
		flag := false
		for next != nil && next.Val == cur.Val {
			next = next.Next
			flag = true
		}
		if flag {
			prev = cur
			cur.Next = next
			cur = next
			continue
		}
		prev = cur
		cur = next
	}
	return dummyHead.Next
}

// 最长重复子数组的长度
func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j]代表 nums1[0:i+1]和 nums2[0:j+1]之间最长重复子串的长度
	m, n := len(nums1), len(nums2)
	if m*n == 0 {
		return 0
	}
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
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}

// 旋转图像，顺时针旋转 matrix 90 度
// 先对角线交换再轴对称交换
func rotate(matrix [][]int) {
	length := len(matrix)
	// 对角线转换
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 竖直轴对称翻转
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}
