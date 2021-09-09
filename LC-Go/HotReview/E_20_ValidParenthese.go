package HotReview
/*
借助栈 stack 实现
	- map[rune]rune 存储左括号和对应的右括号
	- 遍历 string 左括号入栈 有括号比较 栈顶是否为 map 左括号对应的右括号
	- 每次抛出栈顶 stack=stack[:len(stack)]
	- 最后比较stack==0


*/
func isValid(s string) bool {
	stack:=make([]rune,0)
	m:=map[rune]rune{
		// 注意 key 为右括号 使得map 中的 vaule 为对应的栈顶（已经读入的左括号）
		'}':'{',
		']':'[',
		')':'(',
	}
	for _,c:=range s{
		switch c {
		case '[','{','(':
			stack=append(stack,c)
		case ']','}',')':
			if len(stack)==0||stack[len(stack) - 1]!=m[c] {
				return false
			}
			stack=stack[:len(stack) - 1]
		}
	}
	return len(stack) == 0
}
