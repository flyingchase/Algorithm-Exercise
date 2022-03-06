package main

import (
	intership "22intership"
	"fmt"
)

type TreeNode = intership.TreeNode

func main() {
	// 1,2,3,4,5,6
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//}
	//
	//fmt.Println(intership.PreOrderTrvaersalBT(root))
	//
	//matrix := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 0, 0, 0}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	//fmt.Println(len(matrix[0]))

	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	nums := strings.Split(input.Text(), " ")
	// 	res := 0
	// 	for i := 1; i < len(nums); i++ {
	// 		num, _ := strconv.Atoi(nums[i])
	// 		res += num
	// 	}
	// 	fmt.Println(res)
	// }

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intership.MovingAverage(nums, 3))
}
func changeDNA(a, b string) int {
	match := 0
	length := len(a)
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			match++
		}
	}
	if match%2 == 0 {
		return match / 2

	}
	return match/2 - 1
}
