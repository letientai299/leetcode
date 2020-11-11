package main

func decompressRLElist(nums []int) []int {
	res := make([]int, 0, len(nums)*2)
	for i := 0; i < len(nums); {
		for x := 0; x < nums[i]; x++ {
			res = append(res, nums[i+1])
		}
		i += 2
	}

	return res
}
