package topkTags

// 最长公共子串和最长公共子序列
func longestCommonSubstring(a string, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	m, n := len(a), len(b)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i] == b[j] {
				if i >= 1 && j >= 1 {
					// 为二维表格斜上方值加一
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = 0
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}

func longestCommonSubsequence(text1, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	res := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					// 相等则斜上角+1
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
			} else {
				if i > 0 && j > 0 {
					// 不等则正上方和正左方的最大值+1
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				} else if i == 0 && j > 0 {
					// 边界处理
					dp[i][j] = dp[i][j-1]
				} else if i > 0 && j == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = 0
				}
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 最短编辑距离问题
// 将一个字符串 a 编辑为另外一个字符串 b 所需的最小编辑次数
// 编辑动作为：替换 删除 插入

/*
1. 在字符串头部插入示空的字符，将 dp 表格为 (m+1)*(n+1)
2. 第一行，第一列则为单字符串不断插入对应的字符
	dp[0][j]=j dp[i][0]=i
3. dp[i][j]为原字符串以 i-1结尾编辑到原字符串 j-1结尾的最小编辑距离
考虑左上方:
	尾字符相等则拷贝左上方值；否则最多左上方+1(直接替换最后一位)
考虑左侧/上侧：
	left=dp[i][j-1]=1  up=dp[i-1][j]+1
4. 会用到左方 上方 左上方值，因此添加空字符串使得第一行和第一列被初始化
5. res:=dp[m][n]
*/
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	m, n := len(word1), len(word2)
	// 假设 word1和 word2已经插入首字符
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 首行不必再考虑
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			left, up := dp[i][j-1]+1, dp[i-1][j]+1
			left_up := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_up = dp[i-1][j-1] + 1
			}
			dp[i][j] = min(min(left, up), left_up)
		}
	}
	return dp[m][n]
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// 搜索二维矩阵
/*
在二维有序矩阵中查找 target 每行每列均是有序
1. 从右上角开始vertex，左侧 dec 右侧 inc
2. 不断从右上角思考，target>vertex 则排除目前首行row++，target<vertex则排除目前最后一列col--
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	vertex := matrix[row][col]
	for row < m && col >= 0 {
		vertex = matrix[row][col]
		if target == vertex {
			return true
		} else if target < vertex {
			col--
		} else if target > vertex {
			row++
		}
	}
	return false
}

// 给定数组的最大连续子序列之和
// 前缀和，找出 j 左侧的前缀和的最小值
func LargestSumContiguousSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min, sum, max := 0, 0, nums[0]
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		delta := sum - min
		if delta > max {
			max = delta
		}
		if sum < min {
			min = sum
		}
	}
	return max
}

// dp dp[i]代表 i 下标的最大数组和
func LargestSumContiguousSubarrayDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, dp := 0, make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// 加法实现
// 位运算
func Add(a, b int) int {
	if a == 0 {
		return b
	}
	for a != 0 {
		c := (a & b) << 1
		s := a ^ b
		a = c
		b = s
	}
	return b
}

// 乘法
// 1. 重复多次加法
// 2. 使用二分,把较小的数二分
func MulRecursive(a, b int) int {
	var max, min int
	if a > b {
		max, min = a, b
	} else {
		max, min = b, a
	}
	if min%2 == 1 {
		return MulRecursive(min-1, max) + max
	}
	half := MulRecursive(min>>1, max)
	return half << 1
}

// 求平方根
// 1. 二分法
// 2. 牛顿迭代法

func Sqrt(a int) int {
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	l, r := 0, a
	for l < r {
		mid := l + (r-l)/2
		if mid < a/mid {
			l = mid + 1
		} else if mid > a/mid {
			r = mid - 1
		} else {
			return mid
		}
	}
	if l > a/l {
		return l - 1
	}
	return l
}

func SqrtNewton(a int) float64 {
	if a == 0 {
		return 0
	}
	x, d := float64(a)*1.0, 0.0
	for x-float64(a)/x > 0 {
		d = x/2.0 - float64(a)/x/2.0
		if d < 1e-6 {
			break
		}
	}
	return x
}

// 最长回文子串
// 1. 中心轴对称法
//	无须奇偶讨论
// 2. dp[i][j]代表 i-j 位置上最长子串长度
//	注意 j-i<3特殊处理
// 3. 双指针
// 4. 马拉车算法
//	先处理源字符串变为奇数，两端和中间插入异字符
//	构造回文半径数组，p[i]代表i 下标处的回文长度；回文半径==原字符串回文长度
func LongestPalindromeSubstring01(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(s); i++ {
		res = helper01(s, i, i, res)
		res = helper01(s, i, i+1, res)
	}
	return res
}
func helper01(s string, l, r int, res string) string {
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

func LongestPalindromeSubstring02(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || (j-i+1) > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func LongestPalindromeSubstring03(s string) string {
	if len(s) == 0 {
		return ""
	}
	l, r, pl, pr := 0, 0, 0, 0
	for l < len(s) {
		// 找到不相同的边界
		for r < len(s) && s[l] == s[r] {
			r++
		}
		// 移动到回文子串的两端
		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > pr-pl {
			pl, pr = l, r
		}
		// 重置回文中心
		l = (r-l)>>1 + 1
		r = l
	}
	return s[pl : pr+1]
}

// 马拉车
func LongestPalindromeSubstring04(s string) string {
	return s
}
