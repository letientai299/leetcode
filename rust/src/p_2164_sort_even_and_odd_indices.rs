// https://leetcode.com/problems/sort-even-and-odd-indices-independently/description/
struct Solution {}

impl Solution {
  pub fn sort_even_odd(nums: Vec<i32>) -> Vec<i32> {
    let mut odd: Vec<i32> = vec![];
    let mut even: Vec<i32> = vec![];
    for (i, x) in nums.iter().enumerate() {
      if i % 2 == 0 {
        even.push(*x);
      } else {
        odd.push(*x)
      }
    }

    even.sort_unstable();
    odd.sort_unstable_by(|a, b| b.cmp(a));

    // merge
    let mut nums = nums;
    for i in 0..nums.len() {
      nums[i] = if i % 2 == 0 { even[i / 2] } else { odd[i / 2] }
    }

    nums
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn sort_even_and_odd() {
    let tests = [
      (
        vec![
          36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34,
        ],
        vec![
          9, 46, 15, 45, 15, 41, 27, 34, 32, 31, 33, 31, 36, 26, 36, 16, 44, 6,
        ],
      ),
      (vec![4, 1, 2, 3], vec![2, 3, 4, 1]),
      (vec![4, 1, 3], vec![3, 1, 4]),
      (vec![2, 1], vec![2, 1]),
    ];

    for tc in tests {
      let (input, want) = tc;
      let actual = Solution::sort_even_odd(input.clone());
      assert_eq!(
        want, actual,
        "input={:?}, want={:?}, actual={:?}",
        input, want, actual
      );
    }
  }
}
