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
    s.split(',')
        .map(|sub| sub.trim().parse::<T>())
        .collect::<Result<Vec<T>, _>>()
        .unwrap_or_else(|_| panic!("Failed to parse '{}'", s))
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
    let s = s.trim();
    s[1..s.len() - 1]
        .split("],")
        .map(|sub| {
            let inner = sub.trim().trim_matches(|c| c == '[' || c == ']');
            vec_of::<T>(&format!("[{inner}]"))
        })
        .collect()
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
