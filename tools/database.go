package tools

import (
	"math"
	"strings"
)

const (
	panMaskFirstSymbolsCount = 6
	panMaskLastSymbolsCount  = 4
	panMaskedSymbol          = "*"
	precision                = 10
)

func FormatAmount(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}

func MaskBankCardNumber(pan string) string {
	rSymCount := len(pan) - (panMaskFirstSymbolsCount + panMaskLastSymbolsCount)

	return pan[:panMaskFirstSymbolsCount] + strings.Repeat(panMaskedSymbol, rSymCount) + pan[rSymCount+panMaskFirstSymbolsCount:]
}

func toPrecise(val float64) float64 {
	p := math.Pow(10, precision)
	return math.Round(val*p) / p
}
