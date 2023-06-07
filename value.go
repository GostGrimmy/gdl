package gdl

import (
	"go-develop-lib/util"
)

type GetInterface interface {
	Get(s string) (interface{}, bool)
}

func GetK[K any](c GetInterface, key string) (k K, bo bool) {
	data, b := c.Get(key)
	bo = b
	if !b {
		return
	}
	k = util.Convert[K](data)
	return
}
