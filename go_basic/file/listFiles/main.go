package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
	//	fmt.Println(path)
	//	return err
	//})

	files, _ := ioutil.ReadDir("./")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
