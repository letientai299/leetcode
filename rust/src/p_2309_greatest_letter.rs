// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
struct Solution {}

impl Solution {
    pub fn greatest_letter(s: String) -> String {
        let mut alpha = [0; 26];
        let mut best = None;

        for c in s.chars() {
            let (i, v) = Self::consider(&mut alpha, c);
            if v == 3 && (best.is_none() || best.unwrap() < i) {
                best = Some(i)
            }
        }

        best.map_or("".into(), |c| char::from(c + b'A').to_string())
    }
    fn consider(alpha: &mut [i32; 26], c: char) -> (u8, i32) {
        let i = (c.to_ascii_lowercase() as u8 - b'a') as usize;
        alpha[i] |= if c.is_ascii_lowercase() { 1 } else { 2 };
        (i as u8, alpha[i])
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test_case("lEeTcOdE", "E")]
    #[test_case("AbCdEfGhIjK", "")]
    fn greatest_letter(input: &str, want: &str) {
        let actual = Solution::greatest_letter(input.into());
        assert_eq!(
            want, actual,
            "input={:?}, want={:?}, actual={:?}",
            input, want, actual
        );
    }
}
