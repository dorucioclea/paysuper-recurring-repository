package main

import (
	"github.com/paysuper/paysuper-recurring-repository/internal"
)

func main() {
	app := internal.NewApplication()
	app.Init()

	defer app.Stop()
	app.Run()
}