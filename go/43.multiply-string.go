package main

// medium
// https://leetcode.com/problems/multiply-strings/

func multiply(a string, b string) string {
	temp := make([][]byte, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		buf := make([]byte, 0, len(b)+len(a))
		rem := uint8(0)
		x := a[i] - '0'

		for n := len(a) - 1 - i; n > 0; n-- {
			buf = append(buf, 0)
		}

		for j := len(b) - 1; j >= 0; j-- {
			y := b[j] - '0'
			z := x*y + rem
			rem = z / 10
			buf = append(buf, z%10)
		}

		n := i + 1
		if rem > 0 {
			n--
			buf = append(buf, rem)
		}

		for ; n > 0; n-- {
			buf = append(buf, 0)
		}

		temp = append(temp, buf)
	}

	res := make([]byte, len(b)+len(a))
	rem := 0
	m := len(a) + len(b)
	for i := 0; i < m; i++ {
		for _, s := range temp {
			rem += int(s[i])
		}
		res[m-1-i] = byte(rem%10) + '0'
		rem /= 10
	}

	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}

	return string(res)
}

