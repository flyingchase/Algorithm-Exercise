package HotReview

import "sort"

func solution1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)

	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	}
	return float64(nums1[len(nums1)/2] + nums2[len(nums2)/2+1])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums2)+1+len(nums1))>>1, 0, 0
	for low < high {
		nums1Mid = low + (high-low)>>1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			low = nums1Mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max4(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid==len(nums1) {
		midRight=nums2[nums2Mid]
	}else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min4(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2

}

func min4(i int, j int) int {
	if i < j {
		return i
	}
	return j

}

func max4(i int, j int) int {
	if i > j {
		return i
	}
	return j

}
