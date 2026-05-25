use std::cmp::max;

// https://leetcode.com/problems/max-pair-sum-in-an-array/description/
struct Solution {}

impl Solution {
  /// You are given an integer array `nums`. You have to find the **maximum**
  /// sum of a pair of numbers from `nums` such that the **largest digit** in
  /// both numbers is equal.
  ///
  /// Return the **maximum** sum or `-1` if no such pair exists.
  ///
  /// **Constraints:**
  /// - `2 <= nums.length <= 100`
  /// - `1 <= nums[i] <= 10^4`
  pub fn max_sum(nums: Vec<i32>) -> i32 {
    let mut best = [-1; 10];
    let mut ans = -1;

    for x in nums {
      let d = Self::top_digit(x) as usize;
      if best[d] != -1 {
        ans = max(ans, best[d] + x);
      }
      best[d] = max(best[d], x);
    }

    ans
  }

  fn top_digit(mut n: i32) -> i32 {
    let mut d = 0;
    while n > 0 {
      d = max(d, n % 10);
      n /= 10;
    }
    d
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![112,131,411], -1 ; "no pair with same max digit")]
  #[test_case(vec![2536,1613,3366,162], 5902 ; "all share max digit 6")]
  #[test_case(vec![51,71,17,24,42], 88 ; "71+17=88")]
  fn max_pair_sum(nums: Vec<i32>, expected: i32) {
    assert_eq!(Solution::max_sum(nums), expected);
  }
}
