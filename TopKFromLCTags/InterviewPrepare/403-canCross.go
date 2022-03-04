package interviewprepare

func canCross(stones []int) bool {
	if len(stones) == 2 && stones[1] != 1 {
		return false
	}
	for i := len(stones) - 2; i >= 0; i-- {
		res := backtrack(stones, i, stones[len(stones)-1]-stones[i])
		if res {
			return true
		}
	}
	return false
}

func backtrack(stones []int, curIndex, dif int) bool {
	if curIndex == 0 && dif == 1 {
		return true
	}

	for i := curIndex - 1; i >= 0; i-- {
		if stones[curIndex]-stones[i] > dif+1 || stones[curIndex]-stones[i] < dif-1 {
			continue
		}
		res := backtrack(stones, i, stones[curIndex]-stones[i])
		if res {
			return true
		}
	}
	return false
}
