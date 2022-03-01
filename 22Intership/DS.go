package intership

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}
type ListNode struct {
	Next *ListNode
	Val  int
}

func Max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
func Min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
