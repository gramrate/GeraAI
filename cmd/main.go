package main

import (
	"gera-ai/internal/app"
)

func main() {
	GeraApp := app.NewGeraApp()

	app.Start(GeraApp)
}
