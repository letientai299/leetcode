package main

func buildarray1441(target []int, n int) []string {
	res := make([]string, 0, n-len(target))
	last := 0
	for _, x := range target {
		for i := last + 1; i < x; i++ {
			res = append(res, "Push")
			res = append(res, "Pop")
		}

		res = append(res, "Push")
		last = x
	}
	return res
}
