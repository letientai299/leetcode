// https://leetcode.com/problems/split-strings-by-separator/description/
struct Solution {}

impl Solution {
  /// Given an array of strings `words` and a character `separator`, **split**
  /// each string in `words` by `separator`.
  ///
  /// Return *an array of strings containing the new strings formed after the
  /// splits, **excluding empty strings**.*
  ///
  /// **Constraints:**
  /// - `1 <= words.length <= 100`
  /// - `1 <= words[i].length <= 20`
  /// - `separator` is a character from `".,|$#@"`
  pub fn split_words_by_separator(
    words: Vec<String>,
    separator: char,
  ) -> Vec<String> {
    words
      .iter() // work
      // .into_iter() // not working
      .flat_map(|w| w.split(separator))
      .filter(|s| !s.is_empty())
      .map(String::from)
      .collect()
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  fn s(v: &[&str]) -> Vec<String> {
    v.iter().map(|s| s.to_string()).collect()
  }

  #[test_case(
    &["one.two.three", "four.five", "six"], '.',
    &["one", "two", "three", "four", "five", "six"]
    ; "dot separator"
  )]
  #[test_case(
    &["$easy$", "$problem$"], '$',
    &["easy", "problem"]
    ; "dollar separator"
  )]
  #[test_case(
    &["|||"], '|',
    &[] as &[&str]
    ; "all separators"
  )]
  fn split_strings(words: &[&str], separator: char, expected: &[&str]) {
    assert_eq!(
      Solution::split_words_by_separator(s(words), separator),
      s(expected)
    );
  }
}
