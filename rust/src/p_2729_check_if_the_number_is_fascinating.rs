// https://leetcode.com/problems/check-if-the-number-is-fascinating/description/
struct Solution {}

impl Solution {
  /// You are given an integer `n` that consists of exactly **3 digits**.
  ///
  /// We call the number `n` **fascinating** if, after the following
  /// modification, the resulting number contains all the digits from `1` to
  /// `9` **exactly** once and does not contain any `0`s:
  /// - **Concatenate** `n` with the numbers `2 * n` and `3 * n`.
  ///
  /// Return `true` if `n` is fascinating, or `false` otherwise.
  ///
  /// **Constraints:**
  /// - `100 <= n <= 999`
  pub fn is_fascinating(n: i32) -> bool {
    let n2 = n * 2;
    let n3 = n * 3;
    if n % 10 == 5 || n2 >= 999 || n3 >= 999 {
      return false;
    }

    let mut digits = [0; 10];
    let mut check = |mut x: i32| {
      while x > 0 {
        let d = (x % 10) as usize;
        digits[d] += 1;
        x /= 10;
      }
    };

    check(n);
    check(n * 2);
    check(n * 3);

    digits[0] == 0 && digits[1..].iter().all(|&d| d == 1)
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(192, true ; "192 is fascinating")]
  #[test_case(100, false ; "100 is not fascinating")]
  fn check_if_fascinating(n: i32, expected: bool) {
    assert_eq!(Solution::is_fascinating(n), expected);
  }
}
