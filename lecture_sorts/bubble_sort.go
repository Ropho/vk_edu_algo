package lecture_sorts

func BubbleSort(sl []int) []int {

	for iter := 0; iter < len(sl)-1; iter++ {

		for curIndex := 0; curIndex < len(sl)-iter-1; curIndex++ {
			if sl[curIndex] > sl[curIndex+1] {
				tmp := sl[curIndex+1]
				sl[curIndex+1] = sl[curIndex]
				sl[curIndex] = tmp
			}
		}
	}

	return sl
}
