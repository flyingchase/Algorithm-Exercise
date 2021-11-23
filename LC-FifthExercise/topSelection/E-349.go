package topselection

import (
	"math"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	m1 := make(map[int]int, 0)
	m2 := make(map[int]int, 0)
	for _, num := range nums1 {
		m1[num]++
	}
	res := []int{}
	for _, num := range nums2 {
		m2[num]++
	}
	sort.Ints(nums2)
	for i := 0; i < len(nums2)-1; i++ {
		if nums2[i] == nums2[i+1] {
			nums2 = append(nums2[:i], nums2[i+1:]...)
			i--
		}
	}
	for _, num := range nums2 {
		_, ok := m1[num]
		if !ok {
			continue
		}
		count := math.Min(float64(m1[num]), float64(m2[num]))
		for count > 0 {
			count--
			res = append(res, num)
		}
	}
	// sort.Ints(res)
	// for i := 0; i < len(res)-1; i++ {
	// 	if res[i] == res[i+1] {
	// 		res = append(res[:i], res[i+1:]...)
	// 		i--
	// 	}
	// }
	return res
}
