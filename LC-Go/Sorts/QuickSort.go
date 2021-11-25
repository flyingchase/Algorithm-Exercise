package Sorts

func QuickSort(nums []int)[]int  {
	if nums == nil {
		return nil
	}
	quickSort(nums,0,len(nums)-1)
	return nums
}

func quickSort(nums []int, l int, r int) {
	if l>=r {
		return
	}
	p:=partition(nums,l,r)

	quickSort(nums,l,p[0]-1)
	quickSort(nums,p[1]+1,r)
}
func partition(nums []int, l int, r int) []int {
	less,more:=l-1,r
	for l < more {
		if nums[l]<nums[r] {
			less++
			Swap(nums,less,l)
			l++
		}else if nums[l]>nums[r] {
			more--
			Swap(nums,more,l)
		}else {
			l++
		}
	}
	Swap(nums,r,more)
	return []int{less+1,more}
}
