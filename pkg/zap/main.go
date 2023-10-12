package main

import (
	"fmt"
	"math"

	"go.uber.org/zap"
)

func main() {
	field := zap.Any("test", 0.5)
	fmt.Println(field)
	f := zap.Float64("2", 0.5)
	fmt.Println(f)
	f = zap.Float32("2", 0.5)
	fmt.Println(f)

	p, _ := zap.NewProduction()
	p.Info("test", f)
	p.Info("test", zap.Any("v", 0.5))

	intNumber := math.Float32bits(0.5)
	float32Number := math.Float32frombits(intNumber)
	fmt.Println(intNumber, float32Number)
}

// 		enc.AddFloat32(f.Key, math.Float32frombits(uint32(f.Integer)))

// The IEEE Standard for Floating-Point Arithmetic (IEEE 754) is a technical standard for floating-point arithmetic established in 1985 by the Institute of Electrical and Electronics Engineers (IEEE). The standard addressed many problems found in the diverse floating-point implementations that made them difficult to use reliably and portably. Many hardware floating-point units use the IEEE 754 standard.
// https://en.wikipedia.org/wiki/IEEE_754

// 浮点数的二进制表示: https://www.ruanyifeng.com/blog/2010/06/ieee_floating-point_representation.html
