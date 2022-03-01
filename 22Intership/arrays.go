package intership

// 判断是否为单调队列
// 使用位运算&=
func isMonotonic(A []int) bool {
	if len(A) == 0 {
		return false
	}
	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		inc = A[i-1] <= A[i]
		dec = A[i-1] >= A[i]
	}
	return inc || dec
}
