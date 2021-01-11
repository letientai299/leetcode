package main

func thousandSeparator(n int) string {
	buf := make([]byte, 0, 16)
	c := 0
	for {
		c++
		buf = append(buf, byte(n%10+'0'))
		n /= 10
		if c == 3 && n > 0 {
			buf = append(buf, '.')
			c = 0
		}
		if n == 0 {
			break
		}
	}

	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}
	return string(buf)
}
