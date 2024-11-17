package slice

import "toys4go/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zeroValue T
		return nil, zeroValue, errs.ErrorIndexOutOfRange(length, index)
	}
	res := src[index]

	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]

	}

	src = src[:length-1]
	return src, res, nil
}
