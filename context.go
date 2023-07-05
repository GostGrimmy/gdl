package gdl

import "fmt"

type R struct {
	Msg  string
	Data any
}

func NewR(msg string, data any) R {
	return R{
		msg, data,
	}
}
func NewSR(format string, value string, data any) R {
	return NewR(fmt.Sprintf(format, value), data)
}
