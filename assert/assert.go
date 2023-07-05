package assert

import (
	"github.com/GostGrimmy/gdl/er"
	"github.com/pkg/errors"
)

func Nil(err error, msg string) {
	if err != nil {
		newError := er.NewError(err, msg)
		panic(newError)
	}
}
func True(b bool, msg string) {
	if !b {
		err := er.NewError(errors.New(msg), msg)
		panic(err)
	}
}
func TrueWithData(b bool, msg string, data any) {
	if !b {
		err := er.NewErrorWithData(errors.New(msg), msg, data)
		panic(err)
	}
}
