package round1

func searchRange(nums []int, target int) []int {
	dict := make(map[int]struct{}, 0)
	for _, num := range nums {
		dict[num] = struct{}{}
	}
}
