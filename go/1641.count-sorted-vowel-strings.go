package main

// medium
// https://leetcode.com/problems/count-sorted-vowel-strings/

func countVowelStrings(n int) int {
	if n == 0 {
		return 0
	}

	// a : a
	//     e
	//     i
	//     o
	//     u
	//
	// e : e
	//   : i
	//   : u
	//   : o
	//
	// i : 3
	// o : 2
	// u : 1
	//
	// 1 1 1 1 1
	// 1 2 3 4 5
	// 1 4 6 9 10

	v := [5]int{1, 1, 1, 1, 1}
	for n > 1 {
		s := 1
		for i := 1; i < 5; i++ {
			t := s
			s += v[i]
			v[i] += t
		}
		n--
	}

	s := 0
	for i := 0; i < 5; i++ {
		s += v[i]
	}

	return s
}
