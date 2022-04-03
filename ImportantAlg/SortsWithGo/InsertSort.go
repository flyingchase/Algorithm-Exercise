package sortswithgo

// 插入排序
func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		// 第 i 元素大于 i-1元素则直接插入，小于则移动有序表后插入
		if nums[i] < nums[i-1] {
			j := i - 1
			tmp := nums[i]
			nums[i] = nums[i-1]
			// 查找待插元素在有序表内的正确位置
			for tmp < nums[j] {
				nums[j+1] = nums[j]
				j--
			}
			// 插入正确位置
			nums[j+1] = tmp
		}
	}
}
