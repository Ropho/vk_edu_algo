package sem_1

// ZerosEnd moves zeros to the end of slice
// Saves order of other numbers!!
func ZerosEnd(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] == 0 {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}
