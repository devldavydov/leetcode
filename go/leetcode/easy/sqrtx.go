package easy

/*
https://leetcode.com/problems/sqrtx/

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
*/

func Sqrtx(x int) int {
	if x <= 1 {
		return x
	}

	l, r, j := 0, x, x/2
	for {
		if j*j == x {
			return j
		} else if j*j > x {
			r = j
			j /= 2
		} else {
			l = j
			break
		}
	}

	for i := l; i <= r; i++ {
		if i*i > x {
			return i - 1
		}
	}

	return 0
}
