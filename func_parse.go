package muxplus

import (
	"reflect"
)

type Func struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type
	Func reflect.Value
}

func FuncParse(function interface{}) (ret *Func) {

	var (
		funcType reflect.Type
		numIn    int
		numOut   int
	)
	ret = new(Func)

	if function == nil {
		panic("function can't be nil")
	}

	funcType = reflect.TypeOf(function)

	if funcType.Kind() != reflect.Func {
		panic("function must be func")
	}

	ret.Name = funcType.Name()

	numIn = funcType.NumIn()
	numOut = funcType.NumOut()

	for i := 0; i < numIn; i++ {
		ret.In = append(ret.In, funcType.In(i))
	}

	for i := 0; i < numOut; i++ {
		ret.Out = append(ret.Out, funcType.Out(i))
	}

	ret.Func = reflect.ValueOf(function)

	return
}
