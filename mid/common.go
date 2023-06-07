package mid

import (
	"gdl/assert"
	"gdl/er"
	"github.com/gin-gonic/gin"
)

type BindType uint8

const (
	Should     BindType = 0
	Bind                = Should + 1
	BindQuery           = Bind + 1
	BindJson            = BindQuery + 1
	BindHeader          = BindJson + 1
)

// 传入后转成指针
func ParamBind[T any](key string, bindType BindType) gin.HandlerFunc {
	var err error
	return func(c *gin.Context) {
		bindData := new(T)
		switch bindType {
		case Should:
			err = c.ShouldBind(bindData)
		case Bind:
			err = c.Bind(bindData)
		case BindQuery:
			err = c.ShouldBindQuery(bindData)
		case BindJson:
			err = c.ShouldBindJSON(bindData)
		case BindHeader:
			err = c.ShouldBindHeader(bindData)
		}
		assert.Nil(err, er.ParamErrType, "")
		c.Set(key, bindData)
	}
}
func ChangeKeyName(before string, after string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get(before)
		delete(c.Keys, before)
		c.Set(after, user)
	}
}
