package repository

import (
	"context"
	"github.com/paysuper/paysuper-recurring-repository/pkg/constant"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/entity"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
	mongodb "gopkg.in/paysuper/paysuper-database-mongo.v1"
	"testing"
	"time"
)

type RepositoryTestSuite struct {
	suite.Suite
	repository *Repository
}

func Test_Repository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) SetupTest() {
	db, err := mongodb.NewDatabase()
	if err != nil {
		suite.FailNow("db connection failed", "%v", err)
	}

	suite.repository = NewRepositoryService(db)
}

func (suite *RepositoryTestSuite) TearDownTest() {
	if err := suite.repository.db.Drop(); err != nil {
		suite.FailNow("db deletion failed", "%v", err)
	}

	_ = suite.repository.db.Close()
}

func (suite *RepositoryTestSuite) TestRepository_InsertSavedCard_Ok() {
	expireYear := time.Now().AddDate(1, 0, 0)

	req := &repository.SavedCardRequest{
		Token:       bson.NewObjectId().Hex(),
		ProjectId:   bson.NewObjectId().Hex(),
		MerchantId:  bson.NewObjectId().Hex(),
		MaskedPan:   "400000******0002",
		CardHolder:  "MR. CARD HOLDER",
		RecurringId: bson.NewObjectId().Hex(),
		Expire: &entity.CardExpire{
			Month: "12",
			Year:  expireYear.Format("2006"),
		},
	}
	rsp := &repository.Result{}
	err := suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	cursor, err := suite.repository.db.Collection(constant.CollectionSavedCard).Find(context.TODO(), bson.M{"token": req.Token})
	assert.NoError(suite.T(), err)

	var cards []*entity.SavedCard
	err = cursor.All(context.TODO(), &cards)
	assert.NoError(suite.T(), err)

	assert.Len(suite.T(), cards, 1)
	assert.Equal(suite.T(), req.ProjectId, cards[0].ProjectId)
	assert.Equal(suite.T(), req.MerchantId, cards[0].MerchantId)
	assert.Equal(suite.T(), req.MaskedPan, cards[0].MaskedPan)
	assert.Equal(suite.T(), req.CardHolder, cards[0].CardHolder)
	assert.Equal(suite.T(), req.RecurringId, cards[0].RecurringId)
	assert.Equal(suite.T(), req.Expire.Month, cards[0].Expire.Month)
	assert.Equal(suite.T(), req.Expire.Year, cards[0].Expire.Year)
}

func (suite *RepositoryTestSuite) TestRepository_InsertSavedCard_SavedCardExist_Ok() {
	expireYear := time.Now().AddDate(1, 0, 0)

	req := &repository.SavedCardRequest{
		Token:       bson.NewObjectId().Hex(),
		ProjectId:   bson.NewObjectId().Hex(),
		MerchantId:  bson.NewObjectId().Hex(),
		MaskedPan:   "400000******0002",
		CardHolder:  "MR. CARD HOLDER",
		RecurringId: bson.NewObjectId().Hex(),
		Expire: &entity.CardExpire{
			Month: "12",
			Year:  expireYear.Format("2006"),
		},
	}
	rsp := &repository.Result{}
	err := suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	cursor, err := suite.repository.db.Collection(constant.CollectionSavedCard).Find(context.TODO(), bson.M{"token": req.Token})
	assert.NoError(suite.T(), err)

	var cards []*entity.SavedCard
	err = cursor.All(context.TODO(), &cards)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), cards, 1)
}

func (suite *RepositoryTestSuite) TestRepository_InsertSavedCard_InsertError() {
	expireYear := time.Now().AddDate(1, 0, 0)

	req := &repository.SavedCardRequest{
		Token:       bson.NewObjectId().Hex(),
		ProjectId:   "some_not_mongodb_object_id_string",
		MerchantId:  bson.NewObjectId().Hex(),
		MaskedPan:   "400000******0002",
		CardHolder:  "MR. CARD HOLDER",
		RecurringId: bson.NewObjectId().Hex(),
		Expire: &entity.CardExpire{
			Month: "12",
			Year:  expireYear.Format("2006"),
		},
	}
	rsp := &repository.Result{}
	err := suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), constant.ErrDatabase, err)
}

func (suite *RepositoryTestSuite) TestRepository_DeleteSavedCard_Ok() {
	expireYear := time.Now().AddDate(1, 0, 0)

	req := &repository.SavedCardRequest{
		Token:       bson.NewObjectId().Hex(),
		ProjectId:   bson.NewObjectId().Hex(),
		MerchantId:  bson.NewObjectId().Hex(),
		MaskedPan:   "400000******0002",
		CardHolder:  "MR. CARD HOLDER",
		RecurringId: bson.NewObjectId().Hex(),
		Expire: &entity.CardExpire{
			Month: "12",
			Year:  expireYear.Format("2006"),
		},
	}
	rsp := &repository.Result{}
	err := suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	var card *entity.SavedCard
	err = suite.repository.db.Collection(constant.CollectionSavedCard).FindOne(context.TODO(), bson.M{"token": req.Token}).Decode(&card)
	assert.NotNil(suite.T(), card)
	assert.True(suite.T(), card.IsActive)

	req1 := &repository.DeleteSavedCardRequest{
		Id:    card.Id,
		Token: req.Token,
	}
	rsp1 := &repository.DeleteSavedCardResponse{}
	err = suite.repository.DeleteSavedCard(context.TODO(), req1, rsp1)
	assert.NoError(suite.T(), err)

	err = suite.repository.db.Collection(constant.CollectionSavedCard).FindOne(context.TODO(), bson.M{"token": req.Token}).Decode(&card)
	assert.NotNil(suite.T(), card)
	assert.False(suite.T(), card.IsActive)
}

