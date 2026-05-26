// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/description/
struct Solution {}

impl Solution {
  /// Given a **0-indexed** integer array `nums` of length `n` and an integer
  /// `target`, return *the number of pairs* `(i, j)` *where* `0 <= i < j <
  /// n` *and* `nums[i] + nums[j] < target`.
  ///
  /// **Constraints:**
  /// - `1 <= nums.length == n <= 50`
  /// - `-50 <= nums[i], target <= 50`
  pub fn count_pairs(mut nums: Vec<i32>, target: i32) -> i32 {
    let mut ans = 0;
    nums.sort();
    for i in 0..nums.len() - 1 {
      let limit = target - nums[i];
      let count = nums[i + 1..].partition_point(|&x| x < limit);
      ans += count;
    }

    ans as i32
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![-10,-6,-8,-9,6,6,-6,-6,-3], -2, 25)]
  #[test_case(vec![-1,1,2,3,1], 2, 3 ; "target=2 three pairs")]
  #[test_case(vec![-6,2,5,-2,-7,-1,3], -2, 10 ; "target=-2 ten pairs")]
  fn count_pairs_less_than_target(nums: Vec<i32>, target: i32, expected: i32) {
    assert_eq!(Solution::count_pairs(nums, target), expected);
  }
}
