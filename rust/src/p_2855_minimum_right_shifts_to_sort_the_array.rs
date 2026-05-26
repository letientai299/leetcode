// https://leetcode.com/problems/minimum-right-shifts-to-sort-the-array/description/
struct Solution {}

impl Solution {
  /// You are given a **0-indexed** array `nums` of length `n` containing
  /// **distinct** positive integers. Return *the **minimum** number of
  /// **right shifts** required to sort* `nums` *and* `-1` *if this is not
  /// possible.*
  ///
  /// A **right shift** is defined as shifting the element at index `i` to
  /// index `(i + 1) % n`, for all indices.
  ///
  /// **Constraints:**
  /// - `1 <= nums.length <= 100`
  /// - `1 <= nums[i] <= 100`
  /// - `nums` contain distinct integers.
  pub fn minimum_right_shifts(nums: Vec<i32>) -> i32 {
    // must have only 1 or 0 drop
    let mut drop_at: i32 = -1;
    for i in 1..nums.len() {
      if nums[i] < nums[i - 1] {
        if drop_at != -1 {
          return -1; // already found a drop point
        }
        drop_at = i as i32
      }
    }

    if drop_at == -1 {
      return 0;
    }

    if nums[0] < nums[nums.len() - 1] {
      return -1;
    }

    (nums.len() as i32 - drop_at)
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![3,4,5,1,2], 2 ; "two right shifts")]
  #[test_case(vec![1,3,5], 0 ; "already sorted")]
  #[test_case(vec![2,1,4], -1 ; "impossible")]
  fn min_right_shifts(nums: Vec<i32>, expected: i32) {
    assert_eq!(Solution::minimum_right_shifts(nums), expected);
  }
}
