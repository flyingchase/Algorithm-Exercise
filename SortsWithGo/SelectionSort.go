package sortswithgo

// 选择排序
func selectionSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	var Index int
	for i := 0; i < len(nums)-1; i++ {
		Index = i
		// 在未排序序列中找到最值元素，交换
		for j := i + 1; j < len(nums); j++ {
			if nums[Index] < nums[j] {
				Index = j
			}
		}
		nums[i], nums[Index] = nums[Index], nums[i]
	}
}
