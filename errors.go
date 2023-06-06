package errors

import (
	"errors"
)

type Errors struct {
	parents []error
	err     error
}

func (e *Errors) Error() string {
	if e.err == nil {
		if len(e.parents) == 0 {
			return ""
		}
		return e.parents[0].Error()
	}
	return e.err.Error()
}

func New(message string) *Errors {
	return &Errors{err: errors.New(message)}
}

func Wrap(message string, parents ...error) *Errors {
	return &Errors{parents: parents, err: errors.New(message)}
}

func (e *Errors) Unwrap() []error {
	list := make([]error, len(e.parents)+1)
	list[0] = e.err
	copy(list[1:], e.parents)
	if e.err == nil {
		list = list[1:]
	}
	return list
}

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

func Join(err ...error) *Errors {
	switch len(err) {
	case 0:
		return nil
	case 1:
		if err[0] == nil {
			return nil
		}
		return &Errors{err: err[0]}
	default:
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
}
