package easy

/*
https://leetcode.com/problems/merge-sorted-array/

You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
*/

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, n+m)

	i, j, k := 0, 0, 0
	for {
		if i == m && j == n {
			break
		} else if i <= m-1 && j <= n-1 {
			if nums1[i] <= nums2[j] {
				res[k] = nums1[i]
				i++
			} else {
				res[k] = nums2[j]
				j++
			}
			k++
		} else if i <= m-1 {
			res[k] = nums1[i]
			k++
			i++
		} else if j <= n-1 {
			res[k] = nums2[j]
			k++
			j++
		}
	}

	for i, v := range res {
		nums1[i] = v
	}
}
