package main

func generateTheString(n int) string {
	if n%2 == 1 {
		return "b" + generateTheString(n-1)
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = 'a'
	}
	return string(buf)
}
