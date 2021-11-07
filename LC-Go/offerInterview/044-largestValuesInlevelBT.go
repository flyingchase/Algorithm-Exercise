package offerInterview

import (
	"LC-Go/DataStructure"
	"container/heap"
	"fmt"
	"math"
	"reflect"
	"sort"
	"testing"
)

type TreeNode = DataStructure.TreeNode

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		maxElement := math.MinInt32
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxElement {
				maxElement = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if cur.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		res = append(res, maxElement)
	}
	return res
}

type Pair struct {
	value int
	count int
}
type PriorityQueue []Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	top := old[len(old)-1]
	*pq = old[:len(old)-1]
	return top
}
func topKthfrequenctInNums(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	topK := nums[k-1]
	fmt.Println(topK)
	m := map[int]int{}
	for _, v := range nums {
		m[v] += 1
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for val, cout := range m {
		heap.Push(pq, Pair{val, cout})
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(pq).(Pair).value)
	}
	return res
}

func postTraversalBT(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res

	}
	cur := root
	stack := make([]*TreeNode, 0)

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}
		nodeNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = nodeNode.Left
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func Test_postTraversalBT(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "123456",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},

			want: []int{2, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postTraversalBT(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postTraversalBT = %v,want %v", got, tt.want)
			}
		})
	}
}
func mergeExercise(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	mergeHelper(nums, 0, len(nums)-1)
	return nums
}
func mergeHelper(nums []int, l, r int) {
	if l < r {
		mid := l + (r-l)>>1
		mergeHelper(nums, l, mid)
		mergeHelper(nums, mid+1, r)
		mege(nums, l, mid, r)
	}

}
func mege(nums []int, l, mid, r int) {
	p1, p2, helper, i := l, mid+1, make([]int, r-l+1), 0
	for p1 <= mid && p2 <= r {
		if nums[p1] < nums[p2] {
			helper[i] = nums[p1]
			p1++
		} else {
			helper[i] = nums[p2]
			p2++
		}
		i++
	}
	copy(helper[i:], nums[p1:mid+1])
	copy(helper[i:], nums[p2:r+1])
	copy(nums[l:r+1], helper[:])

}
