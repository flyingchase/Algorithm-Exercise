package alg

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j

}

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
