// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
struct Solution {}

impl Solution {
    pub fn count_asterisks(s: String) -> i32 {
        let mut vert_cnt = 0;
        let mut cnt = 0;
        for c in s.chars() {
            if '|' == c {
                vert_cnt = (vert_cnt + 1) % 2;
                continue;
            }
            if '*' == c && vert_cnt == 0 {
                cnt += 1;
            }
        }

        cnt
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test_case("l|*e*et|c**o|*de|", 2)]
    #[test_case("iamprogrammer", 0)]
    #[test_case("yo|uar|e**|b|e***au|tifu|l", 5)]
    fn count_asterisks(input: &str, want: i32) {
        let actual = Solution::count_asterisks(input.into());
        assert_eq!(
            want, actual,
            "input={:?}, want={:?}, actual={:?}",
            input, want, actual
        );
    }
}
