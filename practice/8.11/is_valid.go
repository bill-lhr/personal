package __11

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		b := s[i]
		if b == '(' || b == '{' || b == '[' {
			stack = append(stack, b)
		} else if b == ')' || b == '}' || b == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (b == ')' && top != '(') || (b == '}' && top != '{') || (b == ']' && top != '[') {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}
