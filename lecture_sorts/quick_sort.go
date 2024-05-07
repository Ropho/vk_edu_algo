package lecture_sorts

// func QuickSort(sl []int) []int {
// 	pivot := len(sl) / 2
// 	left := 0
// 	right := len(sl) - 1

// 	for left < right {
// 		for left := 0; left < len(sl); left++ {
// 			if sl[left] > sl[pivot] {
// 				break
// 			}
// 		}

// 		for right := len(sl) - 1; right >= 0; right-- {
// 			if sl[right] <= sl[pivot] {
// 				break
// 			}
// 		}

// 		if left < right {
// 			tmp := sl[left]
// 			sl[left] = sl[right]
// 			sl[right] = tmp
// 		}
// 	}

// 	tmp := sl[left]
// 	sl[left] = sl[pivot]
// 	sl[pivot] = tmp

// 	if pivot > 0
// 	QuickSort(sl[:pivot])
// 	QuickSort(sl[pivot+1:])

// 	return sl
// }
