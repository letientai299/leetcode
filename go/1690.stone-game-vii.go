package main

// medium

// TODO (tai): not done yet
func stoneGameVII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	alice := 0
	bob := 0
	isAlice := true
	for i, j := 0, len(stones)-1; i < j; {
		a, b, c := stones[i]+stones[j], stones[i]+stones[i+1], stones[j]+stones[j-1]
		pickLeft := true
		if a >= b && a >= c {
			if b >= c {
				pickLeft = false
			}
		} else if a >= b { // and a < c
			pickLeft = true
		} else if a >= c { // and a < b
			pickLeft = false
		} else if b >= c { // b>=c>=a
			pickLeft = false
		} else { // c>=b>=a
			pickLeft = true
		}

		pickLeft = pickLeft == isAlice

		if pickLeft {
			sum -= stones[i]
			i++
		} else {
			sum -= stones[j]
			j--
		}

		if isAlice {
			alice += sum
		} else {
			bob += sum
		}

		isAlice = !isAlice
	}

	return alice - bob
}
