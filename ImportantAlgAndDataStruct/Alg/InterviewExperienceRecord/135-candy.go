package interviewexperiencerecord

import (
	"math"
	"sort"
)

/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
*/
// H-135-分发糖果
func candy(ratinngs []int) int {
	if len(ratinngs) == 0 {
		return 0
	}
	if len(ratinngs) == 1 {
		return 1
	}
	n := len(ratinngs)
	low, hight := n, n*(n+1)>>1
	res := math.MaxInt32
	for low <= hight {
		mid := low + (hight-low)>>1
		if validCandy(ratinngs, mid) {
			hight = mid - 1
			if res > mid {
				res = mid
			}
		} else {
			low = mid + 1
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return low
}

func validCandy(raitings []int, sum int) bool {
	temp := []int{}
	temp = append(temp, raitings...)
	sort.Ints(temp)
	count := len(raitings)
	for i := 1; i < len(temp)-1; i++ {
		if raitings[i] > raitings[i+1] && raitings[i] > raitings[i-1] {
			count++
			continue
		}
		if raitings[i] > raitings[i+1] || raitings[i] > raitings[i-1] {
			count++
		}
	}

	if count > sum {
		return false
	}
	return true
}
