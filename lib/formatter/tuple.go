package formatter

import (
	"io"
	"runtime"
	"strconv"
	"strings"
)

type TupleFormatter struct {
	writer   io.Writer
	elements int
}

func (t *TupleFormatter) writeKey(key string) (int, error) {
	n := 0
	if t.elements > 0 {
		written, err := t.writer.Write([]byte{' '})
		if err != nil {
			return n, err
		}
		n += written
	}
	t.elements++
	written, err := t.writer.Write([]byte(key))
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte{'='})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) Byte(s string, b byte) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte{'\'', b, '\''})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) Bytes(s string, bytes []byte) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte{'['})
	if err != nil {
		return n, err
	}
	n += written
	for _, b := range bytes {
		written, err = t.writer.Write([]byte(strconv.FormatInt(int64(b), 10)))
		if err != nil {
			return n, err
		}
		n += written
	}
	written, err = t.writer.Write([]byte{']'})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) Float64(s string, f float64) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte(strconv.FormatFloat(f, 'f', -1, 64)))
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) Int64(s string, i int64) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte(strconv.FormatInt(i, 10)))
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) Uint64(s string, u uint64) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte(strconv.FormatUint(u, 10)))
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) String(s string, s2 string) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte(s))
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) Bool(s string, b bool) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte(strconv.FormatBool(b)))
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (t *TupleFormatter) Err(s string, parentErr error) (int, error) {
	n := 0
	written, err := t.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	content := ""
	switch parentErr.(type) {
	case nil:
	default:
		content = parentErr.Error()
	}
	if !(strings.HasPrefix(content, "{") && strings.HasSuffix(content, "}")) {
		written, err = t.writer.Write([]byte{'"'})
		if err != nil {
			return n, err
		}
		n += written
	}
	written, err = t.writer.Write([]byte(content))
	if err != nil {
		return n, err
	}
	n += written
	if !(strings.HasPrefix(content, "{") && strings.HasSuffix(content, "}")) {
		written, err = t.writer.Write([]byte{'"'})
		if err != nil {
			return n, err
		}
		n += written
	}
	return n, nil
}

func (t *TupleFormatter) Caller() (int, error) {
	n := 0
	written, err := t.writeKey("caller")
	if err != nil {
		return n, err
	}
	n += written
	_, file, line, ok := runtime.Caller(2)
	switch ok {
	case true:
		written, err = t.writer.Write([]byte{'"'})
		if err != nil {
			return n, err
		}
		n += written
		written, err = t.writer.Write([]byte(file))
		if err != nil {
			return n, err
		}
		n += written
		written, err = t.writer.Write([]byte{':'})
		if err != nil {
			return n, err
		}
		n += written
		written, err = t.writer.Write([]byte(strconv.Itoa(line)))
		if err != nil {
			return n, err
		}
		n += written
		written, err = t.writer.Write([]byte{'"'})
		if err != nil {
			return n, err
		}
		n += written
	case false:
		written, err = t.writer.Write([]byte{'"', '"'})
		if err != nil {
			return n, err
		}
		n += written
	}
	return n, nil
}

func (t *TupleFormatter) Msg(s string) (int, error) {
	n := 0
	written, err := t.writeKey("msg")
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte(s))
	if err != nil {
		return n, err
	}
	n += written
	written, err = t.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func NewTupleFormatter(writer io.Writer) (Formatter, error) {
	return &TupleFormatter{
		writer: writer,
	}, nil
}
