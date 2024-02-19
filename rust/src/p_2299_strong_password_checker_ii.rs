// https://leetcode.com/problems/strong-password-checker-ii/
struct Solution {}

impl Solution {
  pub fn strong_password_checker_ii(pw: String) -> bool {
    if pw.len() < 8 {
      return false;
    }

    let mut lower = false;
    let mut upper = false;
    let mut digit = false;
    let mut special = false;
    let chars: Vec<char> = pw.chars().collect();
    for (i, b) in chars.iter().enumerate() {
      if i > 0 && *b == chars[i - 1] {
        return false;
      }

      match b {
        'a'..='z' => lower = true,
        'A'..='Z' => upper = true,
        '0'..='9' => digit = true,
        _ => {
          if "!@#$%^&*()-+".contains(*b) {
            special = true
          }
        }
      }
    }

    lower && upper && digit && special
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn strong_password_checker_ii() {
    let tests = [("IloveLe3tcode!", true)];

    for tc in tests {
      let (input, want) = tc;
      let actual = Solution::strong_password_checker_ii(input.into());
      assert_eq!(
        want, actual,
        "input={:?}, want={:?}, actual={:?}",
        input, want, actual
      );
    }
  }
}
