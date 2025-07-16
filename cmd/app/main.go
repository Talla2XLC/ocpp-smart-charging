package main

import (
	"log"
	"ocpp-smart-charging/internal/app"
	"ocpp-smart-charging/internal/database/mongo"
)

func main() {
	appInstance := app.Run()

	loadBalancers, err := mongo.GetLoadBalancers(appInstance.MongoDB)
	if err != nil {
		log.Fatalf("Can't get LoadBalancers! %+v", err)
	}
	log.Printf("GetLoadBalancers result: %+v", loadBalancers)
}
