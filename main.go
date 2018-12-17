package main

import (
	"github.com/ProtocolONE/payone-repository/internal"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

func main() {
	app := internal.Application{}
	app.Init()
	app.Run()
}
