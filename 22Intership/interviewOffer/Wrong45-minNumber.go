package interviewoffer

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return fmt.Sprint(nums[0])
	}
	res := ""
	sort.Slice(nums, func(i, j int) bool {
		p1, p2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		t1, t2 := []byte(p1), []byte(p2)
		t1 = append(t1, t2...)
		t2 = append(t2, t1...)
		n1, _ := strconv.Atoi(string(t1))
		n2, _ := strconv.Atoi(string(t2))
		return n1 < n2
	})
	for i := 0; i < len(nums); i++ {
		res += fmt.Sprint(nums[i])
	}
	return res
}
