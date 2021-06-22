package Stack

//用栈实现判断有效括号
//遍历字符串S， 如果是字符是左括号就入栈，并跳出当前循环继续执行下一次循环，
//如果字符是右括号，从map中找出映射左括号，如果映射左括号等于栈的最上一层字符，那么就出栈。最后返回栈长度是否为零。
//力扣20
func IsValid(s string) bool {
	stack := make([]byte, 0)
	m := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if value, ok := m[s[i]]; ok {
			if value == stack[len(stack)-1] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

//力扣224
//在遍历s的过程中，再未遇到(之前，一直计算并累计当前的计算结果；遇到(时，将当前计算结果压栈
func calculate(s string) int {
	stack := make([]int, 0)
	res := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for ; i < len(s); i++ {
				if s[i] < '0' || s[i] > '9' {
					break
				}
				num = num*10 + int(s[i]-'0') //这是得到两位数以上的数字 like 222
			}
			res = res + sign*num //sign表示符号位，减号的时候是-1
			i--
		} else {
			switch s[i] {
			case '+':
				sign = 1
			case '-':
				sign = -1 //减号时，符号位是-1
			case '(':
				stack = append(stack, res, sign) //遇到 '(' 把 值和符号位 追加到stack
				res, sign = 0, 1
			case ')':
				sign := stack[len(stack)-1]
				num := stack[len(stack)-2]
				res = num + sign*res //res 是 () 里面的值
				stack = stack[:len(stack)-2]
			}
		}
	}
	return res
}

//力扣 1047
func removeDuplicates(S string) string {
	if len(S) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	stack = append(stack, S[0])
	for i := 1; i < len(S); i++ {
		if len(stack) > 0 && S[i] == stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
			continue
		}
		stack = append(stack, S[i])
	}
	return string(stack)
}
