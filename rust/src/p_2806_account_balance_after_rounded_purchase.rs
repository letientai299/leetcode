// https://leetcode.com/problems/account-balance-after-rounded-purchase/description/
struct Solution {}

impl Solution {
  /// Initially, you have a bank account balance of **100** dollars.
  ///
  /// You are given an integer `purchase_amount` representing the amount you
  /// will spend. The `purchase_amount` is **rounded to the nearest multiple
  /// of 10**. Then that rounded amount is removed from your balance.
  ///
  /// Return your final bank account balance after the purchase.
  ///
  /// **Notes:**
  /// - `0` is considered a multiple of 10.
  /// - When rounding, `5` is rounded **upward**.
  ///
  /// **Constraints:**
  /// - `0 <= purchaseAmount <= 100`
  pub fn account_balance_after_purchase(purchase_amount: i32) -> i32 {
    100 - ((purchase_amount + 5) / 10) * 10
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(9, 90 ; "9 rounds to 10")]
  #[test_case(15, 80 ; "15 rounds to 20")]
  #[test_case(10, 90 ; "10 is already multiple of 10")]
  fn account_balance(purchase_amount: i32, expected: i32) {
    assert_eq!(
      Solution::account_balance_after_purchase(purchase_amount),
      expected
    );
  }
}
