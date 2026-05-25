use std::cmp::min;

// https://leetcode.com/problems/total-distance-traveled/description/
struct Solution {}

impl Solution {
  /// A truck has two fuel tanks. You are given two integers, `main_tank`
  /// representing the fuel present in the main tank in liters and
  /// `additional_tank` representing the fuel present in the additional tank
  /// in liters.
  ///
  /// The truck has a mileage of **10 km per liter**. Whenever **5 liters**
  /// of fuel get used up in the main tank, if the additional tank has at
  /// least 1 liter of fuel, **1 liter** of fuel will be transferred from the
  /// additional tank to the main tank.
  ///
  /// Return *the maximum distance which can be traveled.*
  ///
  /// **Constraints:**
  /// - `1 <= mainTank, additionalTank <= 100`
  pub fn distance_traveled(main: i32, add: i32) -> i32 {
    if main < 5 || add == 0 {
      return main * 10;
    }

    // after the first 5 liters of fuel, main tanks get the first 1 liter, so it only needs to burn
    // 4 liters more to get another 1 liter.
    (min((main - 5) / 4, add - 1) + 1 + main) * 10
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(9, 2, 110)]
  #[test_case(5, 10, 60 ; "main=5 additional=10")]
  #[test_case(1, 2, 10 ; "main=1 additional=2")]
  fn total_distance(main_tank: i32, additional_tank: i32, expected: i32) {
    assert_eq!(
      Solution::distance_traveled(main_tank, additional_tank),
      expected
    );
  }
}
