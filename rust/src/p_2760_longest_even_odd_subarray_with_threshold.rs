// https://leetcode.com/problems/longest-even-odd-subarray-with-threshold/description/
struct Solution {}

impl Solution {
  /// You are given a **0-indexed** integer array `nums` and an integer
  /// `threshold`.
  ///
  /// Find the length of the **longest subarray** of `nums` starting at index
  /// `l` and ending at index `r` (`0 <= l <= r < nums.length`) that
  /// satisfies the following conditions:
  ///
  /// - `nums[l] % 2 == 0`
  /// - For all indices `i` in the range `[l, r - 1]`,
  ///   `nums[i] % 2 != nums[i + 1] % 2`
  /// - For all indices `i` in the range `[l, r]`,
  ///   `nums[i] <= threshold`
  ///
  /// Return _an integer denoting the length of the longest such subarray_.
  ///
  /// **Note:** A **subarray** is a contiguous non-empty sequence of elements
  /// within an array.
  ///
  /// **Constraints:**
  /// - `1 <= nums.length <= 100`
  /// - `1 <= nums[i] <= 10^5`
  /// - `1 <= threshold <= 10^5`
  pub fn longest_alternating_subarray(nums: Vec<i32>, threshold: i32) -> i32 {
    let n = nums.len();

    let mut best: i32 = 0;
    let mut l = 0;

    while l < n {
      while l < n && (nums[l] % 2 != 0 || nums[l] > threshold) {
        l += 1;
      }

      if l >= n {
        break;
      }

      let mut r = l + 1;
      while r < n && nums[r] % 2 != nums[r - 1] % 2 && nums[r] <= threshold {
        r += 1
      }

      if best < (r - l) as i32 {
        best = (r - l) as i32;
      }

      l += 1
    }

    best
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn longest_even_odd_subarray_with_threshold() {
    let tests = [
      (vec![1], 1, 0),
      (vec![3, 2, 5, 4], 5, 3),
      (vec![1, 2], 2, 1),
      (vec![2, 3, 4, 5], 4, 3),
    ];

    for (nums, threshold, want) in tests {
      let actual =
        Solution::longest_alternating_subarray(nums.clone(), threshold);
      assert_eq!(
        want, actual,
        "nums={:?}, threshold={}, want={}, actual={}",
        nums, threshold, want, actual
      );
    }
  }
}
