package main

func restoreString(s string, indices []int) string {
	b := []byte(s)
	i := 0
	n := len(indices)
	rem := n
	c := b[i]
	for rem > 0 {
		v := indices[i]
		if v != i {
			b[v], c = c, b[v]
		}
		indices[i] = -v - 1
		rem--
		i = v
		for cnt := 0; indices[i] < 0 && cnt < n; cnt++ {
			i = (i + n + 1) % n
			c = b[i]
		}
	}

	return string(b)
}
