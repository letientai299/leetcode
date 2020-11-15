package main

func makeGood(s string) string {
	if s == "" {
		return s
	}

	const diff = 'a' - 'A'

	bad := func(a, b byte) bool {
		return a-b == diff || b-a == diff
	}

	buf := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		x := s[i]
		if len(buf) == 0 {
			buf = append(buf, x)
			continue
		}

		last := buf[len(buf)-1]
		if bad(x, last) {
			buf = buf[:len(buf)-1]
			continue
		}

		buf = append(buf, x)
	}

	return string(buf)
}
