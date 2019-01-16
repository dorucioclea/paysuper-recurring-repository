package tools

import (
	"github.com/globalsign/mgo/bson"
	"math"
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