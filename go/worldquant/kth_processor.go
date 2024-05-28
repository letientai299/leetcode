package worldquant

import (
	"fmt"
)

// KthProcessor
//
// Given m jobs distributed on n processors, in which the k-th is the most
// efficient. The jobs should be distributed in a balance way such that the
// different between 2 neighboring processors is not larger than 1. Assuming
// m>=n, what is maximum jobs that k-th processor should have if each processor
// have at least 1 job.
// TODO: not sure if this is optimal yet
func KthProcessor(m int, n int, k int) int {
	if n == 1 {
		return m
	}

	// ensure k is on the right side
	k = max(k, n-k+1)
	k--

	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	var l int
	var r int
	for i := n; i < m; i++ {
		if (k == n-1 && dp[k] == dp[k-1]) || (dp[k] == dp[k-1] && dp[k] == dp[k+1]) {
			dp[k]++
			continue
		}

		for r = k + 1; r < n && dp[r] == dp[r-1]; r++ {
		}
		if r < n {
			dp[r]++
			continue
		}

		for l = k - 1; l > 0 && dp[l] > dp[l-1]; l-- {
		}
		dp[l]++
	}

	fmt.Println(dp)
	return dp[k]
}
