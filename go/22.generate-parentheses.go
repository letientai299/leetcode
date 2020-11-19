package main

// medium

/*
How many string can we generate from n pairs?
         n       f(n)
         0          0
         1          1
         2          2
         3          5
         4         14
         5         42
         6        132
         7        429
         8       1430
         9       4862
        10      16796
        11      58786
        12     208012
        13     742900
        14    2674440
        15    9694845
        16   35357670
        17  129644790

The formula is not easy to find out.
https://stackoverflow.com/a/26048186/3869533
Proofs: https://en.wikipedia.org/wiki/Catalan_number#Applications_in_combinatorics
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	var gen func(open, close int, buf []byte)
	res := make([]string, 0, 2*n)
	gen = func(open, close int, buf []byte) {
		if open == 0 {
			for close > 0 {
				buf = append(buf, ')')
				close--
			}
			res = append(res, string(buf))
			return
		}

		l := len(buf)
		buf = append(buf, '(')
		gen(open-1, close, buf)
		buf = buf[:l]
		if close > open {
			buf = append(buf, ')')
			gen(open, close-1, buf)
			buf = buf[:l]
		}
	}

	gen(n, n, make([]byte, 0, n))
	return res
}
