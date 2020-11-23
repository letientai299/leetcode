package main

// medium
//
// this is exactly the problem with 33, but with better test cases.
func search_81(nums []int, t int) bool {
	return search_33(nums, t) != -1
}
