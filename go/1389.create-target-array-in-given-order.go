package main

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, 0, len(nums))
	for i, n := range nums {
		id := index[i]
		for x := id; x < len(res); x++ {
			res[x], n = n, res[x]
		}
		res = append(res, n)
	}

	return res
}
