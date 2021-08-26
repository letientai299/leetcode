package main

// Iterator for Combination
//
// Medium
//
// https://leetcode.com/problems/iterator-for-combination/
//
// Design the `CombinationIterator` class:
//
// - `CombinationIterator(string characters, int combinationLength)` Initializes
// the object with a string `characters` of **sorted distinct** lowercase
// English letters and a number `combinationLength` as arguments.
// - `next()` Returns the next combination of length `combinationLength` in
// **lexicographical order**.
// - `hasNext()` Returns `true` if and only if there exists a next combination.
//
// **Example 1:**
//
// ```
// Input
// ["CombinationIterator", "next", "hasNext", "next", "hasNext", "next",
// "hasNext"]
// [["abc", 2], [], [], [], [], [], []]
// Output
// [null, "ab", true, "ac", true, "bc", false]
//
// Explanation
// CombinationIterator itr = new CombinationIterator("abc", 2);
// itr.next();    // return "ab"
// itr.hasNext(); // return True
// itr.next();    // return "ac"
// itr.hasNext(); // return True
// itr.next();    // return "bc"
// itr.hasNext(); // return False
//
// ```
//
// **Constraints:**
//
// - `1 <= combinationLength <= characters.length <= 15`
// - All the characters of `characters` are **unique**.
// - At most `104` calls will be made to `next` and `hasNext`.
// - It's guaranteed that all calls of the function `next` are valid.

type CombinationIterator struct {
	s   string
	com []int
	cnt int
	all int
}

func Constructor1286(s string, n int) CombinationIterator {
	com := make([]int, n)
	for i := range com {
		com[i] = i
	}

	l := len(s)
	all := 1
	for i := 1; n > 0; i++ {
		all = all * l / i
		l--
		n--
	}
	return CombinationIterator{
		s:   s,
		com: com,
		cnt: 0,
		all: all,
	}
}

func (ci *CombinationIterator) HasNext() bool {
	return ci.cnt < ci.all
}

func (ci *CombinationIterator) Next() string {
	res := make([]byte, len(ci.com))
	for i := range res {
		res[i] = ci.s[ci.com[i]]
	}

	n := len(ci.s) - 1
	i := len(ci.com) - 1
	for ; i >= 0; i-- {
		if ci.com[i] == n {
			n--
		} else {
			break
		}
	}

	if i != -1 {
		ci.com[i]++
		for j := i + 1; j < len(ci.com); j++ {
			ci.com[j] = ci.com[i] + j - i
		}
	}

	ci.cnt++
	return string(res)
}
