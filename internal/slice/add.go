package slice

import "toys4go/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.ErrorIndexOutOfRange(length, index)
	}

	// add a new element to src first
	var zeroValue T
	src = append(src, zeroValue)

	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}
