package HotReview
func majorityElement(nums []int) int {
	n := len(nums) / 2
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
		if count[v] > n { return v }
	}
	return -1
}