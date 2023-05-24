package errors

import (
	"github.com/snowmerak/errors/lib/bufferedlist"
	"github.com/snowmerak/errors/lib/formatter"
	"io"
	"strings"
)

type Errors struct {
	formatter formatter.Formatter
	buffer    *bufferedlist.BufferedList
}

func New(formatter func(io.Writer) (formatter.Formatter, error)) *Errors {
	bl := bufferedlist.New()
	f, _ := formatter(bl)
	err := &Errors{
		formatter: f,
		buffer:    bl,
	}
	return err
}

func (e *Errors) Byte(s string, b byte) *Errors {
	_, _ = e.formatter.Byte(s, b)
	return e
}

func (e *Errors) Bytes(s string, b []byte) *Errors {
	_, _ = e.formatter.Bytes(s, b)
	return e
}

func (e *Errors) Float64(s string, f float64) *Errors {
	_, _ = e.formatter.Float64(s, f)
	return e
}

func (e *Errors) Int64(s string, i int64) *Errors {
	_, _ = e.formatter.Int64(s, i)
	return e
}

func (e *Errors) Uint64(s string, u uint64) *Errors {
	_, _ = e.formatter.Uint64(s, u)
	return e
}

func (e *Errors) String(s string, s2 string) *Errors {
	_, _ = e.formatter.String(s, s2)
	return e
}

func (e *Errors) Bool(s string, b bool) *Errors {
	_, _ = e.formatter.Bool(s, b)
	return e
}

func (e *Errors) Err(s string, err error) *Errors {
	_, _ = e.formatter.Err(s, err)
	return e
}

func (e *Errors) Caller() *Errors {
	_, _ = e.formatter.Caller()
	return e
}

func (e *Errors) Msg(s string) *Errors {
	_, _ = e.formatter.Msg(s)
	return e
}

func (e *Errors) Error() string {
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
	e.buffer.Free()
	return builder.String()
}
