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
    let s = &s[1..(s.len() - 1)];
    let v = s.split(',').map(|sub| sub.trim().parse::<T>()).collect();

    if let Ok(x) = v {
        x
    } else {
        panic!("Fail to parse {}", s)
    }
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
    let s = &s.trim()[1..(s.len() - 1)];
    let bs = s.as_bytes();
    let mut last = 0;
    let mut res: Vec2<T> = vec![];
    for i in 0..bs.len() {
        if bs[i] == b'[' {
            last = i;
            continue;
        }

        if bs[i] == b']' {
            res.push(vec_of::<T>(&s[last..i + 1]));
        }
    }

    res
}
