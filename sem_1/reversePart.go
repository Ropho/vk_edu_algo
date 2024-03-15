package sem_1

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ReversePart reverts only part of array
// based on input parameter starting from "end" index
func ReversePart(arr []int, endIndex int) ([]int, error) {
	var err error

	validate := validator.New()

	err = validate.Var(len(arr), "gt=0")
	if err != nil {
		return nil, fmt.Errorf("validation failed for array")
	}
	err = validate.Var(endIndex, "gte=0")
	if err != nil {
		return nil, fmt.Errorf("validation failed for endIndex")
	}

	sl, err := ReverseArr(arr, 0, len(arr)-1)
	if err != nil {
		return nil, fmt.Errorf("failed to reverse part: [%w]", err)
	}

	sl, err = ReverseArr(sl, 0, endIndex%len(arr)-1)
	if err != nil {
		return nil, fmt.Errorf("failed to reverse part: [%w]", err)
	}

	sl, err = ReverseArr(sl, endIndex%len(arr), len(arr)-1)
	if err != nil {
		return nil, fmt.Errorf("failed to reverse part: [%w]", err)
	}

	return sl, nil
}
