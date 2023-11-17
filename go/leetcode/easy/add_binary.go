package easy

/*
https://leetcode.com/problems/add-binary/

Given two binary strings a and b, return their sum as a binary string.
*/

func AddBinary(a string, b string) string {
	var maxLen int
	if len(a) > len(b) {
		maxLen = len(a)
	} else {
		maxLen = len(b)
	}

	res, transfer, dr := "", false, ""
	var ai, bi byte
	for i := 0; i < maxLen; i++ {
		ai, bi = '0', '0'
		if len(a)-i-1 >= 0 {
			ai = a[len(a)-i-1]
		}
		if len(b)-i-1 >= 0 {
			bi = b[len(b)-i-1]
		}

		switch {
		case ai == '0' && bi == '0' && !transfer:
			dr, transfer = "0", false
		case (ai == '0' && bi == '0' && transfer) ||
			(ai == '0' && bi == '1' && !transfer) ||
			(ai == '1' && bi == '0' && !transfer):
			dr, transfer = "1", false
		case (ai == '0' && bi == '1' && transfer) ||
			(ai == '1' && bi == '0' && transfer) ||
			(ai == '1' && bi == '1' && !transfer):
			dr, transfer = "0", true
		case ai == '1' && bi == '1' && transfer:
			dr, transfer = "1", true
		}
		res = dr + res
	}

	if transfer {
		res = "1" + res
	}

	return res
}
