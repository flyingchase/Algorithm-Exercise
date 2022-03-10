package interviewoffer

func findRepetNumber(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}
	m := map[int]int{}
	m[nums[0]]++
	for i := 1; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]]++
		} else {
			return nums[i]
		}
	}
	return -1
}
