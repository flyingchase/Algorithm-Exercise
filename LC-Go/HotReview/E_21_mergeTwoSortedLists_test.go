package HotReview

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
	{
		name : "FirstTest",
		args : args {
		    l1 : &ListNode {
		        Val: 1,
		        Next: &ListNode{
		        	Val: 2,
		        	Next: & ListNode {
		        	    Val: 4,
		        	    Next: nil,
		        	},
				},
		    },
		    l2 : & ListNode {
		        Val: 1,
		        Next: & ListNode {
		        	Val: 3,
		        	Next: & ListNode {
		        	    Val: 4,
		        	    Next: nil,
		        	},
		        },
		    },
		},
		want : & ListNode {
		    Val: 1,
		    Next: & ListNode {
		        Val: 1,
		        Next: &ListNode{
		        	2,&ListNode {
		        	    3,&ListNode {
		        	        4,&ListNode{4,nil},
		        	    },
		        	},
				},
		    },
		},
	},	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoList2(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name : "FirstTest",
			args : args {
				l1 : &ListNode {
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: & ListNode {
							Val: 4,
							Next: nil,
						},
					},
				},
				l2 : & ListNode {
					Val: 1,
					Next: & ListNode {
						Val: 3,
						Next: & ListNode {
							Val: 4,
							Next: nil,
						},
					},
				},
			},
			want : & ListNode {
				Val: 1,
				Next: & ListNode {
					Val: 1,
					Next: &ListNode{
						2,&ListNode {
							3,&ListNode {
								4,&ListNode{4,nil},
							},
						},
					},
				},
			},
		},



		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoList2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoList2() = %v, want %v", got, tt.want)
			}
		})
	}
}