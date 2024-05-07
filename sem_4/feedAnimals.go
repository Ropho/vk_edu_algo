package sem_4

func FeedAnimals(animals []int, food []int) int {

	// sort.Ints(animals)
	// sort.Ints(food)
	animals = ShellSort(animals)
	food = ShellSort(food)

	counter := 0
	animalIndex := 0

	for _, foodVal := range food {

		if animalIndex >= len(animals) {
			break
		}

		if foodVal >= animals[animalIndex] {
			counter++
			animalIndex++
		}
	}

	return counter
}
