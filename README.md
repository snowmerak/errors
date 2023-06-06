# errors

errors is a simple error handling library for Go.

## install

```bash
go get github.com/snowmerak/errors
```

## features

### Errors

- `New` - create a new error with message
- `Wrap` - wrap some errors into new error
- `Is` - check if error is equal to some error
- `As` - check if error is equal to some error type
- `Unwrap` - get the previous error
- `(e *Errors) Unwrap` - get the previous errors of Errors
- `Join` - join some errors into new Errors

### Structured Errors

- `From` - create a new structured error with message
- `Get` - get the value of some key
