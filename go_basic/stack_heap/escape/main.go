package main

func main() {
	x := `package test_lottery

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

const lotteryUrlPrefix = "http://localhost:8080/bp/server/user/user1/lottery"
const lottery001UrlPrefix = "http://localhost:8080/bp/server/user/user1/lottery/lottery0001"

func Test_spin(t *testing.T) {
	reader := strings.NewReader("{}")
	resp, err := Post(lottery001UrlPrefix+"/spin", reader)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, string(data), err)
}

func Post(url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      time.Now().Unix(),
		"key_name": "1d93491bf5154452",
	})
	tokenString, err := token.SignedString([]byte("401538e3282d40f2bd0041fdbc34481f"))
	if err != nil {
		panic(err)
	}
	request.Header.Add("X-BytePower-Auth-Token", tokenString)

	client := http.Client{
		Timeout: time.Minute,
	}
	return client.Do(request)
}

func Test_GetLimits(t *testing.T) {
	resp, err := http.Get(lottery001UrlPrefix + "/limits")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, string(data), err)
}

func Test_GetStats(t *testing.T) {
	resp, err := http.Get(lottery001UrlPrefix + "/stats")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, string(data), err)
}

func Test_GetList(t *testing.T) {
	resp, err := http.Get(lotteryUrlPrefix)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, string(data), err)
}

func Test_Get(t *testing.T) {
	resp, err := http.Get(lottery001UrlPrefix)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, string(data), err)
}

func Test_Get_Not_Exist(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/bp/server/user/user1/lottery/lottery0222221")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, string(data), err)
}

func Test_IncrTimes(t *testing.T) {
	resp, err := Post(lottery001UrlPrefix+"/incr_times", reader)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, string(data), err)
}

// 抽奖次数为 0 时，不断调用 GetLimits，每次得到的 next_recovery_time 是固定值
func Test_Recover(t *testing.T) {
	reader := strings.NewReader("{}")
	for i := 0; i < 20; i++ {
		resp, err := Post(lottery001UrlPrefix+"/spin", reader)
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode, string(data), err)

		time.Sleep(time.Minute * 1)
	}
}
`
	a := &x
	//fmt.Println(x)
	sum(a)
}

//go:noinline
func sum(a *string) string {
	return *a + "\n"
}

// go build -gcflags '-m -m' ./main.go  // verbose escape
