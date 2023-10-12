package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	raw := `{"name":"Bob"}`
	fmt.Println(json.Valid([]byte(raw)) == true)
	buf, _ := json.Marshal(json.RawMessage(raw))
	fmt.Println(string(buf))
}
