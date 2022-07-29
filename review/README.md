# review top alg

## Alg

### binary search

**数组存储、有序（无序需要猜答案）**

1. 正常 binary search

```go
func template(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := end + (start-end)>>1
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return nums
}
```



2. 猜答案



3. Lis





## DS