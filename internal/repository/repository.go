package repository

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/globalsign/mgo/bson"
)

const (
	FieldNameId = "_id"

	QueryErrorMask = "Query from table \"%s\" ended with error: %s"

	CollectionOrder = "order"
)

type Repository struct {
	Database *database.Source
}

func ByteToObjectId(b []byte) bson.ObjectId {
	return bson.ObjectId(b[:])
}

func ObjectIdToByte(id bson.ObjectId) []byte  {
	return []byte(string(id))
}
