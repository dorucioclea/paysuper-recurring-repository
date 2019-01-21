package tools

import (
	"github.com/globalsign/mgo/bson"
	"math"
	"strings"
)

const (
	panMaskFirstSymbolsCount = 6
	panMaskLastSymbolsCount  = 4
	panMaskedSymbol          = "*"
)

func FormatAmount(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}

func ByteToObjectId(b []byte) bson.ObjectId {
	return bson.ObjectId(b[:])
}

func ObjectIdToByte(id bson.ObjectId) []byte {
	return []byte(string(id))
}

func MaskBankCardNumber(pan string) string {
	rSymCount := len(pan) - (panMaskFirstSymbolsCount + panMaskLastSymbolsCount)

	return pan[:panMaskFirstSymbolsCount] + strings.Repeat(panMaskedSymbol, rSymCount) + pan[rSymCount + panMaskFirstSymbolsCount:]
}
