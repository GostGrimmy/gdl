package er

type ErrorType uint64

const (
	BindingErrType ErrorType = 9
	ParamErrType   ErrorType = 10
	InterErrType   ErrorType = 11
	NoLoginErrType           = InterErrType + 2
	//以下是给管理员看的
	SQlErrType        ErrorType = NoLoginErrType + 2
	PwdResolveErrType           = SQlErrType + 2
	PwdErrType                  = PwdResolveErrType + 2
	ServiceErrType              = PwdErrType + 2
)
const (
	NotFrontSee bool = false
	FrontSee    bool = true
)

var ErrorMap = map[ErrorType]string{
	ParamErrType:      "参数错误",
	InterErrType:      "服务器异常，请稍后再试",
	NoLoginErrType:    "未登录",
	SQlErrType:        "sql执行错误",
	PwdResolveErrType: "密码解析错误，请重新确认",
	PwdErrType:        "密码错误",
}

type Error struct {
	Err  error
	Msg  string //前端可看消息
	Type ErrorType
}

func (e Error) Error() string {
	return e.Err.Error()
}

func NewError(err error, msg string, t ErrorType) *Error {
	return &Error{
		Err:  err,
		Msg:  msg,
		Type: t,
	}
}
