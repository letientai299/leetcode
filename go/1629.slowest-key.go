package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	releaseTimes = append([]int{0}, releaseTimes...)
	max := 0
	var k byte
	for i, x := range keysPressed {
		dur := releaseTimes[i+1] - releaseTimes[i]
		if max < dur {
			max = dur
			k = byte(x)
			continue
		}

		if max == dur && byte(x) > k {
			k = byte(x)
		}
	}

	return k
}
