// TODO: could be shorter
// https://leetcode.com/problems/dot-product-of-two-sparse-vectors/description/
//
// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     0 <= nums1[i], nums2[i] <= 100

struct SparseVector {
    ns: Vec<i32>,
}

impl SparseVector {
    fn new(nums: Vec<i32>) -> Self {
        let mut ns = vec![];
        let mut zeros = 0;
        for x in nums {
            if x == 0 {
                zeros += 1;
                continue;
            }

            if zeros != 0 {
                // use negative to store consecutive 0, as vector element is >= 0
                ns.push(-zeros);
                zeros = 0;
            }

            ns.push(x);
        }

        if zeros != 0 {
            ns.push(-zeros);
        }

        Self { ns }
    }

    // Return the dotProduct of two sparse vectors
    fn dot_product(&self, other: SparseVector) -> i32 {
        let mut prod = 0;

        let v1 = &self.ns;
        let v2 = other.ns;

        let mut i = 0;
        let mut j = 0;

        // count number of 0 different between v1 and v2.
        // - 0: no different, can process to next element in both vectors.
        // - neg: v1 has more zeroes waiting.
        // - pos: v2 has more zeroes waiting.
        let mut zeroes = 0;

        while i < v1.len() && j < v2.len() {
            if zeroes == 0 {
                let x = v1[i];
                let y = v2[j];
                i += 1;
                j += 1;

                if x > 0 && y > 0 {
                    prod += x * y;
                    continue;
                }

                if x < 0 && y < 0 {
                    zeroes = x - y;
                    continue;
                }

                if x < 0 {
                    zeroes = x + 1; // skip y, store less than 1 zeros for v1 side (thus negative)
                    continue;
                }

                zeroes = -1 - y;
                continue;
            }

            while zeroes > 0 {
                if v1[i] > 0 {
                    i += 1;
                    zeroes -= 1;
                    continue;
                }

                zeroes += v1[i];
                i += 1;
            }

            while zeroes < 0 {
                if v2[j] > 0 {
                    j += 1;
                    zeroes += 1;
                    continue;
                }

                zeroes -= v2[j];
                j += 1;
            }
        }

        prod
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::util::vec_of;
    use test_case::test_case;

    #[test_case("[1,0,0,2,3]", "[0,3,0,4,0]", 8)]
    #[test_case("[0,1,0,0,0]", "[0,0,0,0,2]", 0)]
    #[test_case("[0,1,0,0,2,0,0]", "[1,0,0,0,3,0,4]", 6)]
    fn dot_product(a: &str, b: &str, want: i32) {
        let v1 = SparseVector::new(vec_of(a));
        let v2 = SparseVector::new(vec_of(b));
        let actual = v1.dot_product(v2);

        assert_eq!(
            want, actual,
            "a={:?}, b={:?} want={:?}, actual={:?}",
            a, b, want, actual
        );
    }
}
