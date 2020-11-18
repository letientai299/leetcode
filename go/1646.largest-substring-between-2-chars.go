package main

func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[byte][]int)
	for i, x := range s {
		c := byte(x)
		m[c] = append(m[c], i)
	}

	max := 0
	for _, arr := range m {
		if len(arr) < 2 {
			continue
		}
		n := arr[len(arr)-1] - arr[0] - 1
		if max < n {
			max = n
		}
	}
	return max
}
