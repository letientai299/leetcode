// https://leetcode.com/problems/sum-of-values-at-indices-with-k-set-bits/description/
struct Solution {}

impl Solution {
  /// You are given a **0-indexed** integer array `nums` and an integer `k`.
  ///
  /// Return *the **sum** of elements in* `nums` *whose corresponding
  /// **indices** have **exactly*** `k` *set bits in their binary
  /// representation.*
  ///
  /// **Constraints:**
  /// - `1 <= nums.length <= 1000`
  /// - `1 <= nums[i] <= 10^5`
  /// - `0 <= k <= 10`
  pub fn sum_indices_with_k_set_bits(nums: Vec<i32>, k: i32) -> i32 {
    nums.iter().zip((0..nums.len())).fold(0, |sum, (&num, i)| {
      if i.count_ones() == k as u32 {
        sum + num
      } else {
        sum
      }
    })
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![5,10,1,5,2], 1, 13 ; "k=1")]
  #[test_case(vec![4,3,2,1], 2, 1 ; "k=2")]
  fn sum_k_set_bits(nums: Vec<i32>, k: i32, expected: i32) {
    assert_eq!(Solution::sum_indices_with_k_set_bits(nums, k), expected);
  }
}
