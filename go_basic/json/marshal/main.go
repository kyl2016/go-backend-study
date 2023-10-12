package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	_, err := json.Marshal(math.Inf(1))
	_, ok := err.(*json.UnsupportedValueError) // ok == true
	fmt.Println(ok)

	fmt.Println(json.Marshal([]int{1, 2}))
	fmt.Println(json.Marshal(3))
}

/*
Just to complement Jonathan's answer, the json.Marshal function can return two types of errors: UnsupportedTypeError or UnsupportedValueError

The first one can be caused, as Jonathan said by trying to Marshal an invalid type:

_, err := json.Marshal(make(chan int))
_, ok := err.(*json.UnsupportedTypeError) // ok == true
On the other hand you can also have the Marshal function return an error by passing an invalid value:

_, err := json.Marshal(math.Inf(1))
_, ok := err.(*json.UnsupportedValueError) // ok == true


func newTypeEncoder(t reflect.Type, allowAddr bool) encoderFunc {
    // ignored
    switch t.Kind() {
    case reflect.Bool:
        return boolEncoder
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return intEncoder
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return uintEncoder
    case reflect.Float32:
        return float32Encoder
    case reflect.Float64:
        return float64Encoder
    case reflect.String:
        return stringEncoder
    case reflect.Interface:
        return interfaceEncoder
    case reflect.Struct:
        return newStructEncoder(t)
    case reflect.Map:
        return newMapEncoder(t)
    case reflect.Slice:
        return newSliceEncoder(t)
    case reflect.Array:
        return newArrayEncoder(t)
    case reflect.Ptr:
        return newPtrEncoder(t)
    default:
        return unsupportedTypeEncoder
    }
}

Marshal 不支持的标准类型有 Complex64 ，Complex128 ，Chan ，Func ，UnsafePointer ，这种情况下会返回 UnsupportedTypeError 。对于不支持的数据类型，需要实现 MarshalJSON 或者 encoding.TextMarshaler 接口。对于不支持的值，会返回 UnsupportedValueError 错误，如浮点数的无穷大，无穷小，NaN 和出现循环引用的 map、slice和pointer。

*/
