package datawhale50

func isValid(s string) bool {
	m := make(map[rune]rune, 3)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['
	stack := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		// 左括号则追加
		case '(', '{', '[':
			stack = append(stack, rune(s[i]))
			// 右括号则查看栈顶是否为对应的左括号
		case ']', '}', ')':
			// 栈空则 false
			if len(stack) == 0 || stack[len(stack)-1] != m[rune(s[i])] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	// 最后判断栈是否空，非空则为奇数 false
	return len(stack) == 0
}
