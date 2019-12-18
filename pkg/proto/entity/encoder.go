package entity

import (
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MgoExpire struct {
	Month string `bson:"month" json:"month"`
	Year  string `bson:"year" json:"year"`
}

type MgoSavedCard struct {
	Id          primitive.ObjectID `bson:"_id"`
	Token       string             `bson:"token"`
	ProjectId   primitive.ObjectID `bson:"project_id"`
	MerchantId  primitive.ObjectID `bson:"merchant_id"`
	MaskedPan   string             `bson:"masked_pan"`
	CardHolder  string             `bson:"card_holder"`
	Expire      *MgoExpire         `bson:"expire"`
	RecurringId string             `bson:"recurring_id"`
	IsActive    bool               `bson:"is_active"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (m *SavedCard) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		return nil, err
	}

	projectOid, err := primitive.ObjectIDFromHex(m.ProjectId)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(m.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoSavedCard{
		Id:         oid,
		Token:      m.Token,
		ProjectId:  projectOid,
		MerchantId: merchantOid,
		MaskedPan:  m.MaskedPan,
		CardHolder: m.CardHolder,
		Expire: &MgoExpire{
			Month: m.Expire.Month,
			Year:  m.Expire.Year,
		},
		RecurringId: m.RecurringId,
		IsActive:    m.IsActive,
	}

	t, err := ptypes.Timestamp(m.CreatedAt)

	if err != nil {
		return nil, err
	}

	st.CreatedAt = t

	t, err = ptypes.Timestamp(m.UpdatedAt)

	if err != nil {
		return nil, err
	}

	st.UpdatedAt = t

	return bson.Marshal(st)
}

func (s *SavedCard) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoSavedCard)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	s.Id = decoded.Id.Hex()
	s.Token = decoded.Token
	s.ProjectId = decoded.ProjectId.Hex()
	s.MerchantId = decoded.MerchantId.Hex()
	s.MaskedPan = decoded.MaskedPan
	s.CardHolder = decoded.CardHolder
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

	s.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
