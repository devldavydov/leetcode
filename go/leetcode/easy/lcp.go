package easy

import "strings"

/*
https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".
*/

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	lcp := ""
outer:
	for i := 1; i <= len(strs[0]); i++ {
		pr := strs[0][0:i]
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], pr) {
				break outer
			}
		}
		lcp = pr
	}

	return lcp
}
