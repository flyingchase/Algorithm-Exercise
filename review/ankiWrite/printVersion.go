package ankiwrite

// M-3
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	dict := make(map[byte]int, 0)
	res, l, r := 0, 0, 0
	for r < len(s) {
		if r < len(s) && dict[s[r]-'a'] == 0 {
			dict[s[r]-'a']++
			r++
		} else {
			dict[s[l]-'a']--
			l++
		}
		if res < r-l {
			res = r - l
		}
	}
	return res
}

// H-4
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap) Pop() int {

}
func (h *maxHeap) Push(index int) int {

}

func findMedianSortedArray() {

}
