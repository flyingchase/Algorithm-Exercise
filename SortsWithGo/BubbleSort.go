package sortswithgo

// 冒泡排序
func bubbleSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		// 逐个比较相邻的元素
		for j := 0; j < len(nums)-i-1; i++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
