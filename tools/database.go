package tools

import (
	"math"
	"strings"
)

const (
	panMaskFirstSymbolsCount = 6
	panMaskLastSymbolsCount  = 4
	panMaskedSymbol          = "*"
	precision                = 6
)

func FormatAmount(amount float64) float64 {
	return math.Floor(amount*100) / 100
}

func MaskBankCardNumber(pan string) string {
	rSymCount := len(pan) - (panMaskFirstSymbolsCount + panMaskLastSymbolsCount)

	return pan[:panMaskFirstSymbolsCount] + strings.Repeat(panMaskedSymbol, rSymCount) + pan[rSymCount+panMaskFirstSymbolsCount:]
}

func ToPrecise(val float64) float64 {
	p := math.Pow(10, precision)
	return math.Round(val*p) / p
}

func GetPercentPartFromAmount(amount, rate float64) float64 {
	return amount / (1 + rate) * rate
}
