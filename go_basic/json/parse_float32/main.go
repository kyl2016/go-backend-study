package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"verify_condition":{"withdraw_max":0.1,"filter_ratio":1}}`

	var c config
	err := json.Unmarshal([]byte(s), &c)
	fmt.Println(err, c)
}

type config struct {
	Condition AntiAddictionVerifyCondition `json:"verify_condition"`
}

type AntiAddictionVerifyCondition struct {
	WithdrawMax float32 `json:"withdraw_max"` // 用户总提现最大值
	FilterRatio float32 `json:"filter_ratio"` // 超过总提现最大值后筛选的用户比例
}
