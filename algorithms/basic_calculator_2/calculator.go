package calculator

func calculate(s string) int {
	var stack []int
	sign := byte('+')
	num := 0
	n := len(s)
	for i := 0; i < n; i++ {
		char := s[i]
		if isDigit(char) {
			num = num*10 + int(char-'0')
		}
		if !isDigit(char) && char != ' ' || i == n-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = char
			num = 0
		}
	}
	sum := 0
	for i := range stack {
		sum += stack[i]
	}
	return sum
}

func isDigit(char byte) bool {
	diff := char - '0'
	if diff >= 0 && diff <= 9 {
		return true
	}
	return false
}
