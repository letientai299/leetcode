// https://leetcode.com/problems/count-symmetric-integers/description/
struct Solution {}

impl Solution {
  /// You are given two positive integers `low` and `high`.
  ///
  /// An integer `x` consisting of `2 * n` digits is **symmetric** if the sum
  /// of the first `n` digits equals the sum of the last `n` digits. Numbers
  /// with an odd number of digits are **never** symmetric.
  ///
  /// Return *the **number of symmetric** integers in the range*
  /// `[low, high]`.
  ///
  /// **Constraints:**
  /// - `1 <= low <= high <= 10^4`
  pub fn count_symmetric_integers(low: i32, high: i32) -> i32 {
    // x is bound, so it's either 2 or 4 digits.
    // For 2 digits, to be symmetric as specified, both digits are the same,
    // So, we have 9 candidates.
    // For 4 digits, the left side can be 10->99, whose sum goes from 1 to 18.
    // We need to compute the candidates' count of 2 digits for each sum.
    // Let's see:
    // - 1 -> 01, 10.
    // - 2 -> 02, 20, 11
    // - 3 -> 03, 30, 21, 12
    // - 4 -> 04, 40, 13, 31, 22
    // - 5 -> 05, 50, 14, 41, 23, 32
    // Can't prove, also doens't matter for just 18 numbers.
    // But we can see that 1->9 => x -> x+1
    // and 10->18 => x -> 18-x+1

    // If we use the above logic to solve this problem, we will need to deal with more math.
    // So, let's solve it the dump way.
    let mut count = 0;
    for n in (low..=high) {
      if n < 10 || (n >= 100 && n <= 999) {
        continue; // ignore odd count of digits
      }

      if n < 100 {
        if n % 11 == 0 {
          count += 1;
        }
        continue;
      }

      let left = n % 10 + (n % 100) / 10;
      let right = n / 1000 + (n / 100) % 10;
      if (left == right) {
        count += 1;
      }
    }
    count
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(1, 100, 9 ; "1 to 100")]
  #[test_case(1200, 1230, 4 ; "1200 to 1230")]
  fn count_symmetric(low: i32, high: i32, expected: i32) {
    assert_eq!(Solution::count_symmetric_integers(low, high), expected);
  }
}
