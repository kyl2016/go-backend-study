package main

import (
	"encoding/json"
	"os"
)

func main() {
	truncateAppend(map[string]string{"hi": "ok"}, "temp.txt")

	append2()
}

func truncateAppend(info interface{}, saveTo string) error {
	buf, _ := json.Marshal(info)
	// WriteFile writes data to the named file, creating it if necessary.
	// If the file does not exist, WriteFile creates it with permissions perm (before umask);
	// otherwise WriteFile truncates it before writing, without changing permissions.
	return os.WriteFile(saveTo, buf, os.ModeAppend|os.ModePerm)
}

func append2() {
	reader, err := os.OpenFile("temp.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	reader.WriteString("hello")
	reader.Close()
}
