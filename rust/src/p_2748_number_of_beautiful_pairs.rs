use super::{Solution, util};

impl Solution {
  /// You are given a **0-indexed** integer array `nums`. A pair of indices
  /// `i`, `j` where `0 <= i < j < nums.length` is called **beautiful** if
  /// the **first digit** of `nums[i]` and the **last digit** of `nums[j]`
  /// are **coprime**.
  ///
  /// Return *the total number of beautiful pairs in* `nums`.
  ///
  /// Two integers `x` and `y` are **coprime** if there is no integer greater
  /// than 1 that divides both of them. In other words, `x` and `y` are
  /// coprime if `gcd(x, y) == 1`.
  ///
  /// **Constraints:**
  ///
  /// - `2 <= nums.length <= 100`
  /// - `1 <= nums[i] <= 9999`
  /// - `nums[i] % 10 != 0`
  pub fn count_beautiful_pairs(nums: Vec<i32>) -> i32 {
    let mut n = 0;
    for i in 0..nums.len() {
      for j in i + 1..nums.len() {
        let mut x = nums[i];
        let y = nums[j];
        while x / 10 > 0  {
          x /= 10
        }
        if util::gcd(x, y % 10) == 1 {
          n += 1
        }
      }
    }
    n
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_example1() {
    assert_eq!(Solution::count_beautiful_pairs(vec![2, 5, 1, 4]), 5);
  }

  #[test]
  fn test_example2() {
    assert_eq!(Solution::count_beautiful_pairs(vec![11, 21, 12]), 2);
  }
}
