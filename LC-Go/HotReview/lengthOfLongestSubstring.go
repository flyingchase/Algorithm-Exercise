package HotReview

func lengthofLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int

	res, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}

		res = max(res, right-left+1)

	}
	return res
}

func max(j int, i int) int {
	if i > j {
		return i
	}else {
	    return j
	}


}
func lengthofLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitset [256]bool
	res,left,right:=0,0,0
	for left<len(s) {
		if bitset[s[right]] {
			bitset[s[left]]=false
			left++
		}else {
			bitset[s[right]]=true
			right++
		}
		if right-left>res {
			res=right-left
		}
		if left +res>=len(s)||right>=len(s){
			break
		}
	}
	return res
}
func main() {
	substring := lengthofLongestSubstring("asdasdasdqw")

	print(substring)
}

func lengthOfLongestSubstring3(s string)int  {
	if len(s) == 0 {
		return 0
	}
	res,left,right:=0,0,0
	indexes:=make(map[byte]int,len(s))
	for left<len(s) {
		if idx,ok:=indexes[s[left]];ok&&idx>=right {
			right=idx+1
		}
		indexes[s[left]]=left
		left++
		res=max(res,left-right)
	}
	return res
}

