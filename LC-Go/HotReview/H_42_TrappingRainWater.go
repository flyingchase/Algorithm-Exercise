package HotReview
/*
容器接水  双指针
*/

func trap(height []int) int {
	res,left,right,maxLeft,maxRight:=0,0,len(height)-1,0,0
	// 两指针不断往中间靠
	for left<=right {
		// 左侧指针值小
		if height[left]<=height[right] {
			if height[left]>maxLeft {
				maxLeft=height[left]
			}else {
				// 将左侧 max超出的部分纳入 res
				// 由于 left<maxLeft则小的部分可以视为多出的空间）宽度为 1
				res+=maxLeft-height[left]
			}
			left++
		}else {
			if height[right]>=maxRight {
				maxRight=height[right]
			}else {
				res+=maxRight-height[right]
			}
			right--
		}
	}
	return res
}
