package topkLists

func spiralorder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	count, res, sum := 0, []int{}, m*n
	top, bottom, left, right := 0, m-1, 0, n-1
	for count < sum {
		i, j := top, left
		for j <= right && count < sum {
			res = append(res, matrix[i][j])
			j++
			count++
		}
		i, j = top+1, right
		for i <= bottom && count < sum {
			res = append(res, matrix[i][j])
			i++
			count++
		}
		i, j = bottom, right-1
		for j >= left && count < sum {
			res = append(res, matrix[i][j])
			j--
			count++
		}
		i, j = bottom-1, left
		// 注意最后i>top
		for i > top && count < sum {
			res = append(res, matrix[i][j])
			i--
			count++
		}
		top, bottom, left, right = top+1, bottom-1, left+1, right-1
	}
	return res
}
