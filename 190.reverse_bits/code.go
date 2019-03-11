package main

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		u := uint32(i)
		b := (num >> u) % 2
		res ^= b << (31 - u)
	}
	return res
}
