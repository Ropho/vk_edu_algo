package sem_1

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// MergeArrInPlace merges two sorted arrays w/o additional memory
// arr1 has preallocated space for arr2 elements as well
func MergeArrInPlace(arr1, arr2 []int, lengthFirst int) ([]int, error) {
	var err error
	validate := validator.New()

	err = validate.Var(len(arr1), "gt=0")
	if err != nil {
		return nil, fmt.Errorf("validation failed for length of array 1")
	}
	err = validate.Var(len(arr2), "gt=0")
	if err != nil {
		return nil, fmt.Errorf("validation failed for length of array 2")
	}
	err = validate.Var(lengthFirst, fmt.Sprintf("gt=0,lte=%d", len(arr1)-1))
	if err != nil {
		return nil, fmt.Errorf("validation failed for arr1 length")
	}
	err = validate.Var(len(arr1)-len(arr2)-lengthFirst, "gte=0")
	if err != nil {
		return nil, fmt.Errorf("validation failed for arr1 capacity")
	}

	ptr1 := lengthFirst - 1
	ptr2 := len(arr2) - 1
	ptr3 := ptr1 + ptr2 + 1

	for ptr2 >= 0 {

		if ptr1 >= 0 && arr1[ptr1] > arr2[ptr2] {
			arr1[ptr3] = arr1[ptr1]
			ptr1 -= 1

		} else {
			arr1[ptr3] = arr2[ptr2]
			ptr2 -= 1
		}

		ptr3 -= 1
	}

	return arr1, nil
}
