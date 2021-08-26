package main

// medium
// https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/

func numOfSubarrays1524(arr []int) int {
	n := len(arr)

	// dp[x][i] means number of sub-array end at i that has even/odd sum
	dp := make([][]int, 2)
	dp[0] = make([]int, n) // even
	dp[1] = make([]int, n) // odd

	r := 0
	if arr[0]%2 == 0 {
		dp[0][0] = 1
	} else {
		dp[1][0] = 1
		r++
	}

	for i := 1; i < n; i++ {
		v := arr[i]
		if v%2 == 0 {
			dp[0][i] = dp[0][i-1] + 1
			dp[1][i] = dp[1][i-1]
		} else {
			dp[1][i] = dp[0][i-1] + 1
			dp[0][i] = dp[1][i-1]
		}
		r = (r + dp[1][i]) % (1e9 + 7)
	}

	return r
}
