package sem_6

func PascalTriangle(n int) []int {

	dp := make([][]int, n)

	for i := 0; i < n; i++ {

		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			dp[i][j] = 1
		}
	}

	for row := 2; row < n; row++ {

		for col := 1; col <= row-1; col++ {
			dp[row][col] = dp[row-1][col-1] + dp[row-1][col]
		}
	}

	return dp[n-1]
}
