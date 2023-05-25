package errors

import (
	"github.com/snowmerak/errors/lib/bufferedlist"
	"github.com/snowmerak/errors/lib/formatter"
	"io"
	"strings"
)

type Error struct {
	formatter   formatter.Formatter
	buffer      *bufferedlist.BufferedList
	parentError error
}

func New(formatter func(io.Writer) (formatter.Formatter, error)) *Error {
	bl := bufferedlist.New()
	f, _ := formatter(bl)
	err := &Error{
		formatter: f,
		buffer:    bl,
	}
	return err
}

func (e *Error) Byte(s string, b byte) *Error {
	_, _ = e.formatter.Byte(s, b)
	return e
}

func (e *Error) Bytes(s string, b []byte) *Error {
	_, _ = e.formatter.Bytes(s, b)
	return e
}

func (e *Error) Float64(s string, f float64) *Error {
	_, _ = e.formatter.Float64(s, f)
	return e
}

func (e *Error) Int64(s string, i int64) *Error {
	_, _ = e.formatter.Int64(s, i)
	return e
}

func (e *Error) Uint64(s string, u uint64) *Error {
	_, _ = e.formatter.Uint64(s, u)
	return e
}

func (e *Error) String(s string, s2 string) *Error {
	_, _ = e.formatter.String(s, s2)
	return e
}

func (e *Error) Bool(s string, b bool) *Error {
	_, _ = e.formatter.Bool(s, b)
	return e
}

func (e *Error) Err(s string, err error) *Error {
	_, _ = e.formatter.Err(s, err)
	e.parentError = err
	return e
}

func (e *Error) Caller() *Error {
	_, _ = e.formatter.Caller()
	return e
}

func (e *Error) Msg(s string) *Error {
	_, _ = e.formatter.Msg(s)
	return e
}

func (e *Error) Unwrap() error {
	return e.parentError
}

func (e *Error) Error() string {
	builder := &strings.Builder{}
	reader := e.buffer.Reader()
	buf := [1024]byte{}
	for {
		n, err := reader.Read(buf[:])
		if err != nil && err != io.EOF {
			break
		}
		if n == 0 {
			break
		}
		_, _ = builder.Write(buf[:n])
	}
	return builder.String()
}

func (e *Error) MoveTo(writer io.Writer) error {
	reader := e.buffer.Reader()
	if _, err := io.Copy(writer, reader); err != nil {
		return err
	}
	return nil
}

func (e *Error) Free() {
	e.buffer.Free()
	e.buffer = nil
	e.parentError = nil
}
