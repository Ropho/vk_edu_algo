package sem_5

func BSTCount(n int) int {

	if n == 0 {
		return 1
	}

	sum := 0
	for rootVal := 1; rootVal <= n; rootVal++ {
		sum += BSTCount(rootVal-1) * BSTCount(n-rootVal)
	}

	return sum
}
