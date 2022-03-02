package intership

import (
	"sort"
)

// 判断是否为单调队列
// 使用位运算&=
func isMonotonic(A []int) bool {
	if len(A) == 0 {
		return false
	}
	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		inc = A[i-1] <= A[i]
		dec = A[i-1] >= A[i]
	}
	return inc || dec
}

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

// 最小糖果奖励
// 每个学生至少一个，且高分比相邻低分奖励多
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

// 对角线遍历矩阵
// 避免使用大量指针出现越界操作
// 分解操作再合并
// 第一行斜向左下遍历，最后一列斜向左下遍历
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
