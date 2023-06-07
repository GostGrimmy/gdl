package http

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"reflect"
)

func GetK[K any](c *gin.Context, key string) (K, bool) {
	v, o := c.Get(key)
	var k K
	if !o {
		return k, o
	}
	kV := reflect.ValueOf(k)
	vV := reflect.ValueOf(v)
	if kV.Kind() == reflect.Ptr && vV.Kind() == reflect.Ptr {
		k, _ = v.(K)
	} else if kV.Kind() == reflect.Ptr && vV.Kind() != reflect.Ptr {
		value := reflect.New(reflect.TypeOf(k).Elem())
		value.Elem().Set(vV)
		k = (value.Interface()).(K)
	} else if kV.Kind() != reflect.Ptr && vV.Kind() == reflect.Ptr {
		in := vV.Elem().Interface()
		k, _ = in.(K)
	} else if kV.Kind() != reflect.Ptr && vV.Kind() != reflect.Ptr {
		k, _ = v.(K)
	}
	return k, o
}
func ClearBody(c *gin.Context) {
	writer := c.Writer.(*BodyWriter)
	writer.NewWriter()
}
func GetBody(c *gin.Context) []byte {
	return c.Writer.(BodyWriter).body.Bytes()
}

// 先new 一个writer
// 再writer
// 之后done
type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func NewBodyWriter(c *gin.Context) *BodyWriter {
	return &BodyWriter{c.Writer, &bytes.Buffer{}}
}
func (r BodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b) //r.ResponseWriter.Write(b)
}
func (r *BodyWriter) NewWriter() {
	buffer := &bytes.Buffer{}
	r.body = buffer
}
func (r BodyWriter) Done() {
	r.ResponseWriter.Write(r.body.Bytes())
}
