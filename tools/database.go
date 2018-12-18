package tools

import (
	"github.com/globalsign/mgo/bson"
	"math"
)

type DatabaseHelper struct {}

func (d *DatabaseHelper) ByteToObjectId(b []byte) bson.ObjectId {
	return bson.ObjectId(b[:])
}

func (d *DatabaseHelper) ObjectIdToByte(id bson.ObjectId) []byte {
	return []byte(string(id))
}

func (d *DatabaseHelper) FormatAmount(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}