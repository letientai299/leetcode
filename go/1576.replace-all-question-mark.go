package main

func modifyString(s string) string {
	buf := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != '?' {
			buf = append(buf, c)
			continue
		}

		c = 'a'
		for (i > 0 && buf[i-1] == c) || (i < len(s)-1 && s[i+1] == c) {
			c++
		}
		buf = append(buf, c)
	}
	return string(buf)
}
