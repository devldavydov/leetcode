package easy

/*
https://leetcode.com/problems/climbing-stairs/

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

func ClimbingStairs(n int) int {
	stairs := make([]int, n)
	for i := 0; i < n; i++ {
		switch i {
		case 0:
			stairs[i] = 1
		case 1:
			stairs[1] = 2
		default:
			stairs[i] = stairs[i-1] + stairs[i-2]
		}
	}
	return stairs[n-1]
}
