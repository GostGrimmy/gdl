package gdl

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	Before []EngineFunc
	Host   string
	Port   string
}
type EngineFunc func(engine *gin.Engine) error

func (a App) Run() {
	r := gin.New()
	for _, engineFunc := range a.Before {
		err := engineFunc(r)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	err := r.Run(fmt.Sprintf("%s:%s",
		a.Host,
		a.Port))
	if err != nil {
		_ = fmt.Errorf("server run error , the err is %+v", err)
		return
	}
}
