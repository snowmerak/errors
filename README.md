# errors

errors is a simple error handling library for Go.

## Install

```bash
go get github.com/snowmerak/errors
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/snowmerak/errors"
)

func main() {
    err := errors.New("error message")
    fmt.Println(err.Error()) // error message
}
```
