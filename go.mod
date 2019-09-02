module github.com/paysuper/paysuper-recurring-repository

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/InVisionApp/go-logger v1.0.1 // indirect
	github.com/ProtocolONE/go-micro-plugins v0.3.0
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/protobuf v1.3.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/go-micro v1.1.0
	github.com/micro/go-plugins v1.1.0
	github.com/paysuper/paysuper-database-mongo v0.1.0
	github.com/prometheus/client_golang v1.0.0
	go.uber.org/zap v1.10.0
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.6

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

replace github.com/apache/thrift => github.com/apache/thrift v0.12.0
