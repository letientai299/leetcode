package main

// medium

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{""}
	build := func(s string) {
		next := make([]string, 0, len(res)*len(s))
		for _, c := range s {
			for _, x := range res {
				next = append(next, x+string(c))
			}
		}
		res = next
	}

	for _, d := range digits {
		build(m[byte(d)])
	}

	return res
}
