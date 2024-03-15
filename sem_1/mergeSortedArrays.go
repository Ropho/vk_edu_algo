package sem_1

// MergeSortedArrays merges two given arrays
func MergeSortedArrays(arr1, arr2 []int) []int {

	merged := make([]int, 0, len(arr1)+len(arr2))

	var firstIndex, secondIndex int
	var min int

	for firstIndex < len(arr1) && secondIndex < len(arr2) {

		if arr1[firstIndex] < arr2[secondIndex] {
			min = arr1[firstIndex]
			firstIndex++
		} else {
			min = arr2[secondIndex]
			secondIndex++
		}

		merged = append(merged, min)

	}

	merged = append(merged, arr1[firstIndex:]...)
	merged = append(merged, arr2[secondIndex:]...)

	return merged
}
