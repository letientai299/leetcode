package main

// medium

func findStrobogrammatic(n int) []string {
	if n == 0 {
		return nil
	}

	if n == 1 {
		return []string{"1", "8", "0"}
	}

	buf := make([]byte, n)
	res := make([]string, 0, n*2)

	var gen func(i int)

	gen = func(i int) {
		if i >= n/2 {
			if n%2 == 1 {
				buf[i] = '0'
				res = append(res, string(buf))
				buf[i] = '1'
				res = append(res, string(buf))
				buf[i] = '8'
				res = append(res, string(buf))
				return
			}
			res = append(res, string(buf))
			return
		}

		if i != 0 {
			buf[i] = '0'
			buf[n-1-i] = '0'
			gen(i + 1)
		}

		buf[i] = '1'
		buf[n-1-i] = '1'
		gen(i + 1)

		buf[i] = '8'
		buf[n-1-i] = '8'
		gen(i + 1)

		buf[i] = '6'
		buf[n-1-i] = '9'
		gen(i + 1)

		buf[i] = '9'
		buf[n-1-i] = '6'
		gen(i + 1)
	}

	gen(0)
	return res
}
