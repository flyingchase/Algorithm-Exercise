package exercise

func reverseListNodes(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}
	dummyHead := &ListNode{Val: -1, Next: root}
	prev, cur := dummyHead, root
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

// 测试 revers 函数
// func Test_reverseListNodes(t *testing.T) {
// 	type args struct {
// 		root *ListNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *ListNode
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				root: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
// 			},
// 			want: &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := reverseListNodes(tt.args.root); reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("got: %v, but want: %v", got, tt.want)
// 			}
// 		})
//
// 	}
//
// }
