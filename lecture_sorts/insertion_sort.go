package lecture_sorts

func InsertionSort(sl []int) []int {

	for curIndex := 1; curIndex < len(sl); curIndex++ {

		findIndex := 0
		for findIndex = 0; findIndex < curIndex; findIndex++ {

			if sl[curIndex] < sl[findIndex] {
				break
			}
		}

		tmp := sl[curIndex]

		for moveIndex := curIndex; moveIndex > findIndex; moveIndex-- {
			sl[moveIndex] = sl[moveIndex-1]
		}

		sl[findIndex] = tmp
	}

	return sl
}
