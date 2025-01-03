package main

import (
	"gera-ai/internal/app"
)

func main() {
	// create app
	GeraApp := app.NewGeraApp()

	// start app
	app.Start(GeraApp)
}
