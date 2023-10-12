# nil

## nil is a predeclared identifier in go.

## cannot set 'nil' to string

The nil constant has **no type** — so it cannot substitute for a string.

    package main
    
    func main() {
        temp := []string{}
        temp = append(temp, nil)
    }
    
    Output
    
    # command-line-arguments
    C:\programs\file.go:8: cannot use nil as type string in append

## nil can represent zero values of many types

- pointer types (including type-unsafe ones)
- map types
- slice types
- function types
- channel types
- interface types

## predeclared nil is not a keyword in Go

The predecalred nil can be shadowed.

```
nil := 123
fmt.Println(nil) // 123

// The dollowing line fails to compile,
// for nil represents an int value now 
// in this scope.
var _ map[string]int = nil
```

## the size of nil values with types of different kinds may be different

## two nil values of two different types may be not comparable

```
// Compilation failure reason: mismatched types.
// neither operand can be implicitly converted to the type of the other.
var _ = (*int)(nil) == (*bool)(nil) // error
var _ = (chan int)(nil) == (chan bool)(nil) // error
```

```
type IntPtr *int
// The underlying of type IntPtr is *int.
var _ = IntPtr(nil) == (*int)(nil)

// Every type in Go implements interface{} type.
var _ = (interface{})(nil) == (*int)(nil)

// Values of a directional channel type can be converted to the bidirectional channel type which has the same element type.
var _ = (chan int)(nil) == (chan<- int)(nil)
var _ = (chan int)(nil) == (<-chan int)(nil)
```

## two nil values of the same type may be not comparable

```
// illegal
var _ = ([]int)(nil) == ([]int)(nil)
var _ = (map[string]int)(nil) == (map[string]int)(nil)
var _ = (func())(nil) == (func())(nil)
```

But any types of the above mentioned incomparable types can be compared with
the bare nil identifier.
```
// The following lines compile okay.
var _ = ([]int)(nil) == nil
var _ = (map[string]int)(nil) == nil
var _ = (func())(nil) == nil
```

## ！！！ two nil values may be not equal

If one of the two compared nil values is an interface{} value and the other is not,
assume they are comparable, then the comparison result is always false. The
reson is the not-interface value will be [converted to the type of interface value](https://go101.org/article/interface.html#boxing)
before making the comparison. **The converted interface value has a concrete
dynamic type but the other interface value has not**.

```
fmt.Println( (interface{})(nil) == (*int)(nil) ) // false
```

## Retrieving elements from nil maps will not panic

```
fmt.Println( (map[string]int)(nil)["key"] ) // 0
fmt.Println( (map[int]bool)(nil)[123]) // false
fmt.Println( (map[int]*int64)(nil)[123]) // <nil>
```

## predeclared nil has not a default type

In fact, the predecared nil is the only untyped value who has not a default type in Go. There must be sufficient information for compiler to deduce(推断) the type of a nil from context.

in go, for simplicity and convenience, nil is designed as an identifier which can be used to represent the zero values of some kinds of types. it is **not a single value**. it can represent many values with different memory layouts.

## faq

Why is my nil error value not equal to nil? ¶
Under the covers, interfaces are implemented as two elements, a type T and a value V. V is a concrete value such as an int, struct or pointer, never an interface itself, and has type T. For instance, if we store the int value 3 in an interface, the resulting interface value has, schematically, (T=int, V=3). The value V is also known as the interface's dynamic value, since a given interface variable might hold different values V (and corresponding types T) during the execution of the program.

An interface value is nil only if the V and T are both unset, (T=nil, V is not set), In particular, a nil interface will always hold a nil type. If we store a nil pointer of type *int inside an interface value, the inner type will be *int regardless of the value of the pointer: (T=*int, V=nil). Such an interface value will therefore be non-nil even when the pointer value V inside is nil.

This situation can be confusing, and arises when a nil value is stored inside an interface value such as an error return:

```
func returnsError() error {
var p *MyError = nil
if bad() {
p = ErrBad
}
return p // Will always return a non-nil error.
}
If all goes well, the function returns a nil p, so the return value is an error interface value holding (T=*MyError, V=nil). This means that if the caller compares the returned error to nil, it will always look as if there was an error even if nothing bad happened. To return a proper nil error to the caller, the function must return an explicit nil:

func returnsError() error {
if bad() {
return ErrBad
}
return nil
}
```
It's a good idea for functions that return errors always to use the error type in their signature (as we did above) rather than a concrete type such as *MyError, to help guarantee the error is created correctly. As an example, os.Open returns an error even though, if not nil, it's always of concrete type *os.PathError.

Similar situations to those described here can arise whenever interfaces are used. Just keep in mind that if any concrete value has been stored in the interface, the interface will not be nil. For more information, see The Laws of Reflection.

## refer

[go101/nil](https://go101.org/article/nil.html)