package assert

import (
	"github.com/pkg/errors"
	"go-develop-lib/er"
)

func Nil(err error, errorType er.ErrorType, msg string) {
	if err != nil {
		newError := er.NewError(err, msg, errorType)
		panic(newError)
	}
}
func True(b bool, errorType er.ErrorType, msg string) {
	if !b {
		ginError := er.NewError(errors.New(msg), msg, errorType)
		panic(ginError)
	}
}