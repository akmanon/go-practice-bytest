package reflectiongo

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Array, reflect.Slice:
		for i := range val.Len() {
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := range val.NumField() {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, keys := range val.MapKeys() {
			walkValue(val.MapIndex(keys))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnRes := val.Call(nil)
		for _, res := range valFnRes {
			walkValue(res)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
		return val
	}
	return val
}
