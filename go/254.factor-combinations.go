package main

// Factor Combinations
//
// Medium
//
// https://leetcode.com/problems/factor-combinations/
func getFactors(n int) [][]int {
	memo := make(map[int][][]int)

	var fac func(n int) [][]int
	fac = func(n int) [][]int {
		var r [][]int

		for i := 2; i <= n/i; i++ {
			if n%i == 0 {
				sub := n / i
				r = append(r, []int{i, sub})
				factors, ok := memo[sub]
				if !ok {
					factors = getFactors(sub)
					memo[sub] = factors
				}

				for _, fs := range factors {
					if len(fs) == 0 || fs[0] < i {
						continue
					}

					fs = append([]int{i}, fs...)
					r = append(r, fs)
				}
			}
		}
		return r
	}

	return fac(n)
}

// TODO (tai): can be faster
