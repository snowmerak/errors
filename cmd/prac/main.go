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
	err := errors.New(formatter.NewJsonFormatter)
	err.Caller()
	err.Msg("hello")
	fmt.Println(err.Error())
}
