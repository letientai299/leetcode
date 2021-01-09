package main

// medium
// 1386. Cinema Seat Allocation

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	if n == 0 {
		return 0
	}

	seats := make(map[int]int)
	for _, v := range reservedSeats {
		seats[v[0]-1] |= 1 << (v[1] - 1)
	}
	r := 0
	for _, v := range seats {
		v = v >> 1 &^ (1 << 8)
		if v == 0 {
			r += 2
		} else if (v>>4 == 0) != (v&15 == 0) {
			r += 1
		} else if (v>>2)&^(3<<4) == 0 {
			r += 1
		}
	}
	r += 2 * (n - len(seats))
	return r
}
