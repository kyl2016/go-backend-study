package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/antonmedv/expr"
)

func random_discrete(numbers interface{}, weights interface{}) string {
	fmt.Println(numbers, weights)
	return "ok"
}

func Test1(t *testing.T) {
	//random_discrete[[1,2,5,10],[0.1,0.3,0.4,0.2]]

	var arrI []interface{}
	json.Unmarshal([]byte("[[1,2,5,10],[0.1,0.3,0.4,0.2]]"), &arrI)
	fmt.Println(arrI)

	var arr []int
	json.Unmarshal([]byte("[1,2,5,10]"), &arr)
	fmt.Println(arr)

	//s := `{"values":[1, 2, 5, 10], "weights":[0.1, 0.3, 0.4, 0.2]}`

	env := map[string]interface{}{
		"name":    random_discrete, // "random_discrete",
		"values":  []int{1, 2, 5, 10},
		"weights": []float32{0.1, 0.3, 0.4, 0.2},
		//"discrete": random_discrete,
	}

	code := `name(values, weights)`

	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

func Test2(t *testing.T) {
	fmt.Println(expr.Eval("is_vpn==true&&count>0", map[string]interface{}{
		"is_vpn": true,
		"count":  1,
	}))
	fmt.Println(expr.Eval("is_vpn", map[string]interface{}{
		"is_vpn": true,
		"count":  1,
	}))
}
