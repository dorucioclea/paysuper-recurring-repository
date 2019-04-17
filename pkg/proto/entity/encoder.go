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
    Token       string        `bson:"token"`
    ProjectId   bson.ObjectId `bson:"project_id"`
    MerchantId  bson.ObjectId `bson:"merchant_id"`
    MaskedPan   string        `bson:"masked_pan"`
    Expire      *MgoExpire    `bson:"expire"`
    RecurringId string        `bson:"recurring_id"`
    IsActive    bool          `bson:"is_active"`
    CreatedAt   time.Time     `bson:"created_at"`
    UpdatedAt   time.Time     `bson:"updated_at"`
}

func (m *SavedCard) GetBSON() (interface{}, error) {
    st := &MgoSavedCard{
        Id:         bson.ObjectIdHex(m.Id),
        Token:      m.Token,
        ProjectId:  bson.ObjectIdHex(m.ProjectId),
        MerchantId: bson.ObjectIdHex(m.MerchantId),
        MaskedPan:  m.MaskedPan,
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

    return st, nil
}

func (s *SavedCard) SetBSON(raw bson.Raw) error {
    decoded := new(MgoSavedCard)
    err := raw.Unmarshal(decoded)

    if err != nil {
        return err
    }

    s.Id = decoded.Id.Hex()
    s.Token = decoded.Token
    s.ProjectId = decoded.ProjectId.Hex()
    s.MerchantId = decoded.MerchantId.Hex()
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

    s.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

    if err != nil {
        return err
    }

    return nil
}
