package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writer := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	r, err := os.Open("testdata/flowers_160.jpg")
	if err != nil {
		panic(err)
	}
	buffer, _ := ioutil.ReadAll(r)
	n, _ := writer.Write(buffer)
	fmt.Println()
	fmt.Println(n)

	//enc := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	//s := enc.EncodeToString(buffer)
	//fmt.Println(s)
}
