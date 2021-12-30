package binarysearch

import (
	"reflect"
	"testing"
)

func searchTimes(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}

		// 数组元素可能重复故而 left 和 mid 相等情况，需要移动 left
		for left <= mid && nums[left] == nums[mid] {
			left++
		}
		if left > mid {
			left = mid + 1
			continue
		}
		// left part is order
		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
func Test_searchTimes(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchTimes(tt.args.nums, tt.args.target); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got :%v, but want : %v", got, tt.want)
			}
		})

	}
}
