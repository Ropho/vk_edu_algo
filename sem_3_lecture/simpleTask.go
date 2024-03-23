package sem3

import (
	"github.com/go-playground/validator"
)

func SimpleTask(copyNum, first, second int) int {

	var err error
	validate := validator.New()

	err = validate.Var(first, "gt=0")
	if err != nil {
		return -1
	}
	err = validate.Var(second, "gt=0")
	if err != nil {
		return -1
	}
	err = validate.Var(copyNum, "gt=0")
	if err != nil {
		return -1
	}

	var min, max int

	if first <= second {
		min = first
		max = second
	} else {
		min = second
		max = first
	}

	copyTime := min

	l := 1
	r := copyNum - 1

	for l < r-1 {

		pivot := l + (r-l)/2

		if pivot/min+pivot/max < copyNum-1 {
			l = pivot
		} else {
			r = pivot
		}
	}

	copyTime += r

	return copyTime
}
