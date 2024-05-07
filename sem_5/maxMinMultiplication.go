package sem_5

import "fmt"

func MaxMinMultiplication(data []int) int {

	if len(data) < 3 {
		return -1
	}

	minIndex := 0
	maxIndex := 0

	for i := 0; i < len(data) && data[i] != 0; i = 2*i + 1 {
		fmt.Println(i)
		minIndex = i
	}

	for i := 0; i < len(data) && data[i] != 0; i = 2*i + 2 {
		fmt.Println(i)

		maxIndex = i
	}

	return data[minIndex] * data[maxIndex]
}
