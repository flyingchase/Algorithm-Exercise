package interviewprepare

import (
	"reflect"
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	// window 存储 nums 中的 index
	window := make([]int, 0, k)
	res := make([]int, 0, len(nums)-k)
	for i := 0; i < len(nums); i++ {
		// 保证 window[0]存储的index为最大值
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		// 保证 window 入窗口均大
		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("testResult: %v, want: %v", got, tt.want)
			}
		})
	}
}
