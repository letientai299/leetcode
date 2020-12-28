package main

// medium

// using inequality equations, we can prove that
// the best value to make all element equal to is median.
// simple solution, but O(n log n).
//func minMoves2(nums []int) int {
//	s := 0
//	sort.Ints(nums)
//	n := len(nums)
//	for i:=0; i< n/2; i++ {
//		s+=nums[n-1-i]-nums[i]
//	}
//	return s
//}

func minMoves2(nums []int) int {
	n := len(nums)
	var m int
	if n%2 == 1 {
		m = quickSelect(nums, (n+1)/2)
	} else {
		m = (quickSelect(nums, n/2) + quickSelect(nums, n/2+1)) / 2
	}
	s := 0
	for _, x := range nums {
		if x > m {
			s += x - m
		} else {
			s += m - x
		}
	}
	return s
}
