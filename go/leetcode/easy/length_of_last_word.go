package easy

/*
https://leetcode.com/problems/length-of-last-word/

Given a string s consisting of words and spaces, return the length of the last word in the string.
A word is a maximal substring consisting of non-space characters only.
*/

func LengthOfLastWord(s string) int {
	wlen, state, pStart := 0, 0, 0

	for i := range s {
		switch state {
		case 0:
			// Init
			if s[i] == ' ' {
				state = 1
			} else {
				state = 2
				pStart = i
			}
		case 1:
			// Outside word
			if s[i] != ' ' {
				state = 2
				pStart = i
			}
		case 2:
			// Inside word
			if s[i] == ' ' {
				state = 1
				wlen = i - pStart
			}
		}
	}
	if state == 2 {
		wlen = len(s) - pStart
	}

	return wlen
}
