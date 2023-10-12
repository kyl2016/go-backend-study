package random

import (
	"errors"
	"github.com/kyl2016/Play-With-Golang/utility"
	"math"
	"math/rand"
	"time"
)

const weightAmount int32 = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomNumber(max int) int {
	return rand.Intn(max + 1)
}

// [From, To)
type Range struct {
	From int
	To   int
}

func randSelect(ranges []Range) int {
	lucky := rand.Int31n(weightAmount)
	for i, it := range ranges {
		if it.From <= int(lucky) && int(lucky) < it.To {
			return i
		}
	}

	return -1
}

func getRandomRanges(weights []int) ([]Range, error) {
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
		return nil, errors.New("sum weights should bigger than 0")
	}

	var ranges []Range
	from, to := 0, 0
	for _, w := range weights {
		to += int(math.Round(float64(w) / float64(currentWeightAmount) * float64(weightAmount)))
		ranges = append(ranges, Range{
			From: from,
			To:   to,
		})
		from = to
	}
	if ranges[len(ranges)-1].To != 100 {
		ranges[len(ranges)-1].To = 100
	}
	return ranges, nil
}

func Discrete(values []interface{}, weights []float32) interface{} {
	var weightsInt []int
	for _, it := range weights {
		weightsInt = append(weightsInt, int(it*1000))
	}
	ranges, err := getRandomRanges(weightsInt)
	utility.PanicIfNotNil(err)

	index := randSelect(ranges)
	//fmt.Println(values[index])

	return values[index]
}
