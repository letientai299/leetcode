// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-i/description/
struct Solution {}

impl Solution {
  /// You are given two strings `s1` and `s2`, both of length **4**,
  /// consisting of **lowercase** English letters.
  ///
  /// You can apply the following operation on any of the two strings **any**
  /// number of times:
  /// - Choose any two indices `i` and `j` such that `j - i = 2`, then
  ///   **swap** the two characters at those indices in the string.
  ///
  /// Return `true` *if you can make the strings equal, and* `false`
  /// *otherwise*.
  ///
  /// **Constraints:**
  /// - `s1.length == s2.length == 4`
  /// - `s1` and `s2` consist only of lowercase English letters.
  pub fn can_be_equal(s1: String, s2: String) -> bool {
    let a = s1.as_bytes();
    let b = s2.as_bytes();

    ((a[0] == b[0] && a[2] == b[2]) || (a[0] == b[2] && a[2] == b[0]))
      && ((a[1] == b[1] && a[3] == b[3]) || (a[1] == b[3] && a[3] == b[1]))
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case("abcd", "cdab", true ; "abcd and cdab")]
  #[test_case("abcd", "dacb", false ; "abcd and dacb")]
  #[test_case("bnxw", "bwxn", true)]
  fn can_be_made_equal(s1: &str, s2: &str, expected: bool) {
    assert_eq!(
      Solution::can_be_equal(s1.to_string(), s2.to_string()),
      expected
    );
  }
}
