package sem3

import (
	"github.com/go-playground/validator"
)

func AbsInt(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

func BinarySearchSqrt(target int) int {

	var err error
	validate := validator.New()

	err = validate.Var(target, "gte=0")
	if err != nil {
		return -1
	}

	l := 0
	r := target

	for l < r-1 {
		pivot := l + (r-l)/2

		if pivot*pivot < target {
			l = pivot
		} else {
			r = pivot
		}
	}

	if AbsInt(l*l-target) < AbsInt(r*r-target) {
		return l
	}

	return r
}
