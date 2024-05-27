package sem_6

func ZeroOnes(n int) int {
	if n == 1 {
		return 2
	}

	dp := make([]int, n)

	dp[0] = 2
	dp[1] = 3

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}
