package er

type ErrorType uint64

type Error struct {
	Err  error
	Msg  string //前端可看消息
	Data any
}

func (e Error) Error() string {
	return e.Err.Error()
}

func NewError(err error, msg string) *Error {
	return &Error{
		Err: err,
		Msg: msg,
	}
}
func NewErrorWithData(err error, msg string, data any) *Error {
	return &Error{
		Err:  err,
		Msg:  msg,
		Data: data,
	}
}
