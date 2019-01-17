package main

import (
	"github.com/ProtocolONE/payone-repository/internal"
)

func main() {
	app := internal.NewApplication()
	app.Init()

	defer app.Database.Close()

	app.Run()
}