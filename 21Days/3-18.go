package days

import (
	"strconv"
)

// lca
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}

// lca2
/*
 使用 hashmap 存储父节点；
 从 p、q 分别想上遍历，存储遍历过得结点相同则为公共祖先
*/
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent, visited := map[int]*TreeNode{}, map[int]bool{}
	// 从当前根节点开始存储各个子节点的父节点
	lcaDfs(root, &parent)
	// 从q p 两结点向上遍历，并不断移动 p
	for p != nil {
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
func lcaDfs(r *TreeNode, parent *map[int]*TreeNode) {
	if r == nil {
		return
	}
	if r.Left != nil {
		(*parent)[r.Left.Val] = r
		lcaDfs(r.Left, parent)
	}
	if r.Right != nil {
		(*parent)[r.Right.Val] = r
		lcaDfs(r.Right, parent)
	}
}

// 搜索排序数组
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
			// 左侧有序，判断 target 是否落在 l~mid 之间
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			// 右侧有序，判断 target 是否落在mid~r 之间
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// 相同时移到边界
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

// 数1-n 的十进制中个数字中 1 的个数
// 答案为 n 的数位中各部分 1 之和
/*
例2304 为个十百千四部分出现 1 的次数之和
以十位数为例：**1* 最小为 0010 最大为 2219，其中有229-000+1=230种,为 n 千百位构成数*10
*/
// 暴力超时
func countDigitOne(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		tmp := []byte(strconv.Itoa(i))
		for j := 0; j < len(tmp); j++ {
			if tmp[j] == '1' {
				count++
			}
		}
	}
	return count
}
