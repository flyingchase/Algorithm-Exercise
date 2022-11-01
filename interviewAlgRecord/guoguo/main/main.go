package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{1, 0, 3, 7, 6, 4, 5, 2, 9, 8}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)

}

func lis(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	stack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		stack = append(stack, nums[i])
		if len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			index := binarySearch(stack, nums[i])
			stack[index] = nums[i]
		}
	}

	return len(stack)

}

func binarySearch(stack []int, target int) int {
	l, r := 0, len(stack)-1
	for l <= r {
		mid := l + (r-l)>>1
		if stack[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func splitArray(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	n, m := len(nums), sum>>1
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = false
			} else if j == 0 {
				dp[i][j] = true
			} else {
				if j >= nums[i-1] {
					dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[n][m]
}

func quickSort(nums []int, l, r int) {
	if l > r {
		return
	}

	rand.Seed(time.Now().Unix())
	k := rand.Intn(r-l+1) + l
	nums[k], nums[r] = nums[r], nums[k]
	p := paratition(nums, l, r)
	quickSort(nums, p[1]+1, r)
	quickSort(nums, l, p[0]-1)
}

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
	nums[more], nums[r] = nums[r], nums[more]
	return []int{less + 1, more}
}
