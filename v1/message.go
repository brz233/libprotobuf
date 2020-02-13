package v1

type Message interface {
	Bytes() byte
}

type Record struct {
	Index uint64
	Name  string
	Value interface{}
}

type Field struct {
	Index uint64
}
