package main

import (
	"ocpp-smart-charging/internal/app"
)

func main() {
	appInstance := app.NewApp()
	appInstance.Run()
}
