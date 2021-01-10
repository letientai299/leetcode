package main

// medium
// 1663. Smallest String With A Given Numeric Value
func getSmallestString(n int, k int) string {
	k -= n
	bs := make([]byte, n)

	// go forward faster than backward
	for i := range bs {
		bs[i] = 'a'
	}

	d := int('z' - 'a')
	i := n - 1
	for k > 0 {
		c := d
		if k < c {
			c = k
		}
		bs[i] += byte(c)
		i--
		k -= c
	}

	return string(bs)
}
