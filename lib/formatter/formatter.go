package formatter

type Formatter interface {
	Byte(string, byte) (int, error)
	Bytes(string, []byte) (int, error)
	Float64(string, float64) (int, error)
	Int64(string, int64) (int, error)
	Uint64(string, uint64) (int, error)
	String(string, string) (int, error)
	Bool(string, bool) (int, error)
	Err(string, error) (int, error)
	Caller() (int, error)
	Msg(string) (int, error)
}
