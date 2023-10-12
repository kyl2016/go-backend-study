package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("key", "value")
	buf, _ := json.Marshal(m) // sync.Map 无法序列化和反序列化，没有 exported fields
	fmt.Println(string(buf))
}
