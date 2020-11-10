package main

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := make([]int, 0, len(nums))
	res = append(res, nums[0])
	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]+nums[i])
	}
	return res
}
