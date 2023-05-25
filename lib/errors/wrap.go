package errors

import "errors"

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, &target)
}

func Join(err ...error) error {
	return errors.Join(err...)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}
