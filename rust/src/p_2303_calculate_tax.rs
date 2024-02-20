struct Solution();

use std::cmp::min;

impl Solution {
    pub fn calculate_tax(brackets: Vec<Vec<i32>>, income: i32) -> f64 {
        brackets
            .iter()
            .fold((0.0, income, 0), |(tax, remain, last_cap), bracket| {
                let cap = bracket[0];
                let (effective_cap, percent) = (cap - last_cap, bracket[1]);
                let charged = min(remain, effective_cap);
                let more_tax = (charged * percent) as f64 / 100.0;
                let remain = remain - charged;
                (tax + more_tax, remain, cap)
            })
            .0
    }
}

#[cfg(test)]
mod tests {
    use crate::p_2303_calculate_tax::Solution;
    use crate::util::vec2d_of;
    use test_case::test_case;

    #[test_case("[[3,50],[7,10],[12,25]]", 10, 2.65)]
    #[test_case("[[2,50]]", 0, 0.0)]
    #[test_case("[[1,0],[4,25],[5,50]]", 2, 0.25)]
    #[test_case("[[2,7],[3,17],[4,37],[7,6],[9,83],[16,67],[19,29]]", 18, 7.79)]
    fn test(brackets_raw: &str, income: i32, want: f64) {
        let brackets = vec2d_of::<i32>(brackets_raw);
        let actual = Solution::calculate_tax(brackets.clone(), income);

        assert!(
            (want - actual).abs() < 0.00001,
            "input={:?}, want={:?}, actual={:?}",
            brackets,
            want,
            actual
        );
    }
}
