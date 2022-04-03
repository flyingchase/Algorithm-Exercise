package reviewlc60

func correctParenthess(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	m := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '{', '[', '(':
			stack = append(stack, s[i])
		default:
			if m[s[i]] != stack[len(stack)-1] || len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
