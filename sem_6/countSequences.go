package sem_6

func CountSequences(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 2
	dp[2] = 4

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}
