package datawhale

func isValid(s string) bool {
	m := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '(', '{':
			stack = append(stack, s[i])
		case '}', ']', ')':
			node := stack[len(stack)-1]
			if node != m[node] {
				return false
			}
		}
	}
	return len(stack) == 0

}
