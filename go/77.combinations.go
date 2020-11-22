package main

// medium

func combine(n int, k int) [][]int {
	all := combinator(n, k)
	res := make([][]int, 0, all)
	buf := make([]int, 0, k)
	for i := 1; i <= k; i++ {
		buf = append(buf, i)
	}

	for all > 0 {
		all--
		t := make([]int, k)
		copy(t, buf)
		res = append(res, t)

		for j := k - 1; j >= 0; j-- {
			if (j == k-1 && buf[j] == n) || j < k-1 && buf[j] == buf[j+1]-1 {
				continue
			}
			buf[j]++
			for l := j + 1; l < k; l++ {
				buf[l] = buf[l-1] + 1
			}
			break
		}
	}
	return res
}

func combinator(m int, n int) int {
	v := uint64(1)
	for i := m; i > m-n; i-- {
		v *= uint64(i)
	}
	for i := 2; i <= n; i++ {
		v /= uint64(i)
	}
	return int(v)
}
