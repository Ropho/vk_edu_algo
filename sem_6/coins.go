package sem_6

func Coins(coins []int, target int) int {

	for i := len(coins) - 1; i >= 0; i-- {
		if target-coins[i] > 0 {

			recur := Coins(coins[:i+1], target-coins[i])
			if recur != -1 {
				return 1 + recur
			}

			return -1

		} else if target-coins[i] == 0 {
			return 1
		}
	}

	return -1
}
