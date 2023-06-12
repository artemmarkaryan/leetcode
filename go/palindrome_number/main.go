package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10022201))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	var digits = make([]int, 0)
	var devider = 10
	for {
		if x == 0 {
			break
		}
		digit := x % devider
		digits = append(digits, digit)
		x /= devider
	}

	fmt.Println(digits)

	for i := 0; i < len(digits)/2; i++ {
		a := digits[i]
		b := digits[len(digits)-1-i]
		if a != b {
			return false
		}
	}

	return true
}
