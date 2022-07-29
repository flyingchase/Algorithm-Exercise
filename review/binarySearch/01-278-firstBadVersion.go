package binarysearch

func firstBadVersion(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)>>1
		if !isBadVersion(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func isBadVersion(n int) bool {
	return false
}
