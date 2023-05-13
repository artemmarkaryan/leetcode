package main

func isValid(s string) bool {
	var stack = make([]rune, 0)

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)

		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != map[rune]rune{
				']': '[',
				')': '(',
				'}': '{',
			}[c] {
				return false
			}

		default:
			return false
		}
	}

	return len(stack) == 0
}
