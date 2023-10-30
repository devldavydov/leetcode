package easy

/*
https://leetcode.com/problems/roman-to-integer/

Given a roman numeral, convert it to an integer.
*/

var romanDigits = map[byte]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50,
	'C': 100, 'D': 500, 'M': 1000,
}

func RomanToInt(s string) int {
	var res, cv, rv int

	for i := len(s) - 1; i >= 0; i-- {
		cv = romanDigits[s[i]]
		if cv >= rv {
			res += cv
		} else {
			res -= cv
		}
		rv = cv
	}
	return res
}
