package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var i = nil // can't assign nil without explicit type

	var err error
	fmt.Println("err type is ", reflect.TypeOf(err), "value is", reflect.ValueOf(err))
	// <nil>

	err = newMyError()
	fmt.Println("err type is ", reflect.TypeOf(err), "value is", reflect.ValueOf(err))
	//  *myError
	if err != nil {
		fmt.Println("err is ", err)
		if reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil() {
			fmt.Println("err is a nil of type ", reflect.TypeOf(err))
		}
	}
	// Go is statically typed. Every variable has a static type, that is, exactly one type known and fixed at compile time
	err2 := newMyError()
	fmt.Println("err2 type is ", reflect.TypeOf(err2), "value is", reflect.ValueOf(err2))
	if err2 != nil {
		fmt.Println("err2 is ", err2)
	}
}

type myError struct {
}

func newMyError() *myError {
	return nil
}

func (m *myError) Error() string {
	return ""
}
