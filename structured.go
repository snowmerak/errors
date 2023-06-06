// Package errors provides functionality for structured error handling.
//
// The StructuredError struct represents an error with associated fields.
// It allows for creating errors with named fields and provides methods for error formatting.
//
// Example usage:
//
//	field := make(map[string]interface{})
//	field["userID"] = 123
//	field["username"] = "john"
//	err := errors.From("Invalid input", field)
//	if err != nil {
//	    // Handle the error
//	}
//
//	value, ok := errors.Get(err, "userID")
//	if ok {
//	    // Use the retrieved value
//	}
package errors

import (
	"fmt"
	"strconv"
	"strings"
)

// StructuredError represents an error with associated fields.
type StructuredError struct {
	fields  map[string]any
	message string
}

// Error returns the string representation of the error.
// It formats the error message with the associated fields.
func (e *StructuredError) Error() string {
	buffer := strings.Builder{}
	c := false
	for key, value := range e.fields {
		if c {
			buffer.WriteString("&")
		}
		buffer.WriteString(key)
		buffer.WriteString("=")
		switch v := value.(type) {
		case string:
			buffer.WriteString("\"")
			buffer.WriteString(v)
			buffer.WriteString("\"")
		case int:
			buffer.WriteString(strconv.FormatInt(int64(v), 10))
		case int8:
			buffer.WriteString(strconv.FormatInt(int64(v), 10))
		case int16:
			buffer.WriteString(strconv.FormatInt(int64(v), 10))
		case int32:
			buffer.WriteString(strconv.FormatInt(int64(v), 10))
		case int64:
			buffer.WriteString(strconv.FormatInt(v, 10))
		case uint:
			buffer.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint8:
			buffer.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint16:
			buffer.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint32:
			buffer.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint64:
			buffer.WriteString(strconv.FormatUint(v, 10))
		case float32:
			buffer.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 32))
		case float64:
			buffer.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			buffer.WriteString(strconv.FormatBool(v))
		default:
			_, _ = fmt.Fprint(&buffer, v)
		}
		c = true
	}
	buffer.WriteString(" ")
	buffer.WriteString(e.message)
	return buffer.String()
}

// From creates a new Errors instance with the given message and fields.
// It returns a pointer to the Errors instance.
func From(message string, field map[string]any) *Errors {
	return &Errors{err: &StructuredError{fields: field, message: message}}
}

// Get retrieves the value associated with the specified key from the error's fields.
// It returns the retrieved value and a boolean indicating if the key was found.
// If the error is nil or the key is not found, it returns the zero value of type T and false.
// It recursively checks for the key in the wrapped error chain.
func Get[T any](err error, key string) (T, bool) {
	if err == nil {
		return *new(T), false
	}
	se, ok := err.(*StructuredError)
	if ok {
		value, ok := se.fields[key]
		if ok {
			v, ok := value.(T)
			if ok {
				return v, true
			}
		}
	}
	switch ue := err.(type) {
	case interface{ Unwrap() error }:
		return Get[T](ue.Unwrap(), key)
	case interface{ Unwrap() []error }:
		for _, e := range ue.Unwrap() {
			if v, ok := Get[T](e, key); ok {
				return v, ok
			}
		}
	}
	return *new(T), false
}
