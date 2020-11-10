package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, 0, len(candies))
	m := 0
	for _, v := range candies {
		if m < v {
			m = v
		}
	}
	for _, v := range candies {
		res = append(res, v+extraCandies >= m)
	}
	return res
}
