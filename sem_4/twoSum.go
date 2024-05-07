package sem_4

func twoSum(data []int, target int) (int, int) {

	cache := make(map[int]struct{}, 0)

	for _, num := range data {

		_, ok := cache[target-num]
		if ok {
			return target - num, num
		}

		cache[num] = struct{}{}
	}

	return 0, 0
}
