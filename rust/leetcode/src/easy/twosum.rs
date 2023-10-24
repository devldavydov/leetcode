/*
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

use std::collections::HashMap;

#[allow(unused)]
fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut m: HashMap<i32, i32> = HashMap::new();

    for item in nums.iter().enumerate() {
        match m.get(&(target-*item.1)) {
            None => {
                m.insert(*item.1, item.0 as i32);
            },
            Some(j) => {
                return vec![item.0 as i32, *j]
            }
        }
    }

    vec![]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum() {
        let mut res = two_sum(vec![2, 7, 11, 15], 9);
        assert!(res.contains(&0) && res.contains(&1));

        res = two_sum(vec![3, 2, 4], 6);
        assert!(res.contains(&1) && res.contains(&2));

        res = two_sum(vec![3, 3], 6);
        assert!(res.contains(&0) && res.contains(&1));
    }
}