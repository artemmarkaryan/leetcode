// https://leetcode.com/problems/reverse-integer/

package main

const treshold int32 = (1<<31 - 1) / 10
const maxLastDigit int32 = (1<<31 - 1) % 10

func reverse(x int) int { return int(r(int32(x))) }
func r(x int32) (r int32) {
	var neg bool
	if x <= 0 {
		neg = true
		x = -x
	}
	for ; x > 0; x /= 10 {
		digit := x % 10
		if r > treshold || r == treshold && digit > maxLastDigit {
			return 0
		}

		r *= 10
		r += digit
	}
	if neg {
		return -r
	}
	return
}
