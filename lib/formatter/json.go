package formatter

import (
	"encoding/base64"
	"io"
	"runtime"
	"strconv"
)

type JsonFormatter struct {
	writer   io.Writer
	elements int
}

func (j *JsonFormatter) writeKey(key string) (int, error) {
	n := 0
	if j.elements > 0 {
		written, err := j.writer.Write([]byte{','})
		if err != nil {
			return 0, err
		}
		n += written
	}
	written, err := j.writer.Write([]byte{'"'})
	if err != nil {
		return 0, err
	}
	n += written
	written, err = j.writer.Write([]byte(key))
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"', ':'})
	if err != nil {
		return n, err
	}
	n += written
	j.elements++
	return n, nil
}

func (j *JsonFormatter) Byte(s string, b byte) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'\''})
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{b})
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'\''})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (j *JsonFormatter) Bytes(s string, bytes []byte) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(bytes)))
	base64.URLEncoding.Encode(encoded, bytes)
	written, err = j.writer.Write(encoded)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (j *JsonFormatter) Float64(s string, f float64) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte(strconv.FormatFloat(f, 'f', -1, 64)))
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (j *JsonFormatter) Int64(s string, i int64) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte(strconv.FormatInt(i, 10)))
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (j *JsonFormatter) Uint64(s string, u uint64) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte(strconv.FormatUint(u, 10)))
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (j *JsonFormatter) String(s string, s2 string) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte(s2))
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (j *JsonFormatter) Bool(s string, b bool) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	switch b {
	case true:
		written, err = j.writer.Write([]byte{'t', 'r', 'u', 'e'})
		if err != nil {
			return n, err
		}
		n += written
	case false:
		written, err = j.writer.Write([]byte{'f', 'a', 'l', 's', 'e'})
		if err != nil {
			return n, err
		}
		n += written
	}
	return n, nil
}

func (j *JsonFormatter) Err(s string, err error) (int, error) {
	n := 0
	written, err := j.writeKey(s)
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	switch err.(type) {
	case nil:
	default:
		written, err = j.writer.Write([]byte(err.Error()))
		if err != nil {
			return n, err
		}
		n += written
	}
	written, err = j.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func (j *JsonFormatter) Caller() (int, error) {
	n := 0
	written, err := j.writeKey("caller")
	if err != nil {
		return n, err
	}
	n += written
	_, file, line, ok := runtime.Caller(1)
	switch ok {
	case true:
		written, err = j.writer.Write([]byte{'"'})
		if err != nil {
			return n, err
		}
		n += written
		written, err = j.writer.Write([]byte(file))
		if err != nil {
			return n, err
		}
		n += written
		written, err = j.writer.Write([]byte{':'})
		if err != nil {
			return n, err
		}
		n += written
		written, err = j.writer.Write([]byte(strconv.Itoa(line)))
		if err != nil {
			return n, err
		}
		n += written
		written, err = j.writer.Write([]byte{'"'})
		if err != nil {
			return n, err
		}
		n += written
	case false:
		written, err = j.writer.Write([]byte{'"', '"'})
		if err != nil {
			return n, err
		}
		n += written
	}
	return n, nil
}

func (j *JsonFormatter) Msg(s string) (int, error) {
	n := 0
	written, err := j.writeKey("msg")
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"'})
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte(s))
	if err != nil {
		return n, err
	}
	n += written
	written, err = j.writer.Write([]byte{'"', '}'})
	if err != nil {
		return n, err
	}
	n += written
	return n, nil
}

func NewJsonFormatter(writer io.Writer) (*JsonFormatter, error) {
	if n, err := writer.Write([]byte("{")); err != nil {
		return nil, err
	} else if n != 1 {
		return nil, io.ErrShortBuffer
	}
	return &JsonFormatter{
		writer: writer,
	}, nil
}
