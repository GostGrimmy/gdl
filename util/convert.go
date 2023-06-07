package util

import (
	"reflect"
)

func Convert[K any](v any) (k K) {
	return convert[K](reflect.ValueOf(v))
}
func ConvertArray[K any](v any) (k []K) {
	vV := reflect.ValueOf(v)
	if vV.Kind() == reflect.Array {
		var in []reflect.Value
		vV.CallSlice(in)
		for _, value := range in {
			k = append(k, convert[K](value))
		}
	}
	return
}

func convert[K any](vV reflect.Value) (k K) {
	kV := reflect.ValueOf(k)
	v := vV.Interface()
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
	return
}
