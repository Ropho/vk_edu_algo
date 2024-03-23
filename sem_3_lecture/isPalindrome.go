package sem3

func IsPalindromeStack(s string) bool {

	stk := Stack[rune]{}

	for _, char := range s {
		stk.Push(char)
	}

	for _, char := range s {
		if stk.Peek() != char {
			return false
		}
		stk.Pop()
	}

	return true
}

func IsPalindromePointers(s string) bool {

	runeStr := []rune(s)
	start := 0
	end := len(runeStr) - 1

	for start < end {
		if runeStr[start] != runeStr[end] {
			return false
		}
		start++
		end--
	}

	return true
}
