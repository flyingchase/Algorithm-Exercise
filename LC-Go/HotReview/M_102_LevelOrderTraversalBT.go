package HotReview

import (
	"reflect"
	"testing"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		temp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			temp = append(temp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, temp)
	}
	return res
}
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	cur := root
	queue = append(queue, cur)

	res := make([][]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		lists := make([]int, 0)
		for size > 0 {
			lists = append(lists, queue[0].Val)
			queue = queue[1:]
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			size--
		}
		res = append(res, lists)
	}
	return res
}

func Test_levelOrder2(t *testing.T) {

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				// 1,2,3,4,5,6
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
			want: [][]int{{1}, {2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelorder2 = %v , want %v ", got, tt.want)

			}
		})

	}
}
