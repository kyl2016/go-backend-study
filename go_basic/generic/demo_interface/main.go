package main

import "fmt"

func main() {
	fmt.Println(GetID(aModel{}))
	fmt.Println(GetID(bModel{}))
}

type common interface {
	GetID() string
}

type aModel struct {
}

func (a aModel) GetID() string {
	return "a1"
}

type bModel struct {
}

func (b bModel) GetID() string {
	return "b1"
}

func GetID(T common) string {
	return T.GetID()
}
