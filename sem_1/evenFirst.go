package sem_1

// EvenFirst sorts array that at first go even numbers then odd numbers
func EvenFirst(arr []int) []int {

	evenIndex := 0

	for i := 0; i < len(arr); i++ {

		if arr[i]%2 == 0 {

			tmp := arr[i]
			arr[i] = arr[evenIndex]
			arr[evenIndex] = tmp
			evenIndex++
		}
	}

	return arr
}
