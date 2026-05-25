use std::cmp::max;

// https://leetcode.com/problems/longest-alternating-subarray/description/
struct Solution {}

impl Solution {
  /// You are given a **0-indexed** integer array `nums`. A subarray `s` of
  /// length `m` is called **alternating** if:
  ///
  /// - `m` is greater than `1`.
  /// - `s[1] = s[0] + 1`.
  /// - The subarray `s` follows the pattern:
  ///   `[s[0], s[0]+1, s[0], s[0]+1, ...]`. In other words,
  ///   `s[i] - s[i-1] = 1` if `i` is odd, and `s[i] - s[i-1] = -1` if `i`
  ///   is even (0-indexed within the subarray).
  ///
  /// Return _the length of the **longest alternating** subarray, or `-1` if
  /// no such subarray exists_.
  ///
  /// **Constraints:**
  /// - `2 <= nums.length <= 100`
  /// - `1 <= nums[i] <= 10^4`
  pub fn alternating_subarray(nums: Vec<i32>) -> i32 {
    let mut best = 0;
    for i in 1..nums.len() {
      let x = nums[i - 1];
      let y = nums[i];
      let mut len = 0;
      while y - 1 == x
        && i + len < nums.len()
        && nums[i + len] == y
        && nums[i + len - 1] == x
      {
        len += 2;
      }

      if len > 0 && i + len - 1 < nums.len() && nums[i + len - 1] == x {
        len += 1
      }

      best = max(best, len as i32);
    }

    if best == 0 {
      return -1;
    }

    best
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn longest_alternating_subarray() {
    let tests = [
      (vec![2, 3, 4, 3, 4], 4),
      (vec![21, 9, 5], -1),
      (vec![4, 5, 6], 2),
    ];

    for (nums, want) in tests {
      let actual = Solution::alternating_subarray(nums.clone());
      assert_eq!(
        want, actual,
        "nums={:?}, want={}, actual={}",
        nums, want, actual
      );
    }
  }
}
