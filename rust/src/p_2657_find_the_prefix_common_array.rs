// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/
pub struct Solution();

impl Solution {
    pub fn find_the_prefix_common_array(a: Vec<i32>, b: Vec<i32>) -> Vec<i32> {
        let n = a.len();
        let mut seen = vec![0; n];
        let mut res = vec![0; n + 1];

        for i in 0..n {
            let x = a[i];
            let y = b[i];
            seen[x as usize - 1] |= 1;
            seen[y as usize - 1] |= 2;

            res[i + 1] = res[i]
                + if seen[x as usize - 1] == 3 { 1 } else { 0 }
                + if seen[y as usize - 1] == 3 { 1 } else { 0 }
                - if x == y { 1 } else { 0 }
        }

        res[1..].to_vec()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::util::vec_of;
    use test_case::test_case;

    #[test]
    fn test1() {
        let a = vec_of::<i32>("[1,3,2,4]");
        let b = vec_of::<i32>("[3,1,2,4]");
        let want = vec_of::<i32>("[0,2,3,4]");
        let actual =
            Solution::find_the_prefix_common_array(a.clone(), b.clone());

        assert_eq!(
            want, actual,
            "a={:?}, b={:?}, want={:?}, actual={:?}",
            a, b, want, actual
        );
    }

    #[test_case("[1,3,2,4]", "[3,1,2,4]", "[0, 2,3,4]")]
    #[test_case("[2,3,1]", "[3,1,2]", "[0,1,3]")]
    fn test(raw_a: &str, raw_b: &str, raw_want: &str) {
        let a = vec_of::<i32>(raw_a);
        let b = vec_of::<i32>(raw_b);
        let want = vec_of::<i32>(raw_want);
        let actual =
            Solution::find_the_prefix_common_array(a.clone(), b.clone());

        assert_eq!(
            want, actual,
            "a={:?}, b={:?}, want={:?}, actual={:?}",
            a, b, want, actual
        );
    }
}
