package easy

/*
https://leetcode.com/problems/plus-one/

You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
Increment the large integer by one and return the resulting array of digits.
*/

func PlusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	transfer := 0
	digits[len(digits)-1] += 1

	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i] + transfer
		if d >= 10 {
			transfer, d = 1, d-10
		} else {
			transfer = 0
		}
		res[i+1] = d
	}
	if transfer != 0 {
		res[0] = transfer
	} else {
		res = res[1:]
	}

	return res
}
