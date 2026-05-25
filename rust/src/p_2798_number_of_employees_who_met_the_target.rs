// https://leetcode.com/problems/number-of-employees-who-met-the-target/description/
struct Solution {}

impl Solution {
  /// There are `n` employees in a company. Each employee `i` has worked for
  /// `hours[i]` hours. The company requires each employee to work for **at
  /// least** `target` hours.
  ///
  /// Return *the number of employees who worked at least* `target` *hours*.
  ///
  /// **Constraints:**
  /// - `1 <= n == hours.length <= 50`
  /// - `0 <= hours[i], target <= 10^5`
  pub fn number_of_employees_who_met_target(
    hours: Vec<i32>,
    target: i32,
  ) -> i32 {
    hours.into_iter().filter(|&h| h >= target).count() as i32
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(vec![0,1,2,3,4], 2, 3 ; "target=2")]
  #[test_case(vec![5,1,4,2,2], 6, 0 ; "target=6 none met")]
  fn employees_met_target(hours: Vec<i32>, target: i32, expected: i32) {
    assert_eq!(
      Solution::number_of_employees_who_met_target(hours, target),
      expected
    );
  }
}
