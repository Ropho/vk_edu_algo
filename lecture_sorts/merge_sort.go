package lecture_sorts

import "vk_algo/sem_1"

func MergeSort(sl []int) []int {

	if len(sl) < 2 {
		return sl
	}

	if len(sl) == 2 {
		if sl[0] > sl[1] {
			tmp := sl[1]
			sl[1] = sl[0]
			sl[0] = tmp
		}
		return sl
	}

	left := MergeSort(sl[0 : len(sl)/2])
	right := MergeSort(sl[len(sl)/2:])

	merged := sem_1.MergeSortedArrays(left, right)

	return merged
}
