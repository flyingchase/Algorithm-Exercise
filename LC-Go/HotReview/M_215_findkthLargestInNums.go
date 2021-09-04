package HotReview

import (
	"math/rand"
	"sort"
)

/*给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。*/

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest1(nums []int, k int) int {
	m:=len(nums)-k+1
	return selectsmallest(nums,0,len(nums)-1,m)
}

func selectsmallest(nums []int, l int, r int, i int) int {
	if l >= r {
		return nums[l]
	}
	q:=paratition(nums,l,r)
	k:=q-l+1
	if k == i {
		return nums[q]
	}
	if i<k {
		return selectsmallest(nums, l, q-1, i)
	}else {
		return selectsmallest(nums, q+1, r, i-k)
	}
}

func paratition(nums []int, l int, r int) int {
	k:=l+rand.Intn(r-l+1)
	nums[k],nums[r]=nums[r],nums[k]
	i:=l-1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++;
			nums[i],nums[j]=nums[j],nums[i]
		}
	}
	nums[i+1],nums[r]=nums[r],nums[i+1]
	return i+1

}
