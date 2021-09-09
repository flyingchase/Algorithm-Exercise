package HotReview

/*
给定数组中第一个缺失的正数
	- 将 nums 全部纳入 map 中 再遍历 index 找到存在 value 为 index
	- 第一个不存在的 index 即为缺失的首个正数
*/

func firstMissingPositive(nums []int) int {
	numMap:=make(map[int] int,len(nums))
	// 将 num存在 map 中 以 kv 均为 num 的形式
	for _,num:=range nums {
	    numMap[num]=num
	}
	for index:=1;index<len(nums)+1; index++{
	// 正数遍历中找到对应的第一个 v 不存在map 内即为所求
		if _,ok:=numMap[index];!ok {
			return index
		}
	}
	// 缺失的正数最多是数组长度加一
	return len(nums)+1
}
