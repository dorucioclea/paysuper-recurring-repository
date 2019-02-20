package main

import (
	"github.com/ProtocolONE/paysuper-recurring-repository/internal"
)

func main() {
	app := internal.NewApplication()
	app.Init()

	defer app.Stop()
	app.Run()
}