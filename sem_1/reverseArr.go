package sem_1

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ReverseArr reverts slice from start index to end index
func ReverseArr(arr []int, start, end int) ([]int, error) {
	var err error

	validate := validator.New()

	err = validate.Var(len(arr), "gt=0")
	if err != nil {
		return nil, fmt.Errorf("validation failed for array")
	}
	err = validate.Var(start, fmt.Sprintf("gte=0,lte=%d", len(arr)-1))
	if err != nil {
		return nil, errors.New("validation failed for start index")
	}
	err = validate.Var(end, fmt.Sprintf("gte=0,lte=%d", len(arr)-1))
	if err != nil {
		return nil, errors.New("validation failed for end index")
	}

	for start < end {

		tmp := arr[start]
		arr[start] = arr[end]
		arr[end] = tmp

		start++
		end--
	}

	return arr, nil
}
