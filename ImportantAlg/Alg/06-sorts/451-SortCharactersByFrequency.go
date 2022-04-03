package sorts

import (
	"sort"
)

// M-451-按照字符出现的次数排序
// 区分大小写
// 桶排序
func frequencySort_1(s string) string {
	if len(s) <= 1 {
		return s
	}
	// 存储每个字符出现的次数
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	// 桶，存储同一存储顺序下的字符,类似拉链法
	buckets := make([][]byte, len(s)+1)
	for i, _ := range buckets {
		buckets[i] = make([]byte, 0)
	}
	for char, freq := range m {
		buckets[freq] = append(buckets[freq], char)
	}
	res := []byte{}
	// 桶从后往前遍历，即出现频次由大到小
	for freq := len(buckets) - 1; freq >= 0; freq-- {
		lists := buckets[freq]
		for i := 0; i < len(lists); i++ {
			for j := 0; j < freq; j++ {
				res = append(res, lists[i])
			}
			// 使用 repeatAPI
			// res = append(res, bytes.Repeat([]byte{lists[i]}, freq)...)
		}
	}
	return string(res)
}

type entry struct {
	value byte
	freq  int
}

// 实现 len less swap 方法可使用 sort.sort()
//  实现 len less swap pop push 方法可使用 heap.push
func frequencySort(s string) string {
	if len(s) <= 1 {
		return s
	}
	var entrys []entry
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for value, freq := range m {
		entrys = append(entrys, entry{
			value: value,
			freq:  freq,
		})
	}
	sort.Slice(entrys, func(i, j int) bool {
		return entrys[i].freq > entrys[j].freq
	})
	res := []byte{}
	for i := 0; i < len(entrys); i++ {
		for entrys[i].freq > 0 {
			res = append(res, entrys[i].value)
		}
	}
	return string(res)
}
