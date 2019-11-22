module github.com/paysuper/paysuper-recurring-repository

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/ProtocolONE/go-micro-plugins v0.3.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/protobuf v1.3.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/micro/go-micro v1.8.0
	github.com/micro/go-plugins v1.2.0
	github.com/paysuper/paysuper-billing-server v0.0.0-20191031115520-5d2419c5f2cf
	github.com/paysuper/paysuper-database-mongo v0.1.1
	github.com/prometheus/client_golang v1.1.0
	github.com/stretchr/testify v1.4.0
	go.uber.org/zap v1.10.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

go 1.13
