// https://leetcode.com/problems/sum-of-squares-of-special-elements/description/
struct Solution {}

impl Solution {
  /// You are given a **1-indexed** integer array `nums` of length `n`.
  ///
  /// An element `nums[i]` of `nums` is called **special** if `i` divides
  /// `n`, i.e. `n % i == 0`.
  ///
  /// Return *the **sum of the squares** of all **special** elements of*
  /// `nums`.
  ///
  /// **Constraints:**
  /// - `1 <= nums.length == n <= 50`
  /// - `1 <= nums[i] <= 50`
  pub fn sum_of_squares(nums: Vec<i32>) -> i32 {
    (0..nums.len())
      .filter(|i| nums.len() % (i + 1) == 0)
      .map(|i| nums[i].pow(2))
      .sum()
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![1,2,3,4], 21 ; "nums=[1,2,3,4]")]
  #[test_case(vec![2,7,1,19,18,3], 63 ; "nums=[2,7,1,19,18,3]")]
  fn sum_of_squares_special(nums: Vec<i32>, expected: i32) {
    assert_eq!(Solution::sum_of_squares(nums), expected);
  }
}
