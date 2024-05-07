package sem_4

func GroupAnagrams(words []string) [][]string {

	m := make(map[int][]string, 0)

	for _, word := range words {
		hash := anagramHash(word)

		sl, ok := m[hash]
		if !ok {
			m[hash] = []string{word}
		} else {
			m[hash] = append(sl, word)
		}
	}

	ans := make([][]string, 0, len(m))

	for _, slice := range m {
		ans = append(ans, slice)
	}
	return ans
}

func anagramHash(s string) int {
	sum := 0
	for index := range s {
		sum += int(s[index])
	}

	return sum
}
