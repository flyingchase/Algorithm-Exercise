package interviewoffer

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	if len(s) == 1 {
		return s[0]
	}
	dict := map[byte]int{}
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if dict[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
