// Package errors implements a simple error handling package.
//
// The Errors struct represents an error that can wrap multiple parent errors.
// It provides methods for creating and manipulating errors.
//
// Example usage:
//
//	err := errors.New("Something went wrong")
//	if err != nil {
//	    // Handle the error
//	}
//
//	wrappedErr := errors.Wrap("Another error occurred", err)
//	if wrappedErr != nil {
//	    // Handle the wrapped error
//	}
//
//	if errors.Is(err, targetErr) {
//	    // Check if the error matches the target error
//	}
//
//	var customErr CustomError
//	if errors.As(err, &customErr) {
//	    // Check if the error can be converted to CustomError
//	}
//
//	joinedErr := errors.Join(err1, err2)
//	if joinedErr != nil {
//	    // Handle the joined error
//	}
package errors

import (
	"errors"
)

// Errors represents an error that can wrap multiple parent errors.
type Errors struct {
	parents []error
	err     error
}

// Error returns the string representation of the error.
// If the error itself is nil, it returns the first parent error's string representation.
func (e *Errors) Error() string {
	if e.err == nil {
		if len(e.parents) == 0 {
			return ""
		}
		return e.parents[0].Error()
	}
	return e.err.Error()
}

// New creates a new Errors instance with the given message.
// It returns a pointer to the Errors instance.
func New(message string) *Errors {
	return &Errors{err: errors.New(message)}
}

// Wrap creates a new Errors instance with the given message and parent errors.
// It returns a pointer to the Errors instance.
func Wrap(message string, parents ...error) *Errors {
	return &Errors{parents: parents, err: errors.New(message)}
}

// Unwrap returns a slice of all the wrapped errors, including the error itself.
func (e *Errors) Unwrap() []error {
	list := make([]error, len(e.parents)+1)
	list[0] = e.err
	copy(list[1:], e.parents)
	if e.err == nil {
		list = list[1:]
	}
	return list
}

// Is reports whether any error in the chain matches the target error.
// It recursively checks for matching errors in the wrapped error chain.
func Is(err, target error) bool {
	if err == target {
		return true
	}
	switch err := err.(type) {
	case interface{ Unwrap() error }:
		return Is(err.Unwrap(), target)
	case interface{ Unwrap() []error }:
		for _, e := range err.Unwrap() {
			if Is(e, target) {
				return true
			}
		}
	}
	return false
}

// As checks if the error or any of its wrapped errors can be converted to the target type.
// If successful, it assigns the converted error to the target parameter and returns true.
// It recursively checks for convertible errors in the wrapped error chain.
func As[T any](e error, target *T) bool {
	if e != nil {
		x, ok := e.(T)
		if ok {
			*target = x
			return true
		}
	}
	switch e := e.(type) {
	case interface{ Unwrap() error }:
		return As(e.Unwrap(), target)
	case interface{ Unwrap() []error }:
		for _, e := range e.Unwrap() {
			if As(e, target) {
				return true
			}
		}
	}
	return false
}

// Join combines multiple errors into a single Errors instance.
// It discards nil errors and returns nil if no errors are provided.
// If there is only one non-nil error, it returns an Errors instance wrapping that error.
// If there are more than one non-nil errors, it returns an Errors instance with the first error as the main error
// and the remaining errors as parent errors.
func Join(err ...error) *Errors {
	switch len(err) {
	case 0:
		return nil
	case 1:
		if err[0] == nil {
			return nil
		}
		return &Errors{err: err[0]}
	}
	parents := make([]error, 0, len(err))
	for _, e := range err {
		if e != nil {
			parents = append(parents, e)
		}
	}
	if len(parents) >= 2 {
		return &Errors{
			err:     parents[0],
			parents: parents[1:],
		}
	}
	return Join(parents...)
}
