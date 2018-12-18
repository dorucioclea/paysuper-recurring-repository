package main

import (
	"github.com/ProtocolONE/payone-repository/internal"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	"log"
)

func main() {
	conf, err := internal.NewConfig()

	if err != nil {
		log.Fatalln(err)
	}

	app := internal.NewApplication()
	app.InitDatabase(conf)
	app.InitService()

	defer app.Database.Close()

	app.Run()
}