package main

import (
	"fmt"
	"github.com/snowmerak/errors/lib/errors"
	"github.com/snowmerak/errors/lib/formatter"
)

func main() {
	test()
}

func test() {
	pErr := errors.New(formatter.NewJsonFormatter).Int64("hello", 123).Msg("hello")
	fmt.Println(pErr)

	cErr := errors.New(formatter.NewTupleFormatter).Caller().Err("hello", pErr).Msg("hello")
	fmt.Println(cErr)

	fmt.Println(errors.Is(cErr, pErr))

	pErr.Free()
	cErr.Free()
}

/*
{"hello":123,"msg":"hello"}
caller="/Users/gwon-yongmin/Documents/GitHub/errors/cmd/prac/main.go:17" hello={"hello":123,"msg":"hello"} msg="hello"
true
*/
