// https://leetcode.com/problems/find-the-maximum-achievable-number/description/
struct Solution {}

impl Solution {
  /// Given two integers, `num` and `t`. A number `x` is **achievable** if it
  /// can become equal to `num` after applying the following operation **at
  /// most** `t` times:
  /// - Increase or decrease `x` by `1`, and *simultaneously* increase or
  ///   decrease `num` by `1`.
  ///
  /// Return the **maximum** possible value of `x`.
  ///
  /// **Constraints:**
  /// - `1 <= num, t <= 50`
  pub fn the_maximum_achievable_x(num: i32, t: i32) -> i32 {
    num + 2 * t
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(4, 1, 6 ; "num=4 t=1")]
  #[test_case(3, 2, 7 ; "num=3 t=2")]
  fn max_achievable(num: i32, t: i32, expected: i32) {
    assert_eq!(Solution::the_maximum_achievable_x(num, t), expected);
  }
}
