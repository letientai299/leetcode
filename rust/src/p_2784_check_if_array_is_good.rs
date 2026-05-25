// https://leetcode.com/problems/check-if-array-is-good/description/
struct Solution {}

impl Solution {
  /// You are given an integer array `nums`. We consider an array **good** if
  /// it is a permutation of an array `base[n]`.
  ///
  /// `base[n] = [1, 2, ..., n-1, n, n]` (an array of length `n + 1` which
  /// contains `1` to `n - 1` exactly once, plus **two occurrences** of `n`).
  ///
  /// Return `true` *if the given array is good, otherwise return* `false`.
  ///
  /// **Constraints:**
  /// - `1 <= nums.length <= 100`
  /// - `1 <= nums[i] <= 200`
  pub fn is_good(mut nums: Vec<i32>) -> bool {
    let n = nums.len() - 1;
    nums.sort();
    for i in 1..=n {
      if nums[i - 1] != i as i32 {
        return false;
      }
    }

    nums[n] == n as i32
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![1,1,2], false)]
  #[test_case(vec![1, 3, 3, 2], true ; "good - permutation of base[3]")]
  #[test_case(vec![2, 1, 3], false ; "not good - wrong length")]
  #[test_case(vec![1, 1], true ; "good - permutation of base[1]")]
  #[test_case(vec![3, 4, 4, 1, 2, 1], false ; "not good - wrong length for base[4]")]
  fn check_array_good(nums: Vec<i32>, expected: bool) {
    assert_eq!(Solution::is_good(nums), expected);
  }
}
