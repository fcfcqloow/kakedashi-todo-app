package ldb

type (
	LDBOptions struct {
		read  *func() ([]byte, error)
		write *func([]byte) error
	}
	LDBOption func(*LDBOptions)
)

func ReadByteOption(read func() ([]byte, error)) LDBOption {
	return func(l *LDBOptions) {
		l.read = &read
	}
}

func WriteByteOption(write func([]byte) error) LDBOption {
	return func(l *LDBOptions) {
		l.write = &write
	}
}
