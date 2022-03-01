package reviewlc60

import "sort"

type entry struct {
	start int
	end   int
}

func mergeIntervials(intervials [][]int) [][]int {
	if len(intervials) <= 1 {
		return intervials
	}
	entrys := make([]*entry, 0)
	for _, intervial := range intervials {
		entry := &entry{
			start: intervial[0],
			end:   intervial[1],
		}
		entrys = append(entrys, entry)
	}
	sort.Slice(entrys, func(i, j int) bool {
		if entrys[i].start >= entrys[j].start {
			return true
		}
		return entrys[i].end >= entrys[j].end
	})
	tempRes := make([]*entry, 0)
	tempRes = append(tempRes, &entry{
		start: entrys[0].start,
		end:   entrys[0].end,
	})
	cur := 0
	for i := 1; i < len(entrys); i++ {
		if entrys[i].start > tempRes[cur].end {
			cur++
			tempRes = append(tempRes, entrys[i])
		} else {
			if entrys[i].end > tempRes[cur].end {
				tempRes[cur].end = entrys[i].end
			}
		}
	}
	res := [][]int{}
	for _, entry := range tempRes {
		temp := []int{}
		temp = append(temp, entry.start, entry.end)
		res = append(res, temp)
	}
	return res
}
