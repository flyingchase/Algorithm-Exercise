package round1
func search(nums []int, target int) int  {
	if len(nums)==0 {
		return -1
	}
	low,high:=0,len(nums)-1
	for low<=high {
		mid:=low+(high-low)>>1
		if nums[mid]==target {
			return mid
		}else if nums[mid]>nums[low] {
			if  {
			}
		}
	}
}
