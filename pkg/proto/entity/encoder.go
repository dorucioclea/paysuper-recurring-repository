package entity

import (
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type MgoExpire struct {
	Month string `bson:"month" json:"month"`
	Year  string `bson:"year" json:"year"`
}

type MgoSavedCard struct {
	Id          bson.ObjectId `bson:"_id"`
	Account     string        `bson:"account"`
	ProjectId   bson.ObjectId `bson:"project_id"`
	MaskedPan   string        `bson:"masked_pan"`
	Expire      *MgoExpire    `bson:"expire"`
	RecurringId string        `bson:"recurring_id"`
	IsActive    bool          `bson:"is_active"`
	CreatedAt   time.Time     `bson:"created_at"`
}

func (s *SavedCard) GetBSON() (interface{}, error) {
	st := &MgoSavedCard{
		Id:        bson.ObjectIdHex(s.Id),
		Account:   s.Account,
		ProjectId: bson.ObjectIdHex(s.ProjectId),
		MaskedPan: s.MaskedPan,
		Expire: &MgoExpire{
			Month: s.Expire.Month,
			Year:  s.Expire.Year,
		},
		RecurringId: s.RecurringId,
		IsActive: s.IsActive,
	}

	t, err := ptypes.Timestamp(s.CreatedAt)

	if err != nil {
		return nil, err
	}

	st.CreatedAt = t

	return st, nil
}

func (s *SavedCard) SetBSON(raw bson.Raw) error {
	decoded := new(MgoSavedCard)
	err := raw.Unmarshal(decoded)

	if err != nil {
		return err
	}

	s.Id = decoded.Id.Hex()
	s.Account = decoded.Account
	s.ProjectId = decoded.ProjectId.Hex()
	s.MaskedPan = decoded.MaskedPan
	s.Expire = &CardExpire{
		Month: decoded.Expire.Month,
		Year:  decoded.Expire.Year,
	}
	s.RecurringId = decoded.RecurringId
	s.IsActive = decoded.IsActive
	s.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
