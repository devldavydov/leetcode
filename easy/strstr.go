package easy

/*
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
*/

func StrStr(haystack string, needle string) int {
loop:
	for i := range haystack {
		if haystack[i] == needle[0] {
			for j := 1; j < len(needle); j++ {
				if i+j < len(haystack) && haystack[i+j] == needle[j] {
					continue
				}
				continue loop
			}
			return i
		}
	}

	return -1
}
