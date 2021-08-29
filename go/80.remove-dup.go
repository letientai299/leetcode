package main

func removeDuplicates80(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	i, j, p := 1, 1, nums[0]
	if nums[i] == p {
		i++
		j++
	}

	for ; j < n; j++ {
		if nums[j] == p {
			continue
		}

		p = nums[j]
		nums[i] = nums[j]
		i++
		if j < n-1 && nums[j+1] == p {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