func (suite *RepositoryTestSuite) TestRepository_DeleteSavedCard_CardNotExist_Ok() {
	req1 := &repository.DeleteSavedCardRequest{
		Id:    bson.NewObjectId().Hex(),
		Token: bson.NewObjectId().Hex(),
	}
	rsp1 := &repository.DeleteSavedCardResponse{}
	err := suite.repository.DeleteSavedCard(context.TODO(), req1, rsp1)
	assert.NoError(suite.T(), err)
}

func (suite *RepositoryTestSuite) TestRepository_FindSavedCards_Ok() {
	expireYear := time.Now().AddDate(1, 0, 0)

	req := &repository.SavedCardRequest{
		Token:       bson.NewObjectId().Hex(),
		ProjectId:   bson.NewObjectId().Hex(),
		MerchantId:  bson.NewObjectId().Hex(),
		MaskedPan:   "400000******0002",
		CardHolder:  "MR. CARD HOLDER",
		RecurringId: bson.NewObjectId().Hex(),
		Expire: &entity.CardExpire{
			Month: "12",
			Year:  expireYear.Format("2006"),
		},
	}
	rsp := &repository.Result{}
	err := suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0003"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0004"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0005"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0006"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req1 := &repository.SavedCardRequest{Token: req.Token}
	rsp1 := &repository.SavedCardList{}
	err = suite.repository.FindSavedCards(context.TODO(), req1, rsp1)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), rsp1.SavedCards, 5)
	assert.Equal(suite.T(), req.MaskedPan, rsp1.SavedCards[len(rsp1.SavedCards)-1].MaskedPan)
}

func (suite *RepositoryTestSuite) TestRepository_FindSavedCards_UserNotHaveSavedCards_Ok() {
	req1 := &repository.SavedCardRequest{Token: bson.NewObjectId().Hex()}
	rsp1 := &repository.SavedCardList{}
	err := suite.repository.FindSavedCards(context.TODO(), req1, rsp1)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), rsp1.SavedCards, 0)
}

func (suite *RepositoryTestSuite) TestRepository_FindSavedCardById_Ok() {
	expireYear := time.Now().AddDate(1, 0, 0)

	req := &repository.SavedCardRequest{
		Token:       bson.NewObjectId().Hex(),
		ProjectId:   bson.NewObjectId().Hex(),
		MerchantId:  bson.NewObjectId().Hex(),
		MaskedPan:   "400000******0002",
		CardHolder:  "MR. CARD HOLDER",
		RecurringId: bson.NewObjectId().Hex(),
		Expire: &entity.CardExpire{
			Month: "12",
			Year:  expireYear.Format("2006"),
		},
	}
	rsp := &repository.Result{}
	err := suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0003"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0004"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0005"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	req.MaskedPan = "400000******0006"
	err = suite.repository.InsertSavedCard(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)

	cursor, err := suite.repository.db.Collection(constant.CollectionSavedCard).Find(context.TODO(), bson.M{"token": req.Token})
	assert.NoError(suite.T(), err)

	var cards []*entity.SavedCard
	err = cursor.All(context.TODO(), &cards)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), cards, 5)

	req1 := &repository.FindByStringValue{Value: cards[3].Id}
	rsp1 := &entity.SavedCard{}
	err = suite.repository.FindSavedCardById(context.TODO(), req1, rsp1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), cards[3].Token, rsp1.Token)
	assert.Equal(suite.T(), cards[3].ProjectId, rsp1.ProjectId)
	assert.Equal(suite.T(), cards[3].MerchantId, rsp1.MerchantId)
	assert.Equal(suite.T(), cards[3].MaskedPan, rsp1.MaskedPan)
	assert.Equal(suite.T(), cards[3].CardHolder, rsp1.CardHolder)
	assert.Equal(suite.T(), cards[3].RecurringId, rsp1.RecurringId)
	assert.Equal(suite.T(), cards[3].Expire.Month, rsp1.Expire.Month)
	assert.Equal(suite.T(), cards[3].Expire.Year, rsp1.Expire.Year)
}

func (suite *RepositoryTestSuite) TestRepository_FindSavedCardById_NotFound_Error() {
	req1 := &repository.FindByStringValue{Value: bson.NewObjectId().Hex()}
	rsp1 := &entity.SavedCard{}
	err := suite.repository.FindSavedCardById(context.TODO(), req1, rsp1)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), constant.ErrNotFound, err)
}
