package main

// medium
// https://leetcode.com/problems/jump-game/

// slow, 25%
//func canJump(nums []int) bool {
//	n := len(nums)
//	for j := n - 1; j >= 0; j-- {
//		if nums[j] >= n-1-j {
//			nums[j] = -1
//			continue
//		}
//		for i := 1; i <= nums[j]; i++ {
//			if nums[j+i] == -1 {
//				nums[j] = -1
//				break
//			}
//		}
//	}
//	return nums[0] == -1
//}

func canJump(nums []int) bool {
	max := 0
	for i := 0; i <= max; i++ {
		if i+nums[i] >= len(nums)-1 {
			return true
		}
		if i+nums[i] >= max {
			max = i + nums[i]
		}
	}
	return max >= len(nums)-1
}
