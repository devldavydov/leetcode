package easy

/*
https://leetcode.com/problems/palindrome-number/

Given an integer x, return true if x is a palindrome, and false otherwise.
*/

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversed := 0
	xx := x
	for xx != 0 {
		reversed = reversed*10 + xx%10
		xx /= 10
	}

	return reversed == x
}
