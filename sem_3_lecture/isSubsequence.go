package sem3

import (
	"github.com/go-playground/validator"
)

func isSubsequenceQueue(seq, subseq string) bool {

	var err error
	validate := validator.New()

	err = validate.Var(len(seq)-len(subseq), "gte=0")
	if err != nil {
		return false
	}

	q := Queue[rune]{}

	for _, char := range subseq {
		q.Push(char)
	}

	for _, char := range seq {
		if char == q.PeekEnd() {
			q.Pop()
		}
	}

	return q.Size() == 0
}

func isSubsequencePointers(seq, subseq string) bool {

	seqRune := []rune(seq)
	subseqRUne := []rune(subseq)

	seqIndex := 0
	subseqIndex := 0

	for seqIndex < len(seqRune) && subseqIndex < len(subseqRUne) {
		if seqRune[seqIndex] == subseqRUne[subseqIndex] {
			subseqIndex++
		}
		seqIndex++
	}

	return subseqIndex == len(subseqRUne)
}
