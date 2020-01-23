module github.com/paysuper/paysuper-recurring-repository

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/InVisionApp/go-logger v1.0.1 // indirect
	github.com/ProtocolONE/go-micro-plugins v0.3.0
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/groupcache v0.0.0-20191002201903-404acd9df4cc // indirect
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.5.4 // indirect
	github.com/hashicorp/go-rootcerts v1.0.1 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hashicorp/serf v0.8.3 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.2.0
	github.com/mongodb/mongo-go-driver v0.3.0
	github.com/nats-io/nats-server/v2 v2.1.0 // indirect
	github.com/paysuper/paysuper-proto/go/billingpb v0.0.0-20200122212511-cb73c60b18e4
	github.com/paysuper/paysuper-proto/go/recurringpb v0.0.0-20200122212511-cb73c60b18e4
	github.com/posener/complete v1.2.1 // indirect
	github.com/prometheus/client_golang v1.2.1
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.2.1
	go.uber.org/zap v1.13.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/paysuper/paysuper-database-mongo.v1 v1.0.0-20191120092306-dc35c6f924f1
)

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
)

go 1.13
