package main

import (
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/paysuper/paysuper-recurring-repository/internal"
)

func main() {
	app := internal.NewApplication()
	app.Init()

	defer app.Stop()
	app.Run()
}
