package errs

import "fmt"

func ErrorIndexOutOfRange(len int, index int) error {
	return fmt.Errorf("Out of range, length %d, index %d", len, index)
}

func ErrorInvalidType(want string, got any) error {
	return fmt.Errorf("Type conversion failed, expected type:%s, actual value:%#v", want, got)
}
