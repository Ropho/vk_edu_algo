package sem_6

func findLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
		}
	}

	return sliceMax(dp)
}

func sliceMax(sl []int) int {
	max := sl[0]

	for _, el := range sl {
		if el > max {
			max = el
		}
	}

	return max
}
