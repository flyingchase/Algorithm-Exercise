package interviewoffer

func MinArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	l, r := 0, len(numbers)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)>>1
		if numbers[mid] > numbers[l] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid - 1
		}

	}
	return numbers[mid]
}
