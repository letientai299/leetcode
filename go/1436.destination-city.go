package main

func destCity(paths [][]string) string {
	src := make(map[string]struct{})
	dst := make(map[string]struct{})

	for _, ss := range paths {
		s, d := ss[0], ss[1]
		src[s] = struct{}{}
		dst[d] = struct{}{}
	}

	for d := range dst {
		if _, ok := src[d]; ok {
			return d
		}
	}

	return ""
}
