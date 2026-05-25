// https://leetcode.com/problems/neither-minimum-nor-maximum/description/
struct Solution {}

impl Solution {
  /// Given an integer array `nums` containing **distinct positive integers**,
  /// find and return **any** number from the array that is neither the
  /// **minimum** nor the **maximum** value in the array, or `-1` if there is
  /// no such number.
  ///
  /// Return _any number that is neither the minimum nor the maximum of the
  /// array_. If no such number exists, return `-1`.
  ///
  /// **Constraints:**
  /// - `1 <= nums.length <= 100`
  /// - `1 <= nums[i] <= 100`
  /// - All values in `nums` are distinct.
  pub fn find_non_min_or_max(nums: Vec<i32>) -> i32 {
    if nums.len() < 3 {
      return -1;
    }
    0
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn neither_minimum_nor_maximum() {
    let tests = [(vec![3, 2, 1, 4], 2), (vec![1, 2], -1), (vec![2, 1, 3], 2)];

    for tc in tests {
      let (input, want) = tc;
      let actual = Solution::find_non_min_or_max(input.clone());
      assert_eq!(
        want, actual,
        "input={:?}, want={:?}, actual={:?}",
        input, want, actual
      );
    }
  }
}
