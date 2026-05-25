use std::str::FromStr;

/// Returns the parsed vector from input string.
///
/// # Examples
///
/// ```
/// # use crate::lc_rust::util::vec_of;
/// assert_eq!(vec_of::<i32>("[1, 2, 3]"), vec![1,2,3]);
/// assert_eq!(vec_of::<f32>("[1, 2, 3]"), vec![1., 2., 3.]);
/// assert_eq!(vec_of::<String>("[abc, def]"), vec!["abc", "def"]);
/// ```
pub fn vec_of<T: FromStr>(s: &str) -> Vec<T> {
    let s = &s[1..s.len() - 1];
    let mut res = Vec::new();
    for sub in s.split(',') {
        match sub.trim().parse::<T>() {
            Ok(v) => res.push(v),
            Err(_) => panic!("Failed to parse '{}'", s),
        }
    }
    res
}

type Vec2<T> = Vec<Vec<T>>;

/// Returns the parsed 2d vector from input string.
///
/// # Examples
///
/// ```
/// # use crate::lc_rust::util::vec2d_of;
/// assert_eq!(vec2d_of::<i32>("[[1,2],[3,4]]"), vec![vec![1,2], vec![3,4]]);
/// assert_eq!(vec2d_of::<String>("[[ab, cd], [e, f]]"), vec![vec!["ab","cd"], vec!["e","f"]]);
/// ```
pub fn vec2d_of<T: FromStr>(s: &str) -> Vec2<T> {
    let s = &s.trim()[1..s.len() - 1];
    let bs = s.as_bytes();
    let mut res: Vec2<T> = Vec::new();
    let mut last = 0;
    for (i, &b) in bs.iter().enumerate() {
        if b == b'[' {
            last = i;
        } else if b == b']' {
            res.push(vec_of::<T>(&s[last..=i]));
        }
    }
    res
}

#[cfg(test)]
mod tests {
    use super::*;
    use test_case::test_case;

    #[test_case("[1,2,3]", vec![1, 2, 3]; "no_spaces")]
    #[test_case("[1, 2, 3]", vec![1, 2, 3]; "with_spaces")]
    #[test_case("[42]", vec![42]; "single")]
    #[test_case("[-1, 0, 1]", vec![-1, 0, 1]; "negative")]
    fn vec_of_i32(input: &str, want: Vec<i32>) {
        assert_eq!(vec_of::<i32>(input), want);
    }

    #[test_case("[1.5, 2.0]", vec![1.5, 2.0])]
    fn vec_of_f64(input: &str, want: Vec<f64>) {
        assert_eq!(vec_of::<f64>(input), want);
    }

    #[test_case("[abc, def]", vec!["abc", "def"])]
    fn vec_of_string(input: &str, want: Vec<&str>) {
        assert_eq!(vec_of::<String>(input), want);
    }

    #[test_case("[[1,2],[3,4]]", vec![vec![1,2], vec![3,4]]; "no_spaces")]
    #[test_case("[[1, 2], [3, 4]]", vec![vec![1,2], vec![3,4]]; "with_spaces")]
    #[test_case("[[5]]", vec![vec![5]]; "single")]
    #[test_case("[[1,2,3],[4],[5,6]]", vec![vec![1,2,3], vec![4], vec![5,6]]; "ragged")]
    fn vec2d_of_i32(input: &str, want: Vec<Vec<i32>>) {
        assert_eq!(vec2d_of::<i32>(input), want);
    }

    #[test_case("[[ab, cd], [e, f]]", vec![vec!["ab","cd"], vec!["e","f"]])]
    fn vec2d_of_string(input: &str, want: Vec<Vec<&str>>) {
        assert_eq!(vec2d_of::<String>(input), want);
    }
}

#[cfg(feature = "bench")]
mod benches {
    use super::*;

    struct Named(&'static str, &'static str);

    impl std::fmt::Display for Named {
        fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
            f.write_str(self.0)
        }
    }

    const VEC2D_INPUTS: [Named; 4] = [
        Named("2x2", "[[1,2],[3,4]]"),
        Named("3_ragged", "[[1,2,3],[4],[5,6]]"),
        Named("3x2", "[[3,50],[7,10],[12,25]]"),
        Named("7x2", "[[2,7],[3,17],[4,37],[7,6],[9,83],[16,67],[19,29]]"),
    ];

    const VEC1D_INPUTS: [Named; 3] = [
        Named("3_elem", "[1,2,3]"),
        Named("3_spaced", "[1, 2, 3]"),
        Named("7_elem", "[2,7,3,17,4,37,7]"),
    ];

    #[divan::bench(args = VEC1D_INPUTS)]
    fn vec_of_bench(input: &Named) -> Vec<i32> {
        vec_of::<i32>(divan::black_box(input.1))
    }

    #[divan::bench(args = VEC2D_INPUTS)]
    fn vec2d_of_bench(input: &Named) -> Vec<Vec<i32>> {
        vec2d_of::<i32>(divan::black_box(input.1))
    }
}
