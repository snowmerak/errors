package main

import (
	"fmt"
	"github.com/snowmerak/errors/lib/bufferedlist"
	"io"
)

func main() {
	bl := bufferedlist.New()
	for i := 0; i < 8000; i++ {
		bl.WriteByte(byte(i % 256))
	}

	buf := make([]byte, 3000)

	reader := bl.Reader()
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(buf[:n])
	fmt.Println("----------")

	n, err = reader.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(buf[:n])
	fmt.Println("----------")

	n, err = reader.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(buf[:n])
	fmt.Println("----------")
}
