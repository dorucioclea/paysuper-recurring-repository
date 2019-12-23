module github.com/paysuper/paysuper-recurring-repository

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/ProtocolONE/go-micro-plugins v0.3.0
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/protobuf v1.3.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/micro/go-micro v1.8.0
	github.com/micro/go-plugins v1.2.0
	github.com/paysuper/paysuper-billing-server v1.0.0
	github.com/paysuper/paysuper-database-mongo v0.1.3
	github.com/prometheus/client_golang v1.2.1
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.1.3
	go.uber.org/zap v1.10.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/paysuper/paysuper-database-mongo.v1 v1.0.0-20191120092306-dc35c6f924f1
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

go 1.13
