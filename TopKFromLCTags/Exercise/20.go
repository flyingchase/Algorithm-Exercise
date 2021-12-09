package exercise

func isOkforrune(s string) bool {

	dict := make(map[rune]rune, 0)
	dict[']'] = '['
	dict['}'] = '}'
	dict[')'] = '('

	i := 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	stack := make([]rune, 0)

	for ; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, rune(s[i]))
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != dict[rune(s[i])] {
				return false
			}
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
