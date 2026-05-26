// https://leetcode.com/problems/furthest-point-from-origin/description/
struct Solution {}

impl Solution {
  /// You are given a string `moves` of length `n` consisting only of
  /// characters `'L'`, `'R'`, and `'_'`. The string represents your movement
  /// on a number line starting from the origin `0`.
  ///
  /// For each move:
  /// - move **left** if `moves[i] = 'L'` or `moves[i] = '_'`
  /// - move **right** if `moves[i] = 'R'` or `moves[i] = '_'`
  ///
  /// Return *the **distance from the origin** of the **furthest** point you
  /// can get to after* `n` *moves*.
  ///
  /// **Constraints:**
  /// - `1 <= moves.length == n <= 50`
  /// - `moves` consists only of characters `'L'`, `'R'` and `'_'`.
  pub fn furthest_distance_from_origin(moves: String) -> i32 {
    let (mut pos, mut none): (i32, i32) = (0, 0);
    for m in moves.chars() {
      match m {
        'L' => pos += 1,
        'R' => pos -= 1,
        _ => none += 1,
      }
    }
    pos.abs() + none
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case("L_RL__R", 3 ; "L_RL__R")]
  #[test_case("_R__LL_", 5 ; "_R__LL_")]
  #[test_case("_______", 7 ; "all underscores")]
  fn furthest_point(moves: &str, expected: i32) {
    assert_eq!(
      Solution::furthest_distance_from_origin(moves.to_string()),
      expected
    );
  }
}
