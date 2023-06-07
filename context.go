package gdl

type HandleAble interface {
}

type R struct {
	Msg  string
	Code int `json:",omitempty"`
	Ok   bool
	Data any
}

func NewR(msg string, ok bool, data any) *R {
	return &R{Msg: msg, Ok: ok, Data: data}
}

func Ok(msg string, data any) *R {
	return NewR(msg, true, data)
}
func OkWithMsg(data any) *R {
	return Ok("success", data)
}
func FailWithCode(code int, msg string) *R {
	r := NewR(msg, false, nil)
	r.Code = code
	return r
}
func Fail(msg string) *R {
	return NewR("fail", false, nil)
}
