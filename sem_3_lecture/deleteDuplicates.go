package sem3

func DeleteDuplicates(s string) string {

	stk := Stack[byte]{}

	for index := range s {

		if stk.Peek() == s[index] {
			stk.Pop()
		} else {
			stk.Push(s[index])
		}
	}

	ans := make([]byte, stk.Size())
	index := 0
	for stk.Size() != 0 {
		ans[len(ans)-index-1] = stk.Peek()
		stk.Pop()
		index++
	}

	return string(ans)
}
