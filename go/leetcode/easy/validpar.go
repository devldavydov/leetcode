package easy

import "container/list"

/*
https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

func IsValidParentheses(s string) bool {
	stack := list.New()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.PushFront(c)
		} else {
			p := stack.Front()
			if p == nil {
				return false
			}
			v, _ := p.Value.(rune)
			if (c == ')' && v != '(') ||
				(c == ']' && v != '[') ||
				(c == '}' && v != '{') {
				return false
			}
			stack.Remove(p)
		}
	}
	return stack.Len() == 0
}
