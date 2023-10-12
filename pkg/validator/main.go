package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type rawLotteryRule struct {
	ID             int                `json:"id"`
	Rewards        []rawLotteryReward `json:"rewards" validate:"required,min=1,dive"` // use dive to sub
	Weight         int                `json:"weight" validate:"required,min=0"`
	DefaultReward  rawLotteryReward   `json:",inline" validate:"omitempty"`
	DefaultReward2 rawLotteryReward   `json:"default" validate:"omitempty"`
}

type rawLotteryReward struct {
	Name   string `json:"name" validate:"required"`
	Amount int    `json:"amount" validate:"required,min=0"`
}

// type multiDimensionSlice [][]string `json:"ms" validate:"gt=0,dive,dive,required"`	// with validation tag "gt=0,dive,dive,required"

func main() {
	validate := validator.New()

	rule1 := rawLotteryRule{
		ID: 10,
		Rewards: []rawLotteryReward{
			{Name: "coin", Amount: 10},
			{Name: "hint", Amount: 20},
		},
		Weight:        20,
		DefaultReward: rawLotteryReward{"coin", 10},
	}
	fmt.Println("error:", validate.Struct(rule1))
}

func validateVariable() {
	validate := validator.New()

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
