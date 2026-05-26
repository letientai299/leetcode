// https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/description/
struct Solution {}

impl Solution {
  /// Given an array of strings `words` and a string `s`, determine if `s` is
  /// an **acronym** of `words`.
  ///
  /// The string `s` is considered an acronym of `words` if it can be formed
  /// by concatenating the **first** character of each string in `words` **in
  /// order**.
  ///
  /// Return `true` *if* `s` *is an acronym of* `words`, *and* `false`
  /// *otherwise.*
  ///
  /// **Constraints:**
  /// - `1 <= words.length <= 100`
  /// - `1 <= words[i].length <= 10`
  /// - `1 <= s.length <= 100`
  pub fn is_acronym(words: Vec<String>, s: String) -> bool {
    s.chars().count() == words.len()
      && words
        .iter()
        .zip(s.chars())
        .all(|(w, c)| w.chars().next() == Some(c))
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  fn s(v: &[&str]) -> Vec<String> {
    v.iter().map(|s| s.to_string()).collect()
  }

  #[test_case(&["alice","bob","charlie"], "abc", true ; "abc is acronym")]
  #[test_case(&["an","apple"], "a", false ; "a is not acronym of two words")]
  #[test_case(&["never","gonna","give","up","on","you"], "ngguoy", true ; "ngguoy")]
  fn check_acronym(words: &[&str], acronym: &str, expected: bool) {
    assert_eq!(
      Solution::is_acronym(s(words), acronym.to_string()),
      expected
    );
  }
}
