package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) Set(name string, age int) string {
	u.Name = name
	u.Age = age

	return u.Name
}

func main() {
	user := User{1, "Kitty", 23}
	getValue := reflect.ValueOf(&user)
	userValue := reflect.ValueOf(user)
	fmt.Println("Name:" + userValue.FieldByName("Name").String())

	method := getValue.MethodByName("Set")
	args := []reflect.Value{reflect.ValueOf("hello kitty"), reflect.ValueOf(22)}
	method.Call(args)
	fmt.Printf("%+v\n", user)

	userValue = reflect.ValueOf(user)
	fmt.Println("Name:" + userValue.FieldByName("Name").String())
	v := userValue.FieldByName("Unknow")
	fmt.Println(v.IsValid())

	m := map[string]string{}
	fmt.Println("name:", m["name"])
}
