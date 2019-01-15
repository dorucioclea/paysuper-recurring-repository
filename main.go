package main

import (
	"github.com/ProtocolONE/payone-repository/internal"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

func main() {
	app := internal.NewApplication()
	app.Init()

	defer app.Database.Close()

	app.Run()
}