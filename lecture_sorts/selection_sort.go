package lecture_sorts

func SelectionSort(sl []int) []int {

	for curPos := 0; curPos < len(sl)-1; curPos++ {

		minIndex := curPos

		for i := curPos + 1; i < len(sl); i++ {
			if sl[i] < sl[minIndex] {
				minIndex = i
			}
		}

		tmp := sl[curPos]
		sl[curPos] = sl[minIndex]
		sl[minIndex] = tmp
	}

	return sl
}
