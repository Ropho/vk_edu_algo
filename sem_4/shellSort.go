package sem_4

import "golang.org/x/exp/constraints"

func ShellSort[T constraints.Ordered](sl []T) []T {

	for gap := len(sl) / 2; gap > 0; gap /= 2 {

		curPos := gap

		for ; curPos < len(sl); curPos++ {

			for len := 1; curPos-gap*len >= 0; len++ {

				if sl[curPos-(len-1)*gap] >= sl[curPos-len*gap] {
					break
				}

				tmp := sl[curPos-(len-1)*gap]
				sl[curPos-(len-1)*gap] = sl[curPos-len*gap]
				sl[curPos-len*gap] = tmp
			}
		}
	}

	return sl
}
