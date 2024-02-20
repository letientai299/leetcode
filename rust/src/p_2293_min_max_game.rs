// https://leetcode.com/problems/min-max-game/description/
struct Solution {}

impl Solution {
    pub fn min_max_game(nums: Vec<i32>) -> i32 {
        let mut nums = nums;
        while nums.len() > 1 {
            nums = nums
                .chunks(2)
                .enumerate()
                .map(|(i, pair)| match i % 2 {
                    0 => *pair.iter().min().unwrap(),
                    _ => *pair.iter().max().unwrap(),
                })
                .collect();
        }

        nums[0]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn min_max_game() {
        let tests = [
            (vec![70, 38, 21, 22], 22),
            (vec![1, 3, 5, 2, 4, 8, 2, 2], 1),
            (vec![3], 3),
        ];

        for tc in tests {
            let (input, want) = tc;
            let actual = Solution::min_max_game(input.clone());
            assert_eq!(
                want, actual,
                "input={:?}, want={:?}, actual={:?}",
                input, want, actual
            );
        }
    }
}
