package sem_4

func ExtraLetter(first string, second string) byte {

	firstMap := make(map[byte]int, 0)

	for iter := range first {

		cnt := firstMap[first[iter]]
		firstMap[first[iter]] = cnt + 1

	}

	for iter := range second {

		cnt := firstMap[second[iter]]
		if cnt == 0 {
			return second[iter]
		}
		firstMap[second[iter]] = cnt - 1
	}

	return 0
}
