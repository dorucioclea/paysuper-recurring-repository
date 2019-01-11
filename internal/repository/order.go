package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
	"net"
)

func (r *Repository) UpdateOrder(ctx context.Context, req *billing.Order, rsp *repository.Result) error {
	m := r.toMap(req)
	id := m[FieldNameUnderscoreId]

	delete(m, FieldNameUnderscoreId)

	err := r.Database.Collection(CollectionOrder).UpdateId(id, bson.M{"$set": m})

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}

func (r *Repository) InsertOrder(ctx context.Context, req *billing.Order, rsp *repository.Result) error {
	err := r.Database.Collection(CollectionOrder).Insert(req)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindOrderById(ctx context.Context, req *repository.FindByUnderscoreId, rsp *billing.Order) error {
	err := r.Database.Collection(CollectionOrder).Find(req).All(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}

func (r *Repository) GetUserGeoDataByIp(ctx context.Context, req *repository.FindByStringValue, rsp *billing.PayerData) error {
	ip := net.ParseIP(req.GetValue())
	rec, err := r.GeoReader.City(ip)

	if err != nil {
		log.Printf(GeoIpErrorMask, req.GetValue(), err.Error())
		return err
	}

	rsp.Ip = req.GetValue()
	rsp.CountryCodeA2 = rec.Country.IsoCode
	rsp.CountryName = &billing.Name{En: rec.Country.Names["en"], Ru: rec.Country.Names["ru"]}
	rsp.City = &billing.Name{En: rec.City.Names["en"], Ru: rec.City.Names["ru"]}
	rsp.Timezone = rec.Location.TimeZone

	if len(rec.Subdivisions) > 0 {
		rsp.Subdivision = rec.Subdivisions[0].IsoCode
	}

	return nil
}

func (r *Repository) FindOrderByProjectAndOrderId(ctx context.Context, req *repository.FindByProjectOrderId, rsp *billing.Order) error {
	err := r.Database.Collection(CollectionOrder).Find(r.toMap(req)).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}
