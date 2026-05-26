// https://leetcode.com/problems/minimum-operations-to-collect-elements/description/
struct Solution {}

impl Solution {
  /// You are given an array `nums` of positive integers and an integer `k`.
  ///
  /// In one operation, you can **remove the last element** of the array and
  /// add it to your collection.
  ///
  /// Return *the **minimum number of operations** needed to collect
  /// elements* `1, 2, ..., k`.
  ///
  /// **Constraints:**
  /// - `1 <= nums.length <= 50`
  /// - `1 <= nums[i] <= nums.length`
  /// - `1 <= k <= nums.length`
  pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
    let mut seen = vec![false; k as usize];
    let mut rem = k;
    for i in (0..nums.len()).rev() {
      let x = nums[i] as usize;
      if x > k as usize || seen[x - 1] {
        continue;
      }

      rem -= 1;
      seen[x - 1] = true;
      if rem == 0 {
        return (nums.len() - i) as i32;
      }
    }

    nums.len() as i32
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![3,1,5,4,2], 2, 4 ; "k=2")]
  #[test_case(vec![3,1,5,4,2], 5, 5 ; "k=5 collect all")]
  #[test_case(vec![3,2,5,3,1], 3, 4 ; "k=3")]
  fn min_ops_collect(nums: Vec<i32>, k: i32, expected: i32) {
    assert_eq!(Solution::min_operations(nums, k), expected);
  }
}
