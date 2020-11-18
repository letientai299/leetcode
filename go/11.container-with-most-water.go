package main

func maxArea(height []int) int {
	n := len(height)
	best := 0
	for i, j := 0, n-1; i < j; {
		length := j - i
		if height[i] > height[j] {
			if (height[j] * length) > best {
				best = height[j] * length
			}
			j--
		} else {
			if (height[i] * length) > best {
				best = height[i] * length
			}
			i++
		}
	}
	return best
}
