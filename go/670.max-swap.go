package main

// medium

func maximumSwap(n int) int {
	buf := make([]int, 0, 8)
	for n > 0 {
		buf = append(buf, n%10)
		n /= 10
	}

	for i := len(buf) - 1; i >= 1; i-- {
		id := i
		for j := i - 1; j >= 0; j-- {
			if buf[j] > buf[i] && buf[j] >= buf[id] {
				id = j
			}
		}

		if buf[id] > buf[i] {
			buf[id], buf[i] = buf[i], buf[id]
			break
		}
	}

	n = 0
	for i := len(buf) - 1; i >= 0; i-- {
		n = 10*n + buf[i]
	}

	return n
}
