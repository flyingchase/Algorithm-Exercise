package days

// 相交链表的相交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	cura, curb := headA, headB
	counta, countb := 0, 0
	for cura != nil {
		cura = cura.Next
		counta++
	}
	for curb != nil {
		countb++
		curb = curb.Next
	}
	aPlus, bPlus := false, false
	gap := 0
	if counta > countb {
		aPlus = true
		gap = counta - countb
	} else {
		bPlus = true
		gap = countb - counta
	}
	cura, curb = headA, headB
	for gap > 0 {
		if aPlus {
			cura = cura.Next
		}
		if bPlus {
			curb = curb.Next
		}
		gap--
	}
	for cura != curb {
		cura, curb = cura.Next, curb.Next
	}
	return cura
}

// 二叉树 zigzag 遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, cur := make([]*TreeNode, 0), root
	queue = append(queue, cur)
	flag := false
	res := [][]int{}
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != node {
				queue = append(queue, node.Right)
			}
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

// valid 括号
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	dict := map[byte]byte{
		']': '[',
		')': '(',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '{', '(':
			stack = append(stack, s[i])
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != dict[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 整数拆分
// 将正整数 n 拆分为 k 个整数（和为 n），求最大乘积
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			// i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，在求最大值的时候，一个是j与i-j相乘，一个是j与dp[i-j].
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
