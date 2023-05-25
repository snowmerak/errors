package main

import (
	"fmt"
	"github.com/snowmerak/errors/lib/errors"
	"github.com/snowmerak/errors/lib/formatter"
	"os"
)

func main() {
	test()
}

func test() {
	pErr := errors.New(formatter.NewJsonFormatter).Int64("hello", 123).Msg("hello")

	errors.New(formatter.NewJsonFormatter).Caller().Err("hello", pErr).Msg("hello").MoveTo(os.Stdout)
	fmt.Println()
}
