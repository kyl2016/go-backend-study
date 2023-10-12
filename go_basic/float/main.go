package main

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func main() {
	var f1 float32 = 9.90

	fmt.Println(f1 * 100)

	var f2 float64 = 9.90

	fmt.Println(f2 * 100)

	s := strconv.FormatFloat(1234.5678, 'g', 8, 32)

	fmt.Println(s)

	fmt.Println(strconv.ParseFloat(s, 32))

	s = strconv.FormatFloat(1234.5678, 'g', 8, 64)

	fmt.Println(s)

	fmt.Println(strconv.ParseFloat(s, 64))

	main2()
}

func main2() {

	price, err := decimal.NewFromString("136.02")

	if err != nil {

		panic(err)

	}

	quantity := decimal.NewFromFloat(3)

	fee, _ := decimal.NewFromString(".035")

	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal) // Subtotal: 408.06

	fmt.Println("Pre-tax:", preTax) // Pre-tax: 422.3421

	fmt.Println("Taxes:", total.Sub(preTax)) // Taxes: 37.482861375

	fmt.Println("Total:", total) // Total: 459.824961375

	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

}
