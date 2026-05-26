use std::fmt::Display;
use std::io::repeat;

// https://leetcode.com/problems/maximum-odd-binary-number/description/
struct Solution {}

impl Solution {
  /// You are given a **binary** string `s` that contains at least one `'1'`.
  ///
  /// You have to **rearrange** the bits in such a way that the resulting
  /// binary number is the **maximum odd binary number** that can be created
  /// from this combination.
  ///
  /// Return *a string representing the maximum odd binary number that can be
  /// created from the given combination.*
  ///
  /// **Note** that the resulting string **can** have leading zeros.
  ///
  /// **Constraints:**
  /// - `1 <= s.length <= 100`
  /// - `s` consists only of `'0'` and `'1'`.
  /// - `s` contains at least one `'1'`.
  pub fn maximum_odd_binary_number(s: String) -> String {
    let ones = s.chars().filter(|&c| c == '1').count();
    let zeroes = s.len() - ones;
    format!("{}{}1", "1".repeat(ones - 1), "0".repeat(zeroes))
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case("010", "001" ; "single one")]
  #[test_case("0101", "1001" ; "two ones")]
  fn max_odd_binary(s: &str, expected: &str) {
    assert_eq!(Solution::maximum_odd_binary_number(s.to_string()), expected);
  }
}
