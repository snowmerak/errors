package main

import (
	"fmt"
	"github.com/snowmerak/errors/lib/bufferedlist"
	"github.com/snowmerak/errors/lib/formatter"
	"io"
)

func main() {
	test()
}

func test() {
	bl := bufferedlist.New()

	jf, err := formatter.NewJsonFormatter(bl)
	if err != nil {
		panic(err)
	}

	jf.Byte("byte", 'a')
	jf.Int64("int64", 64)
	jf.Uint64("uint64", 64)
	jf.Float64("float64", 64.0)
	jf.String("string", "string")
	jf.Bool("bool", true)
	jf.Err("err", nil)
	jf.Caller()
	jf.Msg("msg")

	reader := bl.Reader()
	buf := [1024]byte{}
	for {
		n, err := reader.Read(buf[:])
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
