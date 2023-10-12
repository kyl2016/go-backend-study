package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	BaseInfo `json:",inline"`
}

type BaseInfo struct {
	Name string `json:"name"`
}

func main() {
	p := People{BaseInfo: BaseInfo{Name: "Bob"}}
	d, _ := json.Marshal(p)
	fmt.Println(string(d))
}
