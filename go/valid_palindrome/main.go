package main

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for {
		if i >= j {
			return true
		}

		if invalid(s[i]) {
			i++
			continue
		}

		if invalid(s[j]) {
			j--
			continue
		}

		if toLower(s[i]) == toLower(s[j]) {
			i++
			j--
		} else {
			return false
		}
	}
}

func invalid(char uint8) bool {
	return true &&
		(char < 'A' || char > 'Z') &&
		(char < 'a' || char > 'z') &&
		(char < '0' || char > '9')
}

func toLower(char uint8) uint8 {
	if 'a' <= char && char <= 'z' {
		return char
	}

	return char + ('a' - 'A')
}
