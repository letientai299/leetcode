package main

// medium
//
//func sortColors(nums []int) {
//	cnt := [3]int{}
//	for _, c := range nums {
//		cnt[c]++
//	}
//	c := 0
//	for i := range nums {
//		for cnt[c] == 0 {
//			c++
//		}
//		cnt[c]--
//		nums[i] = c
//	}
//}

func sortColors(nums []int) {
	i, j, n := 0, len(nums)-1, len(nums)
	for i < n && nums[i] == 0 {
		i++
	}
	for j >= 0 && nums[j] == 2 {
		j--
	}
	for k := i; k <= j; k++ {
		if nums[k] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			i++
			k = i
		} else if nums[k] == 2 {
			nums[j], nums[k] = nums[k], nums[j]
			j--
			k--
		}
	}
}
