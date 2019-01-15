package tools

import (
	"math"
)

func FormatAmount(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}