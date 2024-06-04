package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	const s = "cbbd"
	p := longestPalindrome(s)
	log.Println(s)
	log.Println(p)
}

func longestPalindrome(str string) string {
	// interleave s with a new character to ensure all palindromes has odd length
	// with a center (even length palindrome doesn't have character center)
	s := interleave(str, "|")
	n := len(s)
	C := 0                     // center
	R := 0                     // radius
	maxRadii := make([]int, n) // all radii for each index in string s.

	for C < n {
		// find the longest radius from the current center,
		// note that we don't reset original radius to 0,
		// its starting value is set in the next logic block, from previous loop
		for C+R+1 < n && C-R-1 >= 0 && s[C+R+1] == s[C-R-1] {
			R++
		}
		maxRadii[C] = R

		// we know the longest palindrome starting at the current center.
		// we need to prepare the radius to check for the next centers.
		oldCenter := C
		oldRadius := R
		R = 0
		for C = oldCenter + 1; C < oldCenter+oldRadius; C++ {
			// C is now a location in the right half side of the old palindrome
			// mirrorC is its corresponding reflection on the left side.
			mirrorC := oldCenter - (C - oldCenter)
			// maxR is the max radius possible for C given our knowledge
			// about its location within the current palindrome.
			maxR := oldRadius - (C - oldCenter)
			mirrorR := maxRadii[mirrorC]

			// if the radius on the left side is exactly maxR, we know that on the
			// right side, we can potentially have a longer palindrome, so the next
			// radius should start at maxR, we can stop checking those right side
			// centers now, continue to the outer loop.
			if mirrorR == maxR {
				R = maxR
				break
			}

			// if the left side palindrome is completely contained within the current
			// palindrome, we can skip continue to the next center on the right side.
			if mirrorR < maxR {
				maxRadii[C] = mirrorR
				continue
			}

			if mirrorR > maxR {
				maxRadii[C] = maxR
				continue
			}
		}
	}

	bestR := 0
	bestI := 0
	for i, v := range maxRadii {
		if bestR < v {
			bestR = v
			bestI = i
		}
	}

	start := bestI/2 - bestR/2
	end := start + bestR
	return str[start:end]
}

func interleave(s string, c string) string {
	var sb strings.Builder
	sb.WriteString(c)
	for _, v := range s {
		sb.WriteRune(v)
		sb.WriteString(c)
	}
	return sb.String()
}

func biggest(ns ...int) int {
	b := ns[0]
	for _, n := range ns {
		b = max(n, b)
	}
	return b
}

func wait() {
	sc := bufio.NewReader(os.Stdin)
	_, _ = sc.ReadByte()
}
