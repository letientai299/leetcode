// https://leetcode.com/problems/greatest-english-letter-in-upper-and-lower-case/
struct Solution {}

impl Solution {
    pub fn greatest_letter(s: String) -> String {
        let mut all = [0; 26];
        let mut best: Option<u8> = None;

        s.chars().for_each(|c| {
            let i = if c.is_ascii_lowercase() {
                all[(c as usize) - 'a' as usize] |= 1;
                (c as u8) - b'a'
            } else {
                all[(c as usize) - 'A' as usize] |= 2;
                (c as u8) - b'A'
            };

            let v = all[i as usize];
            if v == 3 && (best.is_none() || best.is_some_and(|index| index < i))
            {
                best = Some(i)
            }
        });

        match best {
            None => "".to_string(),
            Some(index) => ((index + b'A') as char).to_string(),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn greatest_letter() {
        let tests = [("lEeTcOdE", "E")];

        for tc in tests {
            let (input, want) = tc;
            let actual = Solution::greatest_letter(input.into());
            assert_eq!(
                want, actual,
                "input={:?}, want={:?}, actual={:?}",
                input, want, actual
            );
        }
    }
}
