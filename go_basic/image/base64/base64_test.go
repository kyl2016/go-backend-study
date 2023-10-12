package base64

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/kyl2016/Play-With-Golang/utility"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteStdout(t *testing.T) {
	input := []byte("foo\x00bar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
}

func TestWriteBuffer(t *testing.T) {
	input := []byte("foo\x00bar")
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)
	encoder.Write(input)
	fmt.Println(string(buffer.Bytes()))
}

// 测试 base64 对图片进行编解码
func TestEncodeAndDecodeImageUseBase64(t *testing.T) {
	buf, err := ioutil.ReadFile("./1.png")
	utility.PanicIfNotNil(err)
	dst := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(dst)

	origin, err := base64.StdEncoding.DecodeString(dst)
	utility.PanicIfNotNil(err)
	err = ioutil.WriteFile("1.origin.png", origin, os.ModePerm)
	utility.PanicIfNotNil(err)
}
