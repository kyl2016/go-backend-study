package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	expr "github.com/antonmedv/expr"
)

// {{random_discrete[[1,2,5,10],[0.1,0.3,0.4,0.2]]}}

func main() {
	env := map[string]interface{}{
		"x":               []interface{}{1, 2, 5, 10},
		"y":               []float64{0.1, 0.3, 0.4, 0.1},
		"random_discrete": randomDiscrete,
	}

	code := `random_discrete(x, y)`
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

func randomDiscrete(values []interface{}, weights []float64) interface{} {
	weightsInt := make([]int, 0, len(weights))
	for _, it := range weights {
		weightsInt = append(weightsInt, int(it*1000))
	}

	ranges, err := GetRandomRanges(weightsInt)
	if err != nil {
		fmt.Println(err)
	}

	index := RandomSelect(ranges)

	return values[index]
}

func GetRandomRanges(weights []int) ([]Range, error) {
	if len(weights) == 0 {
		return nil, errors.New("weights is null")
	}

	currentWeightAmount := 0
	for _, w := range weights {
		if w < 0 {
			return nil, errors.New("weight should bigger or equal than 0")
		}
		currentWeightAmount += w
	}
	if currentWeightAmount <= 0 {
		return nil, fmt.Errorf("weights amount is %d, should bigger than 0", currentWeightAmount)
	}

	ranges := make([]Range, 0, len(weights))
	from, to := 0, 0
	for _, w := range weights {
		to += int(math.Round(float64(w) / float64(currentWeightAmount) * float64(weightAmount)))
		ranges = append(ranges, Range{
			From: from,
			To:   to,
		})
		from = to
	}
	ranges[len(ranges)-1].To = 100
	return ranges, nil
}

type Range struct {
	From int
	To   int
}

const weightAmount = 100

func RandomSelect(ranges []Range) int {
	lucky := rand.Int31n(weightAmount)
	for i, it := range ranges {
		if it.From <= int(lucky) && int(lucky) < it.To {
			return i
		}
	}

	return -1
}

func main2() {
	env := map[string]interface{}{
		"greet":   "Hello, %v!",
		"names":   []string{"world", "you"},
		"sprintf": fmt.Sprintf,
	}

	code := `sprintf(greet, names[0])`

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
