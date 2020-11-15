package main

func sortString(s string) string {
	var m [26]int
	for _, x := range s {
		m[x-'a']++
	}
	res := make([]byte, 0, len(s))

	for len(res) < len(s) {
		for i, v := range m {
			if v != 0 {
				res = append(res, byte(i+'a'))
				m[i]--
			}
		}

		for i := range m {
			x := len(m) - 1 - i
			if m[x] != 0 {
				res = append(res, byte(x+'a'))
				m[x]--
			}
		}
	}
	return string(res)
}
