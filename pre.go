package cat

import "reflect"

var TAB = "\t"
var LF = "\n"

type Panic interface{}

type Function interface{}

//Invoke panics if f's Kind is not Func.
//As accurate validation is skipped for performance concern,
//don't call Invoke unless you know what you're doing.
func Invoke(f Function, values ...interface{}) ([]reflect.Value, error) {
	in := make([]reflect.Value, len(values))
	for i, v := range values {
		in[i] = reflect.ValueOf(v)
	}
	return reflect.ValueOf(f).Call(in), nil
}
