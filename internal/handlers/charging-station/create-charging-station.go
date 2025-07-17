package handlers

import (
	"log"
	"net/http"
	"ocpp-smart-charging/internal/database/mongo"
)

func CreateChargingStationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result, err := mongo.AddChargingStation(db)
	if err != nil {
		log.Fatalf("Can't create LoadBalancer! %+v", err)
	}
	log.Printf("CreateLoadBalancer result: %+v", result)
}
