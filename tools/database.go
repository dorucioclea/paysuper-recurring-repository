package tools

import (
	"math"
)

type DatabaseHelper struct {}

func (d *DatabaseHelper) FormatAmount(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}