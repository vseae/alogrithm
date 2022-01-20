package main

func isValid(s string) bool {
	table := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		//匹配左括号，入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			//如果匹配到右括号,弹出
		} else if len(stack) > 0 && stack[len(stack)-1] == table[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
