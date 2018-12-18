package tools

import "github.com/globalsign/mgo/bson"

type DatabaseHelper struct {}

func (d *DatabaseHelper) ByteToObjectId(b []byte) bson.ObjectId {
	return bson.ObjectId(b[:])
}

func (d *DatabaseHelper) ObjectIdToByte(id bson.ObjectId) []byte {
	return []byte(string(id))
}
