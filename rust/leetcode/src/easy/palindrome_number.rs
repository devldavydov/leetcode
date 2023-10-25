/*
https://leetcode.com/problems/palindrome-number/

Given an integer x, return true if x is a palindrome, and false otherwise.
*/

#[allow(unused)]
fn is_palindrome(x: i32) -> bool {
    if x < 0 {
        return false;
    }

    let mut xx = x;
    let mut reversed = 0;

    while xx != 0 {
        reversed = reversed * 10 + xx%10;
        xx /= 10;
    }

    reversed == x
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_palindrome() {
        assert!(is_palindrome(121));
        assert!(is_palindrome(0));
        assert!(is_palindrome(1));
        assert!(!is_palindrome(-121));
        assert!(!is_palindrome(10));
    }
}