// https://leetcode.com/problems/faulty-keyboard/description/
struct Solution {}

impl Solution {
  /// Your laptop keyboard is faulty, and whenever you type a character `'i'`
  /// on it, it **reverses** the string that you have written. Typing other
  /// characters works as expected.
  ///
  /// You are given a **0-indexed** string `s`, and you type each character
  /// of `s` using your faulty keyboard.
  ///
  /// Return *the final string that will be present on your laptop screen.*
  ///
  /// **Constraints:**
  /// - `1 <= s.length <= 100`
  /// - `s` consists of lowercase English letters.
  /// - `s[0] != 'i'`
  pub fn final_string(s: String) -> String {
    s.replace("ii", "")
      .chars()
      .fold(vec![], |mut acc: Vec<char>, c| match c {
        'i' => acc.into_iter().rev().collect(),
        _ => {
          acc.push(c);
          acc
        }
      })
      .into_iter()
      .collect()
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case("string", "rtsng" ; "string")]
  #[test_case("poiinter", "ponter" ; "poiinter")]
  fn faulty_keyboard(s: &str, expected: &str) {
    assert_eq!(Solution::final_string(s.to_string()), expected);
  }
}
