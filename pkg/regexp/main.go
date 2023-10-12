package main

import (
	"fmt"
	"github.com/kyl2016/Play-With-Golang/utility"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	for i := 0; i < 10000; i++ {
		resp, _ := http.Get("kong-test.atcloudbox.com/config")
		fmt.Println(resp.StatusCode)
	}

	reg, _ := regexp.Compile("{{(.*?)}}")
	arr := reg.FindAllStringSubmatch("//\t\"image\": \"{{image}}\",{{image2}}", -1)
	for _, it := range arr {
		for _, v := range it {
			fmt.Println(v)
		}
		fmt.Println()
	}
}

func showlistCreateValueBy(expression string) interface{} {
	if strings.Contains(expression, "random_integer") {
		scope := strings.ReplaceAll(expression, "random_integer", "")
		reg, _ := regexp.Compile("\\d+")
		arr := reg.FindAllString(scope, -1)
		if len(arr) < 2 {
			//	return error
		}
		start, err := strconv.Atoi(arr[0])
		end, err := strconv.Atoi(arr[1])

		utility.PanicIfNotNil(err)

		// TODO end > start

		return int(rand.Int31n(int32(end-start+1))) + start
	}

	panic("not implement")
}
