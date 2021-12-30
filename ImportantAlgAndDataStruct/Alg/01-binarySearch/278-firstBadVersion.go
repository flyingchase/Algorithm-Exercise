package binarysearch

// 278 第一个错误的版本
func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		// mid:=left+(right-left)>>1
		// if !isBadVersion(mid) {
		// 	left=mid+1
		// }else {
		// 	right=mid-1
		// }
	}
	return left
}
