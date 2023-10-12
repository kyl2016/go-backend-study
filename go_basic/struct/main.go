package main

import (
	"fmt"
	"reflect"
)

type T struct {
	num int
}

func (t *T) Increase() {
	t.num++
}

func (t T) F1() {
	fmt.Println(t.num)
}

func (t *T) F2() {
	fmt.Println(t.num)
}

type User struct {
	id      int    // 4
	age     int    // 4
	gender  bool   // 1
	name    string // 16
	active  bool   // 1
	address string // 16
}

type User2 struct {
	id      int    // int is int64 in 64bit os
	age     int    // 8 Byte
	name    string // 16
	address string // 16
	gender  bool   // 1
	active  bool   // 1
	valid   bool   // 1
	valid2  bool   // 1
	valid3  bool   // 1
	valid4  bool   // 1
	valid5  bool   // 1
	valid6  bool   // 1
	id2     int
	age2    int
}

func main() {
	fmt.Println(reflect.TypeOf(int32(1)).Size())
	fmt.Println(reflect.TypeOf(int(1)).Size())
	fmt.Println(reflect.TypeOf(1).Size())
	fmt.Println(reflect.TypeOf(User{}).Size())
	fmt.Println(reflect.TypeOf(User2{}).Size())

	var t T
	t.F1()
	t.F2()

	t.Increase()

	(&t).F1()
	(&t).F2()
}
