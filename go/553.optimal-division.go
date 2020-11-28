package main

import (
	"fmt"
	"strconv"
	"strings"
)

// medium

func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 1 {
		return fmt.Sprint(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d/(%d", nums[0], nums[1]))
	for i := 2; i < n; i++ {
		sb.WriteString("/")
		sb.WriteString(strconv.Itoa(nums[i]))
	}
	sb.WriteString(")")
	return sb.String()
}
