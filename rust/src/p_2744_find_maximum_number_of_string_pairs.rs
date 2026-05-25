use crate::Solution;
use std::collections::HashSet;

/// # 2744. Find Maximum Number of String Pairs
///
/// You are given a **0-indexed** array `words` consisting of **distinct**
/// strings.
///
/// The string `words[i]` can be paired with the string `words[j]` if:
///
/// - The string `words[i]` is equal to the **reversed** string of `words[j]`.
/// - `0 <= i < j < words.length`
///
/// Return the **maximum** number of pairs that can be formed from the array
/// `words`.
///
/// Note that each string can belong in **at most one** pair.
///
/// ## Constraints
///
/// - `1 <= words.length <= 50`
/// - `words[i].length == 2`
/// - `words` consists of distinct strings.
/// - `words[i]` contains only lowercase English letters.
///
/// ## Examples
///
/// ```
/// # use leetcode::p_2744_find_maximum_number_of_string_pairs::maximum_number_of_string_pairs;
/// assert_eq!(maximum_number_of_string_pairs(
///     vec!["cd","ac","dc","ca","zz"].into_iter().map(String::from).collect()
/// ), 2);
/// assert_eq!(maximum_number_of_string_pairs(
///     vec!["ab","ba","cc"].into_iter().map(String::from).collect()
/// ), 1);
/// assert_eq!(maximum_number_of_string_pairs(
///     vec!["aa","ab"].into_iter().map(String::from).collect()
/// ), 0);
/// ```
impl Solution {
  pub fn maximum_number_of_string_pairs(words: Vec<String>) -> i32 {
    let mut n = 0;
    let mut set = HashSet::<String>::new();
    for w in words {
      if set.contains(&w) {
        n += 1;
      } else {
        set.insert(w.chars().rev().collect());
      }
    }
    n
  }
}

#[cfg(test)]
mod tests {
  use crate::Solution;

  #[test]
  fn test_example_1() {
    let words: Vec<String> = vec!["cd", "ac", "dc", "ca", "zz"]
      .into_iter()
      .map(String::from)
      .collect();
    assert_eq!(Solution::maximum_number_of_string_pairs(words), 2);
  }

  #[test]
  fn test_example_2() {
    let words: Vec<String> = vec!["ab", "ba", "cc"]
      .into_iter()
      .map(String::from)
      .collect();
    assert_eq!(Solution::maximum_number_of_string_pairs(words), 1);
  }

  #[test]
  fn test_example_3() {
    let words: Vec<String> =
      vec!["aa", "ab"].into_iter().map(String::from).collect();
    assert_eq!(Solution::maximum_number_of_string_pairs(words), 0);
  }
}
