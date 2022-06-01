package main

import "strings"

// Encode and Decode Strings
//
// Medium
//
// https://leetcode.com/problems/encode-and-decode-strings/
//

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	// input strings contains only ascii, so we use unicode for delim.
	return strings.Join(strs, "｜")
}

func (codec *Codec) Decode(strs string) []string {
	return strings.Split(strs, "｜")
}
